# Docs

Docs here come in two layers. Agents write the first; people read the second.

| Layer | File | Written by | Read by | Edited? |
|-------|------|------------|---------|---------|
| Source | Markdown (`README.md`, `app-walkthrough.md`, …) | Agents, and people when they want to | Agents, and people checking a detail | Yes. This is the source of truth. |
| Explainer | `explainer.html` in the same folder | Generated from the source | People | No. Change the source, then regenerate. |

## Why two layers

Agent docs are complete and exact, which makes them long and organised like the thing they describe: one section per screen, per endpoint, per table. That suits an agent looking something up. It doesn't suit a person trying to understand the subject and decide what to do about it.

An explainer covers the same facts for a reader. It opens with the short version, follows the order a person would think in, uses diagrams for flows and states, puts screenshots next to the point they support, and says what each finding means for Tawn. It should take under ten minutes to read.

## Rules

- **One explainer per source folder**, saved as `explainer.html` beside the markdown it is built from, so relative paths to `screenshots/` work.
- **Never edit an explainer by hand.** Fix the source, then regenerate. Agents should not read explainers as input; they read the source.
- **Every explainer is stamped** with the source files and their hashes, and with its published URL. `make docs-check` uses the stamp to say which explainers are out of date.
- **Explainers carry the source's evidence labels.** If the source says something was seen in the product versus claimed on a website, the explainer shows that too.
- **Not every doc needs one.** Short docs, such as [infrastructure.md](infrastructure.md), are already readable. Write an explainer when a source is long, dense, or spread over several files.

## Making or refreshing one

Ask Claude Code to "make the explainer for `<folder>`" or run `/explainer <folder>`. The steps are in [.claude/skills/explainer/SKILL.md](../.claude/skills/explainer/SKILL.md). Then:

```bash
make docs-check
```

## Explainers

| Folder | Explainer | Published |
|--------|-----------|-----------|
| [competitors](competitors/) | [explainer.html](competitors/explainer.html) (the features matrix) | [Venue Capability Map](https://claude.ai/artifact/QUito4j8RMwMPBBzruDyAC) |
| [competitors/strikee](competitors/strikee/) | [explainer.html](competitors/strikee/explainer.html) | [Strikee, explained](https://claude.ai/artifact/KQbzuCEGsu49yhCThrXHUp) |
