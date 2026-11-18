# Internal tools: research

**Status: research only. Nothing on this page is decided.** It collects what we know so we can discuss and choose. Decisions go in the [Decision log](#decision-log) at the end once they're made.

"Internal tools" are the software Tawn the company uses to find, sell to, support and bill venue owners, and to run its own engineering. They are not features of the venue product.

Last updated 2026-11-18.

## Evidence labels

The proposals below came from a research report written by another agent on 2026-11-18. We spot-checked the claims that matter to a decision. Every claim carries one of these labels:

- **Checked**: we confirmed it ourselves, against the repo, its source code or the GitHub API, on the date given.
- **Report**: the research report says so, and we haven't confirmed it.
- **Inference**: our reasoning, not a fact from a source.

## Working set so far

These were chosen in earlier discussion as the starting point. None is installed yet, and any of them can still be swapped out.

| Need | Tool | Stack and licence | Known gaps |
|------|------|-------------------|------------|
| Find venues and open the first conversation | [Sellbird](https://github.com/ethanplusai/sellbird) | Python, SQLite, AGPL-3.0 | Sends email only. Free OpenStreetMap search is thin for small Indian venues. |
| Talk to owners | Chatwoot, Community Edition | Rails, Postgres, Redis, Sidekiq | Some features (the Captain AI assistant, SLAs, audit logs) are in the paid Enterprise edition. |
| Bill venues | [Billr](https://github.com/ethanplusai/billr) | Next.js, Supabase, Apache-2.0 | Card payments through Stripe only. No Indian GST invoice format. No outgoing webhooks. |
| Connect the tools | Bridge service | Go, written by us | Not written yet. |
| Launch the website | [website-launch-skills](https://github.com/ethanplusai/website-launch-skills) | Node, MIT | None that matter now. |
| Architecture map | [Map](https://github.com/ethanplusai/map) | Rust, MIT | A proof of concept, built for OpenCode. Build it from source rather than run the committed binary. |

Principles we agreed on:

- Each app runs unchanged in its own stack, in a separate internal deployment.
- No database is shared with the venue product.
- Access is private, except Billr's client invoice pages.
- Changes go through extension points, and each fork keeps an `upstream` remote so updates still apply.

## Proposals at a glance

| Area | Proposed tool | Licence (checked 2026-11-18) | When the report says we need it | Our view |
|------|---------------|------------------------------|---------------------------------|----------|
| Sales pipeline (CRM) | Twenty | GitHub reports a mixed licence: AGPL core plus some commercial-licence code | At the first paying venue | Overlaps with Chatwoot labels and Sellbird's pipeline. Decide whether a third tracker is worth it. |
| Accounting and GST | Zoho Books (hosted) | Proprietary | Now | Strong option, but it overlaps heavily with Billr. This is the biggest decision. |
| Subscriptions over UPI | Razorpay Recurring UPI Kit | Apache-2.0, not MIT as the report says | At the first paying venue | Alpha, with 0 stars, and a TypeScript library rather than a service. Too risky as it stands. |
| Error tracking | GlitchTip | MIT (report); main code is on GitLab and active | Now | Reasonable. |
| Uptime and status page | Uptime Kuma | MIT, active | Now | Reasonable, and light. |
| Help center | Chatwoot's built-in help center | Comes with Chatwoot | At 10+ venues | Reasonable. Adds nothing to run. |
| Newsletters | Listmonk | AGPL-3.0, active | At 20+ venues | Reasonable, and light. |
| Booking demo calls | Cal.com | The `calcom/cal.com` repo now redirects to `calcom/cal.diy`, reported as MIT | At 20+ venues | Find out what changed before choosing it. |
| Contracts and e-signature | Documenso | AGPL-3.0, active | At the first paying venue | Check whether a click-to-sign agreement is enough under Indian law. |
| Automation glue | Our Go bridge, not n8n | Ours | Now | Agrees with the plan we already had. |
| Internal dashboards | Metabase, open-source edition | GitHub reports a mixed licence | At 20+ venues | Reasonable later. Heavy, because it runs on Java. |
| Deployments | Coolify | Apache-2.0, active | Now | Conflicts with the shared Postgres in the report's own layout. See below. |
| Passwords | Vaultwarden | AGPL-3.0, not GPL-3.0 as the report says | Now | Reasonable, and light. |
| Private access | Tailscale (hosted) or Headscale | Headscale is BSD-3-Clause | Now | Agrees with our private-access principle. |
| Tasks | Plane | AGPL-3.0, active | Now | About 3 GB of RAM (report). GitHub Projects may be enough for now. |

## Area by area

### Sales pipeline (CRM)

- **Proposal:** Twenty. Chatwoot's webhook creates a deal, and a won deal creates the client in Billr through the bridge.
- **Checked:**
  - Active (last push 2026-11-18).
  - GitHub can't name a single licence for the repo, which usually means an AGPL core plus some code under a commercial licence.
- **Inference:**
  - Twenty also needs Redis, which the report leaves out.
  - Single sign-on may be a paid feature; check before relying on it.
- **Question for us:** With Sellbird tracking prospects and Chatwoot tracking conversations, do we need a third place that tracks the same venues? At low volume, Chatwoot labels plus custom attributes may be enough.

### Accounting and GST

- **Proposal:** Zoho Books, hosted. Billr stays as the client-facing portal, and the bridge copies paid invoices into Zoho Books to produce the official GST invoice.
- **Checked:**
  - Billr has no outgoing webhooks: its API has only `auth`, `health`, `invoices`, `mcp`, `oauth`, `public`, `stripe` and `v1` routes.
  - So the report's "listen to Billr's invoice-paid webhook" step doesn't work. The bridge would have to poll Billr instead.
- **Report:** Zoho Books costs ₹749 a month and handles GST filing.
- **Inference:**
  - Zoho Books also sends invoices, recurring invoices and payment links on its own.
  - If we use it, Billr may be redundant: two systems would hold the same invoices, and the GST invoice would have to win.
  - The real choice is between:
    - (a) Zoho Books for everything, and skip Billr;
    - (b) Billr plus Zoho Books, synced by the bridge;
    - (c) Billr extended with GST, plus an accountant who files from its exports.
- **Question for a CA:** When must Tawn register for GST, and what must an invoice to a venue contain?

### Subscriptions over UPI

- **Proposal:** the Razorpay Recurring UPI Kit (`Questili/razorpay-recurring-upi`), for mandates, renewals, retries and access gating.
- **Checked (README and GitHub API):**
  - Its README calls it "alpha" and "not production-hardened yet".
  - It has 0 stars and 1 fork; it was created 2026-06-22 and last pushed the same day.
  - It's a TypeScript library with a Prisma adapter and a Next.js example. Tawn's API is in Go.
  - Its README says it does no tax or GST.
- **The report gets this wrong:** it calls the kit an MIT-licensed service using under 500 MB of RAM. It's Apache-2.0, and it's a library, not a service.
- **Inference:**
  - Not a fit as it stands.
  - Options to compare:
    - Razorpay's own hosted Subscriptions product, which needs checking for UPI Autopay and retry support;
    - a recurring plan in Zoho Books with a Razorpay payment link;
    - manual UPI or bank transfer with UTR matching, until there are enough venues to justify autopay.
- **Open question from the report:** Will owners in tier 2–4 cities accept UPI Autopay mandates, or do they prefer to pay by static QR or NEFT?

### Errors, uptime and product analytics

- **Proposal:**
  - GlitchTip for errors (Sentry-compatible).
  - Uptime Kuma for uptime and a status page.
  - Plausible or Umami for analytics.
  - Skip PostHog.
- **Checked:**
  - Uptime Kuma is MIT and active.
  - GlitchTip's main backend is on GitLab, active as of 2026-11-18; its GitHub copy is old.
  - Plausible is AGPL-3.0 and Umami is MIT, both active.
- **Report:** PostHog self-hosting now needs at least 16 GB of RAM and 4 vCPUs.
- **Inference:**
  - Plausible and Umami measure website traffic, not in-app product use (funnels, retention, feature use).
  - Product analytics for the venue app is still uncovered.
  - PostHog's hosted free tier may be the practical answer.

### Help center

- **Proposal:** Chatwoot's built-in help center, with a portal per audience.
- **Inference:** Good, because it adds nothing to run. We still need someone to write the articles.

### Marketing

- **Proposal:**
  - Listmonk for newsletters (Go and Postgres, light).
  - Cal.com for booking demo calls.
- **Checked:**
  - Listmonk is AGPL-3.0 and active.
  - `calcom/cal.com` now redirects to `calcom/cal.diy`, which GitHub lists as MIT. We haven't checked what that change means.
- **Inference:**
  - Owners in tier 2–4 cities may not read email newsletters, or book calls through a link.
  - Check how they actually respond during first outreach before adding either tool.

### Contracts and e-signature

- **Proposal:** Documenso. Its webhook moves the CRM stage to "signed".
- **Checked:** AGPL-3.0 and active.
- **Inference:**
  - Under India's IT Act, Aadhaar eSign and digital signature certificates are the formally recognised electronic signatures.
  - A click-to-sign agreement is likely fine for a small subscription, but confirm with a lawyer.
  - Simple online terms accepted at sign-up may be enough at first.

### Automation glue

- **Proposal:** our own Go bridge instead of n8n. n8n's licence isn't OSI open source, and it uses more memory.
- **Inference:** Agrees with our earlier plan. Activepieces is an MIT-licensed alternative if we ever want a visual tool.

### Internal dashboards

- **Proposal:** Metabase, open-source edition, reading a replica of Tawn's database and Billr's database.
- **Checked:** GitHub reports a mixed licence (open-source and enterprise editions).
- **Inference:** Not needed until there's data worth charting. The report estimates about 2 GB of RAM.

### Engineering operations

- **Proposal:**
  - Coolify for deployments, with Traefik and automatic TLS.
  - Vaultwarden for shared passwords.
  - Tailscale or Headscale for private access.
- **Checked:**
  - Coolify is Apache-2.0.
  - Vaultwarden is AGPL-3.0; the report says GPL-3.0.
  - Headscale is BSD-3-Clause.
- **Inference:**
  - Coolify likes to own the server and run each app's database itself. That conflicts with the report's own plan of one shared Postgres and one shared Redis.
  - We'd pick one model: Coolify manages everything, or a plain compose file manages a shared database.
  - Plain compose plus Caddy is the smaller option to compare against.

### Tasks

- **Proposal:** Plane.
- **Checked:** AGPL-3.0 and active.
- **Report:** about 3 GB of RAM.
- **Inference:** For a small team, GitHub Projects next to the repo may do the job, and costs nothing to run.

## Cross-cutting issues

### Overlaps to resolve

| Need | Tools that cover it |
|------|---------------------|
| Pipeline tracking | Sellbird, Chatwoot labels, Twenty |
| Invoicing | Billr, Zoho Books |
| Subscriptions | Billr's recurring flag, Zoho Books, Razorpay Subscriptions |

### Infrastructure

- **Report:** about 12 services, 12–14 GB of RAM in total, on two 8 GB servers.
- **Inference:**
  - Adding only what's needed now (Chatwoot, Billr with Supabase, Sellbird, GlitchTip, Uptime Kuma, Vaultwarden) is much smaller.
  - Postpone the rest until it's needed.

### Sharing Postgres and Redis

- **Report:** one Postgres server with a database per app, and one Redis split by database index.
- **Inference:**
  - One Postgres server with a database per app is fine, though Billr's Supabase stack brings its own Postgres.
  - Sharing one Redis is riskier. Apps need different eviction settings (Sidekiq needs `noeviction`), and one app clearing its data can affect the others.

### Hosting location

- **Report:** Hetzner.
- **Inference:**
  - Hetzner has no India region; its nearest is Singapore.
  - The DPDP Act allows data to leave India except to restricted countries, but it's worth knowing where data sits.
  - Indian providers or AWS Mumbai are alternatives.

### The DPDP Act

The report says self-hosting a CRM "avoids third-party processor liabilities". That's misleading:
- Under the DPDP Act, Tawn is the data fiduciary and is responsible wherever the data sits.
- Hosted processors are allowed, under a contract.
- Venue owners' names and phone numbers are personal data.

The report also suggests encrypting Aadhaar or PAN if we collect them. **Don't collect Aadhaar at all.** The Aadhaar Act restricts who may store it (inference; confirm with a lawyer).

### AGPL

AGPL tools are fine for internal use. If we modify one and let people outside Tawn use it over the network (venues, for example), we must publish our changes. That applies to Sellbird, Listmonk, Documenso, Plane and Vaultwarden.

## Not covered by the report

The report was asked about these and didn't answer, or answered only partly:

- Single sign-on across the internal tools. Authentik and Keycloak are the usual open-source options, and some apps put SSO behind a paid plan.
- Backups: what is backed up, where to, how often, and a tested restore.
- Logs: one place to search all the apps' logs.
- Feature flags for the venue product.
- An internal wiki. `docs/` in this repo may already be the answer.
- Product analytics for the venue app, as noted above.
- Phone support. The report suggests Chatwoot with an Indian provider such as Exotel; we haven't checked whether Chatwoot supports voice calls.

## Decisions to make

1. **Invoicing and accounting:**
   - (a) Zoho Books only;
   - (b) Billr plus Zoho Books;
   - (c) Billr extended with GST.
2. **Collecting subscriptions:**
   - Razorpay Subscriptions;
   - Zoho Books recurring invoices;
   - manual UPI with UTR matching to start.
3. **CRM:** Chatwoot labels and attributes, or Twenty.
4. **Deployment model:** Coolify, or plain compose plus Caddy on one server.
5. **What runs now and what waits** (the report's "now" list versus a smaller start).
6. **Hosting provider and region.**
7. **Contracts:** online terms at sign-up, or Documenso agreements.
8. **Tasks:** Plane, or GitHub Projects.

### Questions for people outside the team

- **CA:** GST registration timing, and what an invoice to a venue must contain.
- **Lawyer:** whether click-to-sign agreements are enough, DPDP consent for outreach, and what never to collect.
- **First venue owners:**
  - whether they read email or only WhatsApp;
  - whether they'll accept UPI Autopay;
  - whether they'll book a demo call or need a visit.

## Decision log

No decisions yet.

| Date | Decision | Why | Who |
|------|----------|-----|-----|

## Sources

- Research report from another agent, received 2026-11-18. Claims labelled **Report** come from it, and it cited these sources:
  - PostHog self-hosting docs
  - Chatwoot features page
  - the Questili repository
  - Frappe and ERPNext docs
  - articles from Web4Cause and TechDino, which we haven't checked
- GitHub API, checked 2026-11-18: licence, last push and archived status for Twenty, Uptime Kuma, Listmonk, Cal.com, Documenso, Metabase, Coolify, Vaultwarden, Plane, Headscale, Chatwoot, Plausible, Umami and the Questili kit.
- GitLab API, checked 2026-11-18: GlitchTip backend activity.
- `Questili/razorpay-recurring-upi` README, read 2026-11-18.
- The Billr, Sellbird, Map and website-launch-skills repositories, cloned and read 2026-11-18.
