#!/usr/bin/env bash
set -euo pipefail

usage() {
  echo "Usage: $0 <vMAJOR.MINOR.PATCH[-PRERELEASE][+BUILD]> [os/arch ...]"
  echo "Defaults: linux/amd64 linux/arm64 darwin/amd64 darwin/arm64 windows/amd64 windows/arm64"
  echo "Output: DIST_DIR, or <repository>/dist/<version> (must be absent or empty)"
}

if [[ ${1:-} == --help || ${1:-} == -h ]]; then
  usage
  exit 0
fi
if [[ $# -eq 0 ]]; then
  usage >&2
  exit 1
fi

version=$1
shift
# Restrict the version before using it in paths or linker arguments.
if [[ ! $version =~ ^v(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)(-[0-9A-Za-z-]+(\.[0-9A-Za-z-]+)*)?(\+[0-9A-Za-z-]+(\.[0-9A-Za-z-]+)*)?$ ]]; then
  echo "Invalid release version: $version" >&2
  exit 1
fi

if [[ $# -eq 0 ]]; then
  set -- linux/amd64 linux/arm64 darwin/amd64 darwin/arm64 windows/amd64 windows/arm64
fi
for target in "$@"; do
  case "$target" in
    linux/amd64|linux/arm64|darwin/amd64|darwin/arm64|windows/amd64|windows/arm64) ;;
    *) echo "Unsupported target: $target" >&2; exit 1 ;;
  esac
done

for tool in go tar zip; do
  if ! command -v "$tool" >/dev/null 2>&1; then
    echo "Required tool not found: $tool" >&2
    exit 1
  fi
done
if command -v sha256sum >/dev/null 2>&1; then
  checksum=(sha256sum)
elif command -v shasum >/dev/null 2>&1; then
  checksum=(shasum -a 256)
else
  echo "Required tool not found: sha256sum or shasum" >&2
  exit 1
fi

repo_root=$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)
output_dir=${DIST_DIR:-"$repo_root/dist/$version"}
mkdir -p "$output_dir"
output_dir=$(cd "$output_dir" && pwd)
if [[ -n $(ls -A "$output_dir") ]]; then
  echo "Output directory must be empty: $output_dir" >&2
  exit 1
fi

staging_dir=$(mktemp -d "${TMPDIR:-/tmp}/mojo-release.XXXXXX")
trap 'rm -rf "$staging_dir"' EXIT
mkdir -p "$staging_dir/archives"

for target in "$@"; do
  target_os=${target%/*}
  target_arch=${target#*/}
  package_name="mojo_${version}_${target_os}_${target_arch}"
  package_dir="$staging_dir/$package_name"
  mkdir -p "$package_dir"
  binary=mojo
  if [[ $target_os == windows ]]; then
    binary=mojo.exe
  fi

  echo "Building $version for $target..."
  (
    cd "$repo_root/go"
    CGO_ENABLED=0 GOWORK=off GOOS="$target_os" GOARCH="$target_arch" \
      go build -mod=readonly -trimpath \
      -ldflags "-s -w -X github.com/mojo-lang/mojo/go/cmd/mojo/cmd.version=$version" \
      -o "$package_dir/$binary" ./cmd/mojo
  )
  cp "$repo_root/README.md" "$package_dir/README.md"

  if [[ $target_os == windows ]]; then
    (cd "$package_dir" && zip -q "$staging_dir/archives/$package_name.zip" "$binary" README.md)
  else
    COPYFILE_DISABLE=1 tar -czf "$staging_dir/archives/$package_name.tar.gz" \
      -C "$package_dir" "$binary" README.md
  fi
done

# Publish local output only after every requested target has built successfully.
(
  cd "$staging_dir/archives"
  "${checksum[@]}" mojo_* > checksums.txt
)
cp "$staging_dir/archives/"* "$output_dir/"
echo "Release archives and checksums: $output_dir"
