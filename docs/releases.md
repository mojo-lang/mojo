# Mojo CLI releases

The release workflow builds `go/cmd/mojo` when a version tag such as `v0.1.1`
is pushed to GitHub. All six targets must build, the CLI/build-command tests
must pass, and the Linux amd64 binary must pass a smoke test before the workflow
uploads the archives and `checksums.txt` to a GitHub Release. New releases stay
as drafts until the upload completes. Tags with a prerelease suffix such as
`v0.2.0-rc.1` are marked as prereleases.

| OS | Architectures | Archive |
| --- | --- | --- |
| Linux | amd64, arm64 | `.tar.gz` |
| macOS (`darwin`) | amd64, arm64 | `.tar.gz` |
| Windows | amd64, arm64 | `.zip` |

Each archive contains `mojo` (or `mojo.exe`) and `README.md`. The CLI embeds its
version and reports it through `mojo --version`. Builds disable CGO and embed
the compiler's standard packages and templates already included in the Go
sources; the archive does not need a local checkout to run the CLI.

## Local packaging

Use macOS or Linux with Bash, the Go version required by `go/go.mod` (or newer),
`tar`, `zip`, and either `sha256sum` or `shasum`. Dependencies must be available
in the Go module cache or downloadable.

From the repository root:

```sh
# Build all six targets from the current working tree.
bash scripts/build-release.sh v0.1.1

# Build selected targets into a separate output directory.
DIST_DIR=/tmp/mojo-preview bash scripts/build-release.sh v0.1.1 darwin/arm64 linux/amd64
```

The default output is `dist/<version>/`, with names such as
`mojo_v0.1.1_linux_amd64.tar.gz`. The output directory must be absent or empty;
use a new `DIST_DIR` for another build. Failed builds leave no release assets
in the output directory. The script only builds archives; it does not create
tags, commit files, or publish releases. The supplied version labels the current
working tree; it does not check out that tag.

Verify downloaded archives with `sha256sum --check checksums.txt` on Linux, or
`shasum -a 256 --check checksums.txt` on macOS. On Windows, compare
`Get-FileHash .\mojo_v0.1.1_windows_amd64.zip -Algorithm SHA256` with the matching
entry in `checksums.txt`.

## Publishing a version

Commit and push the release workflow, script, and CLI changes before tagging.
The tagged commit must contain those files:

```sh
git tag -a v0.1.1 -m 'Release v0.1.1'
git push origin refs/tags/v0.1.1
```

Watch **Actions → Release Mojo CLI**, then download the assets from **Releases**.
The workflow uses the repository's built-in `GITHUB_TOKEN` with `contents: write`;
no personal token is required. GitHub Actions must be enabled for the repository.
If a run fails, fix the cause and rerun it when possible; source fixes require a
new commit and version tag. Uploads can be retried by rerunning the workflow.

The existing `v0.1.0` tag predates this workflow and will not automatically get
binary assets when the workflow is added. Publish a new version containing the
workflow; do not move an already published tag.

References: [GitHub tag-push triggers](https://docs.github.com/en/actions/reference/workflows-and-actions/events-that-trigger-workflows#push),
[GitHub CLI release publishing](https://cli.github.com/manual/gh_release_create).
