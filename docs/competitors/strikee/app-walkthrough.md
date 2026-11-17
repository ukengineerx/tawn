# Strikee app walkthrough

What the logged-in product does, screen by screen, as seen in the public live demo on 2026-11-17. The demo logs into a sample venue, "PlayArena Multi-Sport Arena" (Andheri West, Mumbai), on the Ultra trial.

The review had two passes. The first opened every screen and form without submitting anything. The second ran the flows that only show themselves when you commit: settling a tab, refunding, deleting, the tournament wizard past step 1, and losing the connection. The demo account is shared (a change made in one browser shows up for every other visitor), so the second pass used ₹1–₹120 amounts on an existing sample customer and cleaned up afterwards wherever the app allowed it. All times are India time; the browser was set to Asia/Kolkata so the business day lined up.

Route names are the app's own URLs on `strikee.in`. Screenshots are in [screenshots/](screenshots/).

## App shell

- Sidebar in two groups. **Run the floor:** Cashbook, Sessions (with a live count), Slot Board, Booking requests, Members (with a count), Memberships, Add-ons, Facilities, Tournaments. **Owner only:** Dashboard, Reports, Settings, Billing.
- **Venue picker** (`/welcome`): "Welcome, Joy. Pick a venue to enter", one row per venue with its sports, and "Set up a new venue · same login, separate book". A side panel shows a sample of "what tonight looks like" and says "Live numbers fold in post-launch". The demo account has three venues: Mumbai, Surat and Rajkot. ([screenshot](screenshots/venue-picker.jpg))
- Footer: trial countdown with an Upgrade link, "Tutorials · 4 how-to videos · 2 min", "Email support · replies within 4 business hours", and the signed-in user.
- First visit prompts "Help us reach you" for a mobile number so the founder can WhatsApp setup tips. It can be skipped with "Later", but comes back on later visits.
- A **Quick start** checklist (Tables → Customer → Session → Settle) pops up over Sessions until it's skipped. ([screenshot](screenshots/onboarding-quick-start.jpg))
- A 9-step guided tour runs on first load: Cashbook, Sessions, Members, Facilities, Owner Dashboard, Reports, Tournaments, then "Typical flow: customer walks in → Sessions → Record Session → pick member → pick facility → done." It can be replayed from Settings.
- Staff use it as an installable web app (PWA); there is no native app.

## Cashbook (`/cashbook`)

- One business day at a time, with a **05:00 AM cutoff**, so play after midnight counts toward the previous evening. Today and Yesterday buttons plus a date picker. ([screenshot](screenshots/cashbook-full.jpg))
- Summary cards: Today's balance (live), Cash in hand, Expenses, UPI / Card, Credit used, Credit settled.
- Transaction table with All / Income / Expenses / Refunds filters. Columns: time, description, type, mode, amount. Descriptions are generated sentences, for example "Harsh Mehta's credit settlement of ₹920 via UPI. Previous balance: ₹1450. Remaining balance: ₹530".
- Types seen: Income, Settled (credit paid down), Store credit (refund converted to credit), membership sales.
- A payment split across modes becomes **one row per mode**, all carrying the same invoice number.
- **Transaction drawer:** clicking a row shows its ID (`INC-…`, `EXP-…`), amount, time, mode, customer, who recorded it and their role, and the note.
  - Rows created by a session payment are read-only here: "To edit or delete it, open the Sessions screen and update the matching session — changing it here would put the cashbook out of sync."
  - Manual rows have Edit and Delete. Delete confirms with "This will permanently remove the ₹1 expense… Today's KPIs will recalculate. This can't be undone." ([screenshot](screenshots/cashbook-expense-drawer.jpg))
