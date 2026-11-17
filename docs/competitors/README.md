# Competitors

One directory per competitor. Each holds a `README.md` profile and a `screenshots/` folder. Profiles start with YAML front matter using the same keys, so they can be compared side by side or parsed later. [features.md](features.md) is the cross-competitor view: one row per capability, one column per competitor.

| Competitor | What it is | Price | Last reviewed |
|------------|------------|-------|---------------|
| [Strikee](strikee/README.md) | Sports venue management for Indian snooker halls, courts, turfs and gaming zones | ₹449–599/month | 2026-11-17 |
| [CourtReserve](courtreserve/README.md) | Not reviewed yet | | |

## Adding a competitor

1. Copy [_template.md](_template.md) to `<slug>/README.md`, where the slug is the product name in lowercase.
2. Fill in the front matter and sections, and list every source with the date you checked it.
3. If you get into the product (trial or demo), write what each screen does in `<slug>/app-walkthrough.md` and keep the profile as the summary. [Strikee's](strikee/app-walkthrough.md) is the example.
4. Save screenshots in `<slug>/screenshots/` as `<module>-<screen>.jpg` (or `.png`), for example `cashbook-day-close.jpg`.
5. Add a row to the table above and a column to [features.md](features.md).
6. Once the profile is substantial, generate its `explainer.html` for human readers (see [docs/README.md](../README.md)). [Strikee's](strikee/explainer.html) is the example.

## Keeping it honest

- Treat marketing-site claims as claims until someone sees them in the app. Profiles say which is which.
- Update `last_reviewed` whenever a profile is rechecked, and note pricing changes rather than overwriting them silently.
