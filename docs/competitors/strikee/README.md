---
name: Strikee
slug: strikee
url: https://strikee.in
category: sports venue management
segments: [snooker and pool, billiards, badminton, pickleball, table tennis, tennis, squash, box cricket, football turf, arcades, gaming zones]
region: India
founded: 2024
hq: Mysore, Karnataka
stage: public beta (product v0.7)
pricing: Pro ₹449/month or ₹4,499/year; Ultra ₹599/month or ₹5,999/year; GST included
platform: mobile-responsive web app (installable PWA), no native app
last_reviewed: 2026-11-17
evidence: public website, plus a full walkthrough of the public live demo including test settles, refund and delete attempts, and an offline test
---

# Strikee

Strikee is bookkeeping-first software for Indian sports and gaming venues. Counter staff run timed sessions per table, court or console, and bill them per frame, hour or slot. Bills settle across cash, UPI, card, store credit and customer credit (khata), and everything lands in a daily cashbook. Around that core it has a slot board, prepaid memberships and coupons, snack inventory, an owner dashboard, exports for the accountant, a public venue page that takes booking requests, and tournaments with public bracket pages.

The screen-by-screen detail, including form fields and business rules, is in [app-walkthrough.md](app-walkthrough.md). This page is the summary.

## Who it targets

- Owner-operated venues where staff bill from one counter, especially in tier 2–4 cities. Snooker and pool are the home base; courts, turf and PS5/Xbox gaming zones came later. A live customer seen during review is a PS5 gaming café in Morbi, Gujarat.
- Copy is written for the owner in Indian small-business vocabulary: khata, udhaar, "send to CA", Hinglish WhatsApp mock-ups.
- The pitch is replacing three tools: Excel for the cash, WhatsApp groups for bookings, and a notebook for credit.

## Positioning

- "Built by an owner": the founder ran a snooker and pool club in Gujarat for two years, and that club still tests every feature. The company runs three venues in Navsari (two PS5 gaming lounges and a pool and snooker hall), listed on strikee.in like any customer.
- Against Khatabook, Vyapar and myBillBook, the FAQ argues those are generic ledgers that don't know what a frame, table, court slot or session is.
- No payment lock-in: Strikee records payments but doesn't process them or take a cut. Even booking advances are collected by the venue over a WhatsApp QR.

## Pricing

| Plan | Monthly | Yearly | What you get |
|------|---------|--------|--------------|
| Pro | ₹449 | ₹4,499 (₹375/mo) | Up to 4 tables, courts, consoles or PCs. Unlimited sessions, customers and team. Khata, add-ons, cashbook, every report. |
| Ultra | ₹599 | ₹5,999 (₹500/mo) | Everything in Pro, no unit limit, tournaments, public venue page, memberships. |

- Prices include GST. Yearly is priced as 10 months. Full refund within 7 days. Payment by UPI, card or netbanking.
- Every venue starts on a 30-day Ultra trial with no card. When a trial lapses the venue goes read-only: new sessions are blocked and staff are logged out, but data stays readable.
- Price scales with playable units, not staff seats.
- The billing page sells the upgrade with the venue's own numbers: cost per day, share of a weekend day's revenue, and cost per logged session.

## Product at a glance

| Area | What it does |
|------|--------------|
| Sessions | Six-step start wizard covering bill mode (Regular, Losers Pay, Split Pay), facility, walk-in or booking, price package, customer, and an optional deposit across several payment modes. Settle screen that applies store credit automatically, takes coupons, splits across methods and won't submit until the amounts balance. Invoice-style receipts with Print and PDF. |
| Cashbook | Day view with a 5 AM cutoff; cash, UPI/card, credit used and credit settled; income, expense and refund entries; a toggle for whether an expense came out of the till. Rows created by sessions are read-only. No day close, and the summary cards disagreed with the list. |
| Slot Board | 30-minute grid across all bookable facilities, today and tomorrow. |
| Booking requests | Inbox for requests from the public venue page: new, awaiting payment, confirmed, closed. Push notifications per staff phone. |
| Members and credit | Watchlist banded by share of credit limit used; WhatsApp reminder template; profile with credit, store credit, exposure, and ledgers that record which staff member made each entry. |
| Memberships and coupons | Prepaid hour or frame bundles with expiry; printed single-use or shared coupon codes; ₹ off, % off, free play or free item. |
| Add-ons | Snack and drink inventory with cost and margin, restock, stocktake and sales history. |
| Facilities | 12 sports plus custom, each with default pricing; per-facility hours, booking toggle, price variations, repair and retire states. |
| Tournaments (Ultra) | Templates; a four-step wizard covering knockout, round robin or groups then knockout, match length, tiebreaks, scoring, dates, WhatsApp and Instagram entry, and a fixed or entry-funded pot with a house cut; players, pot, sponsors, public pages. Snooker, EA FC 26 and FC 27 only. |
| Dashboard | Overview, financial, customer and operations tabs (P&L, revenue by game type, peak hours, retention, written insights). Opens on the current month, which was empty; Yearly shows the data. |
| Reports | P&L, member ledger, staff payout summary, and a custom multi-dataset report, as PDF or CSV. |
| Settings | Receipt identity, GSTIN, six currencies, operating hours with late flags, Admin / Co-owner / Owner permissions, 4-digit till passcode, public page editor. |

## What the demo changed from the marketing-site picture