- **Record expense** ("Log expense"): amount with quick chips (₹100 / ₹500 / ₹1,000); category (Floor supplies, Cue maintenance, Utilities, Staff, Snacks restock, Rent, Other); paid from Cash (from till), UPI or Card; paid to; "a note for the CA". A **"Paid from sales till"** toggle: "On: cash/UPI came out of the drawer — reduces cash-in-hand. Off: paid separately (your own funds), tracked but doesn't touch the till." "Mark as recurring" is marked coming soon. The save button reads "Log ₹X expense".
- **Record income:** amount chips; source (Walk-in cash, Tab settled, Equipment rental, Membership, Other); paid by Cash, UPI or Card; optional note.
- **No till count or day close.** There is no control to count cash, match it, or close the day, either today or on past days opened from the date picker ([screenshot](screenshots/cashbook-date-picker.jpg)). The marketing site's headline promise (a one-tap end-of-day match of counted cash against booked cash) isn't in the product.
- **Summary cards don't match the list.** With ₹2,300 of cash income and ₹650 of UPI income listed for the day, Cash in hand and UPI / Card both read ₹0. After a new ₹70 cash plus ₹50 UPI payment they still read ₹0. ([screenshot](screenshots/cashbook-after-settle.jpg))

## Sessions (`/sessions`)

- Header: "Live now · N tabs open · ₹ owed". Floor and Tabs views. The list doesn't refresh after a settle or close; it needs a page reload, and it sometimes sits on "Loading active sessions…" for a long time. ([screenshot](screenshots/sessions-stuck-loading.jpg))
- **Floor view:** a timeline of today's schedule (one row per facility in use, 10 AM to midnight) with live blocks, above the open tabs. ([screenshot](screenshots/sessions-floor-view.jpg))
- **Tab cards:** one per customer, showing the running total, "Settle tab →", one line per session (facility, status, units × rate, start time, amount, and a "···" menu), and "Add-ons on tab · + Add-on".
- **Older unpaid:** tabs from earlier that were never settled, with sessions marked **Overstayed** and an End button on each. ([screenshot](screenshots/sessions-older-unpaid.jpg))
- **Paid:** settled tabs per customer (sessions, settle time, modes, amount) with a Receipt link and nothing else.
- **Session menu ("···") on an open session:** Transfer to another customer, Edit session, Delete session. ([screenshot](screenshots/session-more-actions-menu.jpg))
  - *Edit session:* frames (with +/−), price per frame, new total, start time, end time ("Still active") and notes. The price field is labelled "Console" even on a snooker table. ([screenshot](screenshots/session-edit-dialog.jpg))
  - *Delete session:* "Owner only · logged with your name. If money was taken, use Cancel instead." It's meant to be a clean delete for sessions with no payment, but it called a session holding a ₹10 deposit "No payment was taken — this is a clean delete". No separate cancel action was found. ([screenshot](screenshots/session-delete-dialog.jpg))
- **Quick sell:** a three-step cart (Cart, Customer, Confirm) for selling add-ons without a game.
- **Book/Record** opens a six-step wizard:
  1. **Mode:** how the bill splits. *Regular*: one customer pays everything and can keep several sessions open until they leave. *Losers Pay*: pick several players, staff mark the losers at the end, and the cost becomes their debt. *Split Pay*: pick several players, and the cost splits equally into a debt per player.
  2. **Facility:** grouped by sport with price range and "Free now". Tables under maintenance or archived are hidden. A table showed "Free now" while an overstayed session from earlier was still open on it.
  3. **Kind:** Walk-in starts the timer immediately; Booking reserves a future date and slot, held until staff check the customer in.
  4. **Variation:** pick a saved package, for example Per Frame ₹120 or Per Hour ₹250. The price can be overridden for this session only, and the package's collection policy (at the end or upfront) can be switched for this session.
  5. **Customer:** search by name or phone, or add a new customer inline. Frames to play starts at 1 and is bumped from the session card as play continues; the estimated total updates.
  6. **Payment:** "Collect now (optional)". Apply the customer's store credit, take a deposit in Cash, UPI, Card or Online, split across several modes ("Add another mode"), add notes. The button changes from "Start session" to "Collect ₹10 & start" and shows "₹110 due at play". ([screenshot](screenshots/session-wizard-payment-step.jpg))

