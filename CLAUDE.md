# Tawn

## Docs have two layers

Markdown under `docs/` is the source of truth. Agents write it and read it. Some folders also have an `explainer.html`, a page generated from that markdown for people to read. See `docs/README.md`.

- Never edit an `explainer.html` by hand, and don't use one as input. Read the markdown.
- After you substantially change markdown in a folder that has an explainer, regenerate the explainer with the `explainer` skill, or tell the user it's out of date.
- `make docs-check` lists stale explainers, and long docs that have none.