- **Day close:** the site's headline promise is a cashbook that matches the till at close. The demo cashbook shows running totals but has no step for entering counted cash or closing the day, on today or any past day. Its cash and UPI summary cards also read ₹0 while the list below had cash and UPI income.
- **Refunds:** the permission matrix and till passcode both mention refunds, and the sample data has a refund to store credit, but no refund action was reachable for a settled payment. A tab that ends up overpaid can't be closed.
- **Offline:** it's installable, but nothing loads without a connection. A write made offline waits and saves once the connection returns, as long as the tab stays open, with no indicator to staff.
- **Player booking:** the site lists online booking as a roadmap item, but a request-based version exists. Players request from the venue's public page, the venue collects an advance over WhatsApp QR, and staff confirm it onto the slot board. There's still no instant online booking or payment.
- **Exports:** both PDF and CSV are available, plus a staff payout summary and a custom report builder.
- **Roles:** Admin can change pricing and credit limits but can't see financials, export, refund or manage the team. Only the Owner can refund or delete sessions.
- **Extras not on the marketing site:** coupons and printed free-play cards, add-on stocktake and margin, a till passcode, six currencies, store credit from refunds, push notifications for booking requests.

## Gaps

From their own site and in-app "coming soon" labels:

- No automated reminders or customer notifications; reminders are a WhatsApp link staff send by hand.
- No GST summary export (GSTIN does print on invoices).
- No floor-only staff role.
- Recurring expenses (rent, salaries) are "coming soon".
- Tournaments cover only Snooker, EA FC 26 and EA FC 27.
- No native app; it's an installable web app.

Seen in the demo:

- No till count or day-close step, despite the marketing.
- No refund for a settled payment, and no way to close an overpaid tab.
- Not usable offline beyond the page already open.
- No online payment for bookings or tournament entries; everything goes through WhatsApp, UPI QR, or pay at the venue.
- Lists that don't refresh, summary numbers that contradict the ledger, and an empty member history (listed at the end of the walkthrough).

## Ideas worth borrowing

These are useful whatever Tawn turns out to be:

1. **5 AM business-day cutoff**, so late-night play counts toward the right evening.
2. **Pricing as saved packages per facility**, each with a unit (frame, hour, session), a duration and a collection policy (upfront or at the end), overridable per session.
3. **Bill modes for group play:** Regular, Losers Pay and Split Pay, each producing a debt per player.
4. **"Paid from sales till" toggle on expenses**, which separates till cash from the owner's own money.
5. **Refunds as store credit**, offered automatically at the next payment.
6. **Till passcode for risky actions** (refunds, deletions, credit settlement, pricing) instead of a heavy role system.
7. **Credit watchlist banded by share of limit used**, with a one-tap message template.
8. **Coupon design:** printed single-use codes that fail when photocopied, versus shared codes with total and per-phone limits.
9. **Trial expiry that goes read-only instead of deleting**, and an upgrade page priced in the venue's own numbers.
10. **Public venue and tournament pages** that act as SEO pages, booking-request inboxes and "Powered by Strikee" acquisition, all without a payment gateway.
11. **A settle screen that can't go wrong:** store credit applied by default, pick which unpaid sessions to settle, split across methods, and the submit button stays disabled until the amounts balance.
12. **Ledger rows owned by their source:** cashbook entries created by a session payment can only be changed from the session, so the books can't drift from the floor.
13. **A tournament wizard that does the arithmetic:** it turns field size and group settings into "4 groups of 4 · 24 group matches · 7 knockout matches", and an entry fee and house cut into the pot.
14. **Prices by player count with peak and off-peak columns**, for consoles shared by up to four players.

## Distribution

- Search-driven content: a landing page per sport, "how to start a [sport] venue in India" guides, a blog, free calculators, and every customer's public venue and tournament pages hosted on strikee.in.
- Founder-led sales and support over WhatsApp. The app asks for the owner's mobile number on first use for exactly that.
- Onboarding: a 9-step in-app tour, 4 short tutorial videos, and email support within 4 business hours.

## Company and tech

- Founder: Joy Patel (product and engineering). Started in Gujarat in 2024; now building in Mysore. The terms of service name LETSTHRIVE.IO; the in-app invoice preview names Strikee Technologies Pvt. Ltd.
- Contact: hello@strikee.in, X [@strikee_club](https://x.com/strikee_club).
- Visible stack: Next.js, Supabase (Postgres with row-level security, and storage for images), Radix-style dialogs, Lucide icons, driver.js for the product tour, installable web app.
- Versions at review: product "public beta v0.7", marketing site v4.2.1.

## Open questions

1. Where do refunds of settled payments happen for real customers? The demo shows refunded store credit in its sample data but no button for it.
2. Do the cashbook summary cards disagree with the ledger in real accounts too, or only in the generated demo data?
3. How many venues are paying, given the beta label, founder-led support and the company running three venues of its own?

Answered in the second pass: there is no day close anywhere; the dashboard has full data under Yearly; offline, open pages keep working and writes wait for the connection, but nothing loads.

## Sources

Checked 2026-11-17.

- [strikee.in](https://strikee.in): modules, pricing, FAQ
- [strikee.in/features](https://strikee.in/features): module deep dives
- [strikee.in/about](https://strikee.in/about): founder story, roadmap
- [strikee.in/terms-of-service](https://strikee.in/terms-of-service): legal entity, usage rules
- Live demo via "View live demo" on the homepage: every sidebar section and form, then a second pass that settled, edited, deleted and went offline with small test amounts
- [strikee.in/clash-x-gamecafe](https://strikee.in/clash-x-gamecafe): a live customer's public venue page
- [strikee.in/strikee-gaming-lounge-eru](https://strikee.in/strikee-gaming-lounge-eru): one of Strikee's own venues
- [strikee.in/strikee-gaming-lounge](https://strikee.in/strikee-gaming-lounge): Strikee's flagship lounge page, which didn't load
- A demo tournament's public page at `strikee.in/public/tournaments/{id}`