### Settling a tab

"Settle tab →" opens **Payment for {customer}** ([screenshot](screenshots/settle-tab-screen.jpg)):

- Member card: unpaid sessions and games, unpaid amount, credit available, and **net due** ("₹120 bill less ₹120 credit").
- **Store credit is applied automatically** with a switch ("Apply ₹120 store credit · Auto-applied. ₹280 stays on credit."). Turning it off puts the full bill back and adds Credit as a payment method. ([screenshot](screenshots/settle-store-credit-block.jpg))
- "Have a coupon or free-play card?" opens a code field with Check. ([screenshot](screenshots/settle-coupon-field.jpg))
- Game sessions: tick which unpaid sessions to settle now, or "Select all unpaid".
- Payment "split across methods": each method is Cash, Card, UPI or Credit with an amount; Cash takes a note, UPI takes a reference or UPI ID. "Add payment method" adds another.
- A balance check reads "Payment balanced · Total entered matches selected sessions" or "Over by ₹70". **Process payment stays disabled until it balances.** A footer shows "You are collecting ₹X". ([screenshot](screenshots/settle-split-payment.jpg))
- Afterwards the tab closes, the cashbook gets one row per method under one invoice, and the receipt lists each method. The Sessions list keeps showing the tab as open until the page is reloaded ([screenshot](screenshots/settle-result.jpg)).

**Receipts** are invoices: venue name and address, invoice number, bill-to with phone, one line per game with date, time range, count and rate, each payment method, and Print / Save PDF. Footer reads "Generated by Strikee". ([screenshot](screenshots/session-receipt.jpg))

### Refunds

- **No refund action was found for a settled payment.** Receipts offer Print and Save PDF; settled tab cards offer only Receipt; cashbook rows from sessions are read-only and point back to Sessions; the member profile has no per-payment action.
- Refunds do exist in the data model. The sample data has "Refund DEMO-MS-40045-R1 leg: CREDIT 400", where an advance was refunded as ₹400 of store credit, and refund legs are numbered per invoice (`-R1`).
- **An overpaid tab can't be closed.** After the price of a session with a ₹10 deposit was edited down to ₹0, the header showed "₹-10 owed" and closing the tab failed with "Payment methods total (0) does not match amount owed (unpaid sessions -10 + add-ons due 0 = -10)." The only way out was to edit the price back up. ([screenshot](screenshots/tab-close-overpaid.jpg))
- Settings lists refunds as owner-only and protected by the till passcode, so the button may sit behind a role or a state the demo doesn't reach. In the demo, an owner couldn't find it.

## Slot Board (`/slots`)

- A 30-minute grid of every bookable facility (9 in the demo) from 10:00 AM to 10:30 PM, for Today or Tomorrow.
- Click a free cell to create a booking. Click a booked cell to start it, view the customer, or cancel. Past slots are read-only.

## Booking requests (`/booking-requests`)

- An inbox for requests from the venue's public page, with four states: New → Awaiting payment → Confirmed → Closed.
- Stated flow: staff WhatsApp the customer their payment QR, confirm once the advance arrives, and the booking then appears on the Slot Board. There is no online payment step.
- Each staff member can turn on push notifications on their own phone so requests arrive even when the app is closed.

## Members (`/customers`)

- Active and Archived lists (name and phone). **Add member** needs only a name and phone number.
- **Credit watchlist:** members who owe money, total outstanding with the change this week, in three bands by share of their limit used: "Over the edge" (80%+), "Approaching" (50–80%) and "Healthy" (below 50%). Each row has **Send reminder** and **Collect ₹X**.
- Send reminder opens WhatsApp with the text "Hi {name}, this is a friendly reminder about your outstanding balance." Staff have to press send; nothing is automated.
- **Member profile** (`/customers/{id}`) ([screenshot](screenshots/member-profile.jpg)):
  - Credit usage against the limit, with Edit limit and a health badge.
  - Store-credit balance with Apply credit.
  - What they owe: tonight's unpaid sessions, standing credit balance and total exposure.
  - Actions: Send WhatsApp reminder, Settle standing credit, Make payment, Edit member, Archive.
  - Three ledgers: Game history ("click a row to see the invoice"), Credit settlements and Credit history ("every time this member deferred a payment"). Each row shows date, type, amount, mode, note, and **which staff member created it**.
  - Game history said "No game sessions or payments yet" for a customer who had paid twice that day.
