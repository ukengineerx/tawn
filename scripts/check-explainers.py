#!/usr/bin/env python3
"""Report which docs explainers are out of date with their source markdown.

Each docs/**/explainer.html carries tags like
  <meta name="explainer:source" content="docs/x/README.md sha256:abcdef012345">
The script recomputes each source hash and reports ok, stale, or missing.
It also lists long markdown files whose folder has no explainer yet.
Exits 1 if any explainer is stale or names a missing source.
"""

import hashlib
import re
import sys
from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent
DOCS = ROOT / "docs"
LONG_DOC_BYTES = 8000
META = re.compile(
    r'<meta\s+name="explainer:source"\s+content="([^"\s]+)\s+sha256:([0-9a-f]+)"\s*/?>'
)


def short_hash(path: Path) -> str:
    return hashlib.sha256(path.read_bytes()).hexdigest()[:12]


def main() -> int:
    failed = False
    covered = set()

    for explainer in sorted(DOCS.rglob("explainer.html")):
        rel = explainer.relative_to(ROOT)
        sources = META.findall(explainer.read_text(encoding="utf-8"))
        covered.add(explainer.parent)
        if not sources:
            print(f"unstamped  {rel}  (no explainer:source tags)")
            failed = True
            continue
        problems = []
        for src, want in sources:
            path = ROOT / src
            if not path.is_file():
                problems.append(f"{src} is missing")
            elif short_hash(path) != want:
                problems.append(f"{src} changed")
        if problems:
            failed = True
            print(f"stale      {rel}")
            for p in problems:
                print(f"           - {p}")
        else:
            print(f"ok         {rel}")

    for md in sorted(DOCS.rglob("*.md")):
        if md.stat().st_size >= LONG_DOC_BYTES and md.parent not in covered:
            print(f"no page    {md.relative_to(ROOT)}  ({md.stat().st_size // 1000} KB, no explainer in folder)")

    return 1 if failed else 0


if __name__ == "__main__":
    sys.exit(main())
