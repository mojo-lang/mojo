#!/usr/bin/env python3
"""Copy selected NCraft YAML templates without overwriting project configuration."""

import argparse
from pathlib import Path


TEMPLATE_DIR = Path(__file__).resolve().parent.parent / "configs"
COMPONENTS = ("db", "jwt", "logs", "metrics", "redis", "sd", "server", "tracing")
DEFAULT_COMPONENTS = ("server", "logs", "metrics", "tracing", "sd")


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--dest", type=Path, required=True, help="Target configuration directory")
    selection = parser.add_mutually_exclusive_group()
    selection.add_argument("--components", nargs="+", choices=COMPONENTS,
                           help="Select components (default: server logs metrics tracing sd)")
    selection.add_argument("--all", action="store_true", help="Copy all eight generic templates")
    args = parser.parse_args()

    components = COMPONENTS if args.all else (args.components or DEFAULT_COMPONENTS)
    sources = [TEMPLATE_DIR / f"{name}.yaml" for name in dict.fromkeys(components)]
    for source in sources:
        if not source.is_file():
            parser.error(f"Missing template: {source}")

    try:
        args.dest.mkdir(parents=True, exist_ok=True)
        for source in sources:
            target = args.dest / source.name
            data = source.read_bytes()
            try:
                with target.open("xb") as output:
                    output.write(data)
            except FileExistsError:
                print(f"SKIP existing: {target} (compare and merge manually)")
            else:
                print(f"COPIED: {target}")
    except OSError as exc:
        parser.exit(1, f"Copy failed: {exc}\n")

    print("Templates copied without environment expansion; review component settings before startup.")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