- Refunds can be paid out as store credit, which the settle screen then applies automatically.

## Memberships (`/memberships`)

- **Plans:** prepaid bundles defined by price, units (hours or frames), sport (any table of that sport, or pinned to one) and validity in days. Demo examples: Snooker 20-Frame Pack ₹2,000 for 120 days; Gaming 10-Hour Pack ₹1,200 for 90 days; Pickleball 10-Hour ₹3,500 for 90 days.
- Rules from the in-app help:
  - Sold from a customer's profile and counted as income on the day of sale, once.
  - Drawn down automatically when the customer settles a visit; anything beyond the balance is paid normally.
  - Unused balance lapses at expiry. An owner or admin can extend the date, and the change is logged.
  - A customer can hold several, each with its own balance (say snooker frames and PS5 hours).
  - Explicitly not khata: "Khaata is play now, pay later. A membership is paid already, play it down."
- **Coupons:** free play, ₹ off, % off, or a free snack or drink.
  - *Printed cards* carry a unique code each, single-use across every branch they're valid for, so photocopies fail. Codes download as CSV.
  - *Shared codes* (for posters or posts) have an optional total limit, a per-customer limit matched by phone, and dates.
  - Discounts can target game time, snacks and drinks, or the whole bill. Staff apply at settle and can Undo before payment. A deposit already paid is never touched. Free-play coupons become a ₹0 membership.
  - Coupons can be paused or ended. Membership-bonus coupons are listed as coming next.

## Add-ons (`/add-ons`)

- Items by category (Food, Drink, Equipment, Other), with price, stock on hand and an optional cost. Without a cost the card says "add it to track margin". Each item has Restock; items can be archived.
- **Stocktake:** enter the shelf count per item against the system count. It flags shortages and resets the system to your count.
- **Sales history:** Today, 7 days, 30 days or a date; filters for open tabs, settled and voided; totals for sold, settled, add-on profit and margin.
- Quick sale (also on Sessions) sells items on their own.

## Facilities (`/facilities`)

- "Venue pulse": how many facilities are running, down or idle, utilisation, peak hour, and tonight's revenue per sport group.
- **Add a facility** is a three-step wizard. Picking a sport presets the pricing model:

  | Sport | Default pricing |
  |-------|-----------------|
  | Snooker, Pool, Billiards | per frame |
  | Table tennis, Badminton, Tennis | 60-minute slot |
  | Pickleball, Squash | 45-minute slot |
  | Box cricket | 60-minute slot |
  | Football | 90-minute slot |
  | Gaming zone | 60-minute slot |
  | Arcade | per session |
  | Custom | named by the owner |

- **Facility detail:**
  - Name and sport. Changing the sport only affects new variations; past sessions keep their pricing.
  - Opening hours per facility (or 24 hours), otherwise the venue's hours.
  - A toggle for time-slot booking; off means walk-in only.
  - **Variations:** each is a price per frame, hour or session, a duration (30 minutes to 3 hours), and a collection policy (at the end or upfront). Extra variations are meant for peak hours, tournament rates and the like.
  - Send for repair (status Down); Retire (status Off floor) with Reactivate.
- Live customers' public pages show variations also used for **player count** (1–4 players on one console), peak and off-peak rates, and 30-minute trial slots.

## Tournaments (`/tournaments`, Ultra)

