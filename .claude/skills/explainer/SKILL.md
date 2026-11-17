---
name: explainer
description: Build or refresh the human-readable explainer.html for a docs folder (for example docs/competitors/strikee) from the agent-written markdown in it, stamp it with source hashes, and publish it as an Artifact. Use when asked to make, update or regenerate an explainer, a readable version, or the human layer of a doc, or after substantially changing a source doc that already has an explainer.
---

# Explainer

Turns the agent-written markdown in one `docs/` folder into `explainer.html`, a page a person can read in under ten minutes. The rules for the two layers are in `docs/README.md`. The reference example is `docs/competitors/strikee/explainer.html`.

Argument: the folder, for example `docs/competitors/strikee`. If none is given, run `make docs-check` and offer the stale or missing ones.

## 1. Read the source

- Read every markdown file in the folder in full, plus any file it links to that it depends on (for competitors, also `docs/competitors/features.md`).
- List the folder's `screenshots/` and look at the ones you might use. Note anything a screenshot shows that the text doesn't say. Screenshots are sometimes captured after the text was written.
- If `explainer.html` already exists, read its `explainer:artifact` meta tag so you republish to the same URL.

## 2. Plan for the reader

Decide what a person needs to understand and decide, not what order the source is in.

- Open with the short version: what it is, three things that matter on each side, and what it means for Tawn.
- Arrange the body around how the subject is actually used (a night at the counter, a request's life cycle, a user's journey), not around the source's headings or the product's menu.
- Use a diagram wherever the source describes a sequence, a state machine, a split or a scale. Draw it with HTML/CSS or small inline SVG, or a `<pre class="mermaid">` block.
- When the source mixes up similar concepts, add a comparison table.
- Keep the source's evidence labels (seen in the product, claimed on the website, missing, roadmap) as small tags.
- Add a "For Tawn" note where a finding implies something for Tawn. Don't invent facts. Opinions go only in those notes.
- If you used anything that the source text doesn't contain (usually something seen in a screenshot), say so in the footer.
- Cut repetition. The explainer should be well under half the source's length.

## 3. Write it

- Load the `artifact-design` guidance through the Artifact tool's quickstart (`intent: "other"`) and follow its page contract: light and dark themes, phone width, and fonts only from Google Fonts.
- Write plainly, in short active sentences, and define each term of art once.
- Use `screenshots/<file>` as relative image paths, and let people click a screenshot to enlarge it.
- Put these tags right after `<title>`, one `explainer:source` per source file:

  ```html
  <meta name="explainer:source" content="docs/competitors/strikee/README.md sha256:<first 12 hex>">
  <meta name="explainer:artifact" content="https://claude.ai/artifact/<id>">
  ```

  Get each hash with `shasum -a 256 <file> | cut -c1-12`. Add the `explainer:artifact` tag after the first publish.
- End with a footer that names the source files and the build date, and says that the page is generated and should not be edited by hand.

## 4. Publish and record

- Publish with the Artifact tool. Map each screenshot used in `files` as `"screenshots/x.jpg": "docs/<folder>/screenshots/x.jpg"`. Pass `url` from the `explainer:artifact` tag when it exists.
- Add or update the folder's row in the Explainers table in `docs/README.md`.
- Run `make docs-check` and confirm the folder reports `ok`.