- List filtered by status: Live, Registration, Ready, Completed, Draft, Cancelled. Cards show date, game, singles or doubles, format, entry fee, fill rate and prize pool.
- **Templates:** presets (Nightly, Daily, Weekend) that carry the format, fees and prize split. They unlock once a snooker table exists.
- **Create wizard** (`/tournaments/create`): four steps with Save as draft at any point and a live preview of the public card that fills in as you go. Nothing is saved until you save a draft or open registration; leaving the wizard leaves no draft behind.
  1. **Basics:** name, sport, Singles or Doubles. Snooker (per frame), EA SPORTS FC 27 and FC 26 (PS5, goals) are live; every other sport, FC 25 and eFootball say "coming soon".
  2. **The draw** ([screenshot](screenshots/tournament-step2-draw.jpg)):
     - Bracket format: Knockout (single elimination), Round robin, or Groups → KO.
     - Field size 8, 12, 16, 20, 24 or 32 ("keep it a multiple of 4 so every group is full. Extra sign-ups go to the waitlist").
     - Match length 6, 8, 10 or 12 minutes ("4-minute halves — set the same half length in the game's match settings").
     - Group stage: players per group (3–6), top 1/2/3 go through, best third-placed (none, 2, 4), random or seeded draw, points 3·1·0 or 2·1·0. A summary line works it out: "16 players → 4 groups of 4 · 24 group matches · everyone plays at least 3. 8 go through → straight into the Quarter-finals. 7 knockout matches."
     - Tiebreaks: points → goal difference → goals scored → head-to-head, then the organiser decides.
     - Knockout rules: extra time before penalties (on by default), 3rd-place match.
     - Scoring on the day: Final score (one scorer enters it at the end) or Live goals.
  3. **When & where** ([screenshot](screenshots/tournament-step3-when-where.jpg)): single session or multi-day; start date (required) and first-cue time; registration close date and time (optional); venue; WhatsApp number for entries; Instagram handle for DMs. The last two become "Message on WhatsApp" and "DM on Instagram" buttons on the public page.
  4. **The pot** ([screenshot](screenshots/tournament-step4-pot.jpg)): prize pot is either *Fixed / guaranteed* (a sponsor-backed amount independent of entries) or *Entry-funded* (entry fee with ₹200 / ₹500 / ₹1,000 chips and a 0%, 10% or 20% house cut, worked out as "16 entries × ₹500 = ₹8,000 collected · house keeps ₹800"). Prize split Top 2, Top 3 or Top 8, each place in % or ₹ or as a trophy, plus a top-scorer prize, and "Add place". The final button is **Open registration**. ([screenshot](screenshots/tournament-step4-entry-funded.jpg))
  - The default split (₹1,500 / ₹700 / ₹300) doesn't rescale when the pool changes: with an entry-funded ₹7,200 pool it reads "₹2,500 in cash prizes of a ₹7,200 pool — the rest stays with the house".
- **Tournament detail:** a public share link, Add player, signups against capacity, entry fees collected, prize pool. Tabs for Overview, Players (with payment mode per entry), Pot (collected and pending), Sponsors and Settings. "Make the draw" needs at least 2 players; "Just close registration" is the alternative.
- **Public tournament page** (`/public/tournaments/{id}`, no login): spots filled, prize split (winner, runner-up, top scorer), "Enter on WhatsApp" and "DM to enter on Instagram" buttons, "pay at the lounge or via UPI", format rules (groups of 6, win 3, draw 1, top 2 go through; knockouts with extra time and penalties), house rules, directions, and a "Run your venue's tournaments on Strikee" link at the bottom.

## Dashboard (`/owner-dashboard`)

- Period options: Daily, Weekly, Monthly, Yearly, or a date range. **It defaults to Monthly, which was empty in the demo; Yearly shows the full sample data.**
- **Overview** ([screenshot](screenshots/dashboard-yearly-overview.jpg)): revenue, expenses, refunds (count and amount), add-on collections, net profit with margin, new customers, daily average against the previous period, credit used and settled, peak hour, and revenue by game type.
- **Financial** ([screenshot](screenshots/dashboard-yearly-financial.jpg)): profit and loss, daily income split into cash and online, revenue against expenses and net, revenue by game type, and a transaction breakdown.
- **Customers** ([screenshot](screenshots/dashboard-yearly-customers.jpg)): customer metrics, top 10 customers, retention (Overview, Segments, Visit intervals), average session duration, and rule-based "Insights overview" and "Recommendations" written as plain sentences.
- **Operations** ([screenshot](screenshots/dashboard-yearly-operations.jpg)): peak hours by revenue and session count, and session duration by weekday, by hour and as a trend.

## Reports (`/reports`)

- Pitched as an "exports center"; the dashboard is for insight, this page is for documents to share. ([screenshot](screenshots/reports-last-month.jpg))
- Period presets: 7 days, This month, Last month, FY 25-26. Default format is PDF or CSV. "This month" was empty in the demo; "Last month" had data.
- Ready-made documents, each showing what it will contain before download:
  - **P&L Statement**, "signed PDF ready for your CA" (118 transactions for last month in the demo).
  - **Member ledger:** members on credit with balance, limit and days since the last update.
  - **Staff payout summary:** transactions, cash, online, credit issued and revenue per teammate.
- **Custom report:** pick any of Sessions, Payments, Credit ledger, Members, Facilities and Staff activity (Tournaments "soon"). It shows an estimated row count and a live preview, then builds one document with a section per dataset.

## Settings (`/settings`)

- **Identity:** club logo, name and address, with a live receipt preview "as your members see it". Changes apply to future receipts only.
- **Business details:**
  - GST number (printed on tax invoices; blank if not registered) and phone number.
  - Currency: INR, USD, GBP, EUR, AED or SGD.
  - Operating hours per day; sessions started outside them get a "late" flag in reports.
- **Team:** invite teammates and see a permission matrix:

  | Permission | Admin | Co-owner | Owner |
  |------------|:-----:|:--------:|:-----:|
  | Start and complete sessions | yes | yes | yes |
  | Take payments and settle credit | yes | yes | yes |
  | Add and archive members | yes | yes | yes |
  | Set credit limits | yes | yes | yes |
  | Change pricing and facility setup | yes | yes | yes |
  | See owner dashboard | no | yes | yes |
  | Business financials (revenue, profit) | no | yes | yes |
  | Export reports | no | yes | yes |
  | Refund and delete sessions | no | no | yes |
  | Invite and remove teammates | no | yes | yes |
  | Transfer ownership, delete venue | no | no | yes |

- **Security (till lock):** an optional 4-digit passcode that protects refunds, deletions, credit settlement and pricing, even for logged-in staff. It auto-locks after 15 minutes idle. It isn't set in the demo, so deletes and edits went through without a prompt.
- **Plan:** current plan, invoice downloads, cancel renewal. The demo shows "Professional ₹399/month", which isn't on the public price list.
- **Danger:** replay the tour; deleting the venue goes through contacting Strikee.
- **Public page** (`/settings/public-page`): sections for Identity (web address `strikee.in/{slug}`, public name, tagline, about, cover and logo), Layout, Gallery, Hours, Contact, Catalog, Offerings, Stations and Deals. The page stays a draft until a web address is saved.

## Billing (`/subscribe`)

- Trial status, days used, and **what stops when the trial ends**: new sessions are blocked, staff logins are disabled, and facilities, tournaments and bracket pages freeze. Data stays readable; nothing is deleted.
- Pro and Ultra cards, monthly or yearly. Payment by UPI, card or netbanking, GST included, full refund within 7 days.
- **ROI math from the venue's own data:** price per day (₹16.44), share of a typical weekend day's revenue (0.62%), the number of box-cricket slots it equals (7), and cost per logged session (₹4.76, based on the last 30 days).
- First-invoice preview from "Strikee Technologies Pvt. Ltd." with IGST at 18%.
- FAQ: cancelling mid-year, switching yearly to monthly, extra seats, GST, what Pro loses, adding a second venue.

## Public venue page (`strikee.in/{slug}`)

Seen on two live pages: Clash X Gamecafe in Morbi, and Strikee Gaming Lounge Eru in Navsari, which belongs to Strikee itself. That page describes itself as "our second gaming lounge in Navsari, part of the Strikee family alongside Strikee Gaming Lounge and Strikee Pool & Snooker", so the company runs three venues of its own. The flagship lounge's page, `strikee.in/strikee-gaming-lounge`, returned HTTP 500 in the first pass and in the second loaded as "Application error: a client-side exception has occurred".

- About text written for search ("a PS5 game centre in Navsari"), address with Google Maps directions, and a "Review us on Google" link.
- Prices per station group by player count (1–4), with **peak and off-peak** columns ("Peak = weekdays after 5 PM, weekends & public holidays all day") and a cheaper PS4 station; 30-minute trial rates on another page.
- Station cards ("Pick your console"), each with its own named line-up ("Blockbusters & Showdowns", "Party & 4-Player Chaos") and a Book button.
- A game menu (27–28 titles) filterable by genre, with player counts, "New" badges, and a "coming soon" title.
- Photo gallery, opening hours per day, amenities (AC, parking, food, Wi-Fi, card/UPI accepted, tournaments, kid friendly and so on).
- An FAQ written for search engines (cost, hours, games, group size, location, how to book, facilities), plus geo and Open Graph tags and a generated share image.
- "Call to book" and WhatsApp with a prefilled message; "Page by Strikee" at the bottom.

## Offline behaviour

The app registers a service worker (`/sw.js`) and a web manifest, but its cache storage is empty, so nothing is available without a connection.

- **Pages don't load offline.** Clicking a sidebar link or reloading while offline shows the browser's own "No internet" page. ([screenshot](screenshots/offline-reload.jpg))
- **An open page keeps working.** Dialogs still open, and there's no offline banner.
- **Writes wait and retry.** A ₹1 expense saved while offline sat on "Saving…" with no message ([screenshot](screenshots/offline-expense-save.jpg)). When the connection came back it saved exactly once, with no toast. This only works if the tab stays open; closing or reloading it loses the entry.

## Rough edges seen in the demo

Some of these may come from the demo's sample data rather than the product.

- No till-count or day-close step, even though the marketing site's main promise is the matched till.
- Cashbook summary cards show ₹0 cash and ₹0 UPI while the list below has cash and UPI income, including a payment made during the review.
- No way to refund a settled payment, and an overpaid tab can't be closed at all.
- The delete dialog calls a session with a deposit "a clean delete" and points to a "Cancel" option that doesn't exist.
- Sessions lists don't refresh after settling or closing a tab, and sometimes hang on "Loading active sessions…". The sidebar said "2 LIVE" while Sessions said "0 tabs open".
- A member's game history was empty although they had paid twice that day.
- The dashboard and reports open on the current month, which was empty; the data is under Yearly and Last month.
- A table shows "Free now" while an overstayed session is still open on it.
- The session edit dialog labels a snooker table's price "Console".
- A member with store credit shows "-80%" credit usage.
- The venue switcher says "No tables yet" for a venue with 13 facilities.
- The plan in Settings (Professional ₹399/month) doesn't match public pricing (Pro ₹449/month).
- Tournament prize splits don't rescale with the pool; a public tournament page lists a ₹2,500 prize pool but a ₹15,000 winner's prize.
- The public page of Strikee's own flagship lounge doesn't load.

## Test data left in the demo

The demo has no way to refund a settled payment, so two small payments made during the review remain on the sample customer Chandan (+916366311856), dated 17 Nov 2026:

- Invoice DEMO-MS-40046: one frame on Table 2, ₹120, paid ₹70 cash and ₹50 UPI.
- A one-frame session on Table 4 at ₹10, paid by a ₹10 UPI deposit.

The ₹1 test expense and the test tournament were removed (the tournament was never saved).
