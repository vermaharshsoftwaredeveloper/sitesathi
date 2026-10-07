# SiteSathi — Application Architecture

## Key decisions

One Flutter app for all four roles talks to one Go service on Postgres in AWS Mumbai.
Every action saves on the phone first and syncs when the network allows. It implements
the Site Management Spec and the design canvas.

| Area | Choice | Why |
| :--- | :--- | :--- |
| Mobile app | Flutter, one app; role decides the home screen | One codebase for Android and iPhone; smooth on budget Androids |
| State | Riverpod | Simple, testable, works well with streams from the local database |
| Local database | Drift (SQLite) | Typed queries; app works fully offline; reactive lists |
| Sync | Own outbox and change feed (below); PowerSync as a fallback option | No extra vendor; Go stays the single place where rules run |
| Backend | Go modular monolith, chi router, sqlc + pgx | One deployable; fast; clear module borders to split later if needed |
| Database | PostgreSQL | Relational site data, strong transactions, also runs the job queue |
| Background jobs| River (jobs stored in Postgres) | Jobs commit in the same transaction as the data, so none get lost |
| Files | S3 in Mumbai, direct upload with signed links | Photos never pass through the Go server |
| Login | Phone OTP, then short-lived access token + refresh token | No passwords for supervisors |
| Hosting | AWS Mumbai (ap-south-1) | Data stays in India; good network from Indian carriers |
| Web for engineers (later) | Same Go API; Flutter web first, React only if large tables feel slow | Nothing in the backend changes |

## Free pilot setup (₹0)

For the pilot (5–10 sites, about 50 users, Android only) the same Go and Flutter code runs on free services: one Oracle Always Free server in India, Cloudflare R2 for photos and Firebase's free tools. Moving to the paid AWS setup later is a configuration change, not a rewrite.

### What changes for the pilot

| Piece | Paid plan (above and below) | Free pilot |
| :--- | :--- | :--- |
| Server | AWS ECS containers | One Oracle Always Free ARM VM running Docker Compose |
| Database | Amazon RDS Postgres | Postgres in a container on the same VM |
| Background jobs| River workers | Same River workers, same VM |
| Photos, audio, backups | S3 + CloudFront | Cloudflare R2 (S3-compatible, so same code): 10 GB, free downloads |
| HTTPS | Load balancer certificate | Caddy web server with a free Let's Encrypt certificate |
| Login | Phone OTP by SMS | Invite code shared on WhatsApp by the engineer, or Google sign-in; SMS OTP costs money |
| Voice to text | Sarvam AI on the server | Phone's built-in speech recognition, plus the audio saved for later |
| Push, crash reports, test builds | FCM, Sentry, Play Store | FCM, Crashlytics and App Distribution — free on every Firebase plan |
| WhatsApp messages| WhatsApp API | “Share to WhatsApp” button that opens WhatsApp with the text filled in |
| UPI payments | Razorpay | Owner pays as usual and types the UPI reference |
| Bill reading with AI | Vision AI model | Left out; rule-only smart checks still run |

### The Oracle server

* Pick **Mumbai or Hyderabad as the home region** at sign-up. Always Free servers and disks only work in the home region, and it can't be changed later.
* Free ARM allowance is 2 CPUs and 12 GB RAM in total, with 200 GB disk and 10 TB outbound data a month.
* **Avoid losing the server to idle reclaim.** Oracle reclaims a free server if, over 7 days, CPU, network *and* memory all stay under 20%. A pilot will be quiet, so make 1 CPU / 6 GB the VM size and give Postgres about 1.5 GB of cache; memory then stays above 20%.
* Containers in Docker Compose: `caddy` (HTTPS), `api` (Go), `worker` (Go, River), `postgres`. One `docker compose up` brings it all back.
* **Backups off the server:** a River job runs `pg_dump` every night, compresses it and stores it in R2 for 14 days. Test a restore once before the pilot starts.

### Voice in the pilot

Android's built-in recognizer ends a session after about 5 seconds of silence, and the app can't change that. So the daily report asks short questions one at a time, each answered with a hold-to-speak button: “What work was done?”, “Any problem?”. The audio is also recorded and uploaded, so it can be re-transcribed with Sarvam once there's budget.

### Code rules that keep the move to paid easy

* All settings come from environment variables: database URL, storage URL and keys, login methods, speech provider.
* Storage uses the S3 API only, so R2 now and S3 later need no code change.
* Login, speech, messaging and payments each sit behind a small Go interface, with a “free” and a “paid” implementation.

### When to leave the free setup

| Trigger | Move to |
| :--- | :--- |
| R2 storage nears 10 GB (around 30,000 compressed photos) | Paid R2 or S3 |
| More than about 20 active sites, or the API feels slow | AWS Mumbai setup in this doc |
| First paying customers, or a customer asks about uptime and backups | Managed database with point-in-time restore |
| Users ask for phone OTP or iPhone | SMS OTP provider; Apple developer account |

### Setup checklist

1. Create the Oracle Cloud account with Mumbai or Hyderabad as home region; create a 1 CPU / 6 GB ARM VM with Ubuntu.
2. Point a free subdomain at the server's IP (a free dynamic-DNS name, or a cheap domain later).
3. Install Docker; deploy the Compose file; Caddy fetches the certificate.
4. Create the Cloudflare R2 bucket and keys; set a usage alert.
5. Create the Firebase project: Cloud Messaging, Crashlytics, App Distribution.
6. Build the Android app; add pilot users as testers in App Distribution.
7. Turn on the nightly backup job and test one restore.

---

## System overview

The phone talks to the Go API only for sync and login; photos and audio go straight to S3, and everything slow (voice, checks, PDFs) runs in the workers.

## Mobile app (Flutter)

The app is organised by feature, and every screen reads from the local database, never straight from the network. That is what makes it work on a site with no signal.

### Folder layout

```text
lib/
  app/       router, theme, languages (ARB files: hi, en, mr, gu, te, ta, kn, bn)
  core/
    db/      Drift database, tables, migrations
    sync/    outbox, change pull, retry, conflict handling
    api/     HTTP client, auth tokens, refresh
    media/   in-app camera, compression, background upload queue
    voice/   recorder, upload, transcript status
    location/ GPS with mock-location check
    auth/    OTP login, session, role per project
  features/
    today/ safety/ attendance/ report/ material/ petty_cash/
    quality/ issues/ drawings/ machines/ snags/
    owner_home/ approvals/ payments/ timeline/
    boq/ schedule/ bills/ labour/ change_orders/ compliance/
    each: data/ (tables, DAO, repository) · ui/ (screens, widgets) · providers.dart
```

### Main packages

| Need | Package | Note |
| :--- | :--- | :--- |
| Navigation | go_router | One shell per role: supervisor, engineer, architect, owner |
| State | flutter_riverpod | Providers watch Drift queries, so screens update when sync lands |
| Local DB | drift | All records, plus the outbox table |
| HTTP | dio | Token refresh, retries |
| Camera | camera | In-app capture only, no gallery, so every photo is fresh |
| Compression | flutter_image_compress | Target about 200–400 KB per photo before upload |
| Voice | record | Records AAC on the phone; works offline |
| Location | geolocator | Reads the mock-location flag on each fix |
| Background work | workmanager | Upload photos and sync when the app is closed |
| Push | firebase_messaging | Approvals, reminders, alerts |
| Drawings | a PDF viewer package | Pins drawn on top as an overlay |
| Crashes | sentry_flutter | Errors from real phones in the field |

**Roles:** after login the server returns the user's projects with a role in each. The app picks the shell for that role. The server checks the role again on every request, so hiding a button is never the only protection.

**Voice reports:** the phone records and saves the audio, then uploads it like a photo. The server transcribes it in a background job and sends back the text and filled fields. Offline, the report shows “will be transcribed when online”.

### Offline-first sync

The phone is the first place every record is saved. A small outbox pushes changes up in order, and a change feed pulls other people's changes down. Every record gets its ID on the phone (UUID v7), so a retried upload can never create a duplicate.

**Push (phone → server)**
1. A save writes the record to Drift and adds an operation to the outbox in the same local transaction: operation ID, table, record ID, create/update/delete, changed fields, the version the phone last saw.
2. When online, the sync engine sends outbox operations in order, in batches: `POST /v1/sync/push`.
3. The Go server handles each batch in one database transaction. It skips operations it has already applied, checks the user's role, runs the business rules, saves, writes a change-log row and queues any jobs (River) in that same transaction.
4. Each operation comes back as *applied*, *rejected* (with a reason shown to the user) or *conflict*. Applied operations leave the outbox.

**Pull (server → phone)**
* Every saved change gets a sequence number per project. The phone asks `GET /v1/sync/pull?project=…&since=<last number>` and stores what comes back, page by page, until it is up to date.
* A push notification (“project has changes”) triggers a pull, so updates arrive quickly without constant polling.
* The phone only receives projects it belongs to, and only what its role may see. For example, a supervisor gets no owner payment details.

**Conflict rules**
| Kind of record | Examples | Rule |
| :--- | :--- | :--- |
| Add-only | Photos, attendance, goods received, daily reports, machine logs | Never conflict; corrections are new entries that point at the old one |
| Plain editable fields | Task names, notes, BOQ descriptions | Later save wins per field; the audit log keeps both |
| Decisions and money | Approvals, quality OK, bill certify, payments | Server decides. The phone sends an intent (“approve”); if the state already moved (say, already paid), it is rejected with a clear message |
| Deletes | Any | Soft delete (`deleted_at`), synced like any change |

**Photos and audio** sync in two parts: the record goes through the outbox at once with status *uploading*. The file uploads separately in the background straight to storage, then the record is marked *uploaded*. The phone's clock time is kept as evidence next to the server's time, and a large gap between them is flagged.

If time is short: **PowerSync** offers the same model as a service. Its client queues changes and uploads them to your backend, while its service streams Postgres changes down to the phone's SQLite. It fits behind the same Go rules, at the cost of one more service to run or pay for.

---

## Backend (Go)

Start as one Go service with clear modules: the API and the background workers share one codebase and one database. Split a module out only when it truly needs its own scaling.

### Folder layout

```text
cmd/
  api/          HTTP server
  worker/       River background workers (same code, different entry)
internal/
  platform/     config, Postgres pool (pgx), S3, push, SMS, logging, tracing
  auth/         OTP, tokens, devices
  access/       roles and permissions per project
  sync/         push and pull endpoints, change log
  projects/     projects, members, invites
  drawings/ boq/ schedule/
  material/     requests, purchase orders, goods received, stock, vendors
  labour/       attendance, labour bills, registers
  reports/      daily reports, voice transcription
  quality/      checks, tests
  issues/       problems, instructions, hindrance
  machines/
  money/        running bills, supplier bills, petty cash, payments, TDS
  changes/      change orders
  handover/     snags, warranty complaints
  compliance/   RERA packs, register exports
  checks/       smart-check engine (rules + bill and quote readers)
  notify/       push, WhatsApp, SMS templates
db/
  migrations/   schema changes (goose or golang-migrate)
  queries/      SQL per module, turned into Go code by sqlc
```

Each module has the same three layers: **handler** (HTTP in and out), **service** (business rules and permission checks), **store** (sqlc queries). Modules call each other's services, never each other's tables.

### API
* REST + JSON under `/v1`. Most writes from the phone arrive through `/v1/sync/push`. Separate endpoints cover login, signed upload links, exports, and payment webhooks.
* An OpenAPI file describes every endpoint. The Dart client in the app is generated from it, so app and server never drift apart.

### Login and roles
* Phone + OTP, then a 15-minute access token and a long-lived refresh token tied to the device. The refresh token is stored hashed and replaced on each use.
* Roles per project: supervisor, engineer, architect, owner (contractor and viewer later). The permission table lives in `access/` and every service checks it.

### Background jobs (River)
Jobs are saved in Postgres in the same transaction as the data that caused them, so a job is never lost or run for data that didn't save.

| Job | Triggered by |
| :--- | :--- |
| Transcribe voice report, fill fields | Audio uploaded |
| Make thumbnails, read photo location | Photo uploaded |
| Run smart checks | Goods received, supplier bill, running bill, attendance week closed |
| Build registers and PDFs | Weekly, or on request |
| Reminders (report due, test day, RERA date) | Scheduled (periodic jobs) |
| Send push, WhatsApp, SMS | Any of the above |

**Smart-check engine:** when a record that a check depends on is saved, a job runs the matching rules and writes a *check flag* (amount at risk, records compared, status). Norms and limits, such as cement per m³ or wastage %, live in a settings table per project, not in code. Bill, quote and work-order photos are read by an AI step into structured fields with a confidence score. A person confirms them before any check uses them.

### Data and storage
Every table follows the same few conventions, so sync, permissions and the audit log work the same way everywhere.

**Every synced table has:**
| Column | Purpose |
| :--- | :--- |
| `id` uuid | (v7, made on the phone) Same ID offline and online; sorts by time |
| `project_id` uuid | Every query is scoped to a project the user belongs to |
| `created_by`, `created_at`, `updated_at` | Who and when (server time) |
| `client_time` | Phone time, kept as evidence |
| `version` int | Bumped on each change; used to detect conflicts |
| `deleted_at` | Soft delete, synced like any change |

**Key tables:**
* **People and access:** `users`, `devices`, `projects`, `project_members` (role)
* **Plan:** `drawings`, `drawing_revisions`, `boq_items`, `schedule_stages`, `payment_stages`, `vendors`, `contractors`, `work_orders`
* **Site:** `attendance`, `daily_reports`, `material_requests`, `purchase_orders`, `goods_receipts`, `stock_counts`, `quality_checks`, `tests`, `issues`, `instructions`, `machine_logs`, `petty_cash`
* **Money:** `supplier_invoices`, `running_bills`, `bill_measurements`, `change_orders`, `payments`
* **Finish:** `snags`, `warranty_complaints`
* **Platform:** `media`, `check_flags`, `change_log` (project_id, seq), `sync_ops` (applied operation IDs), `audit_log`, `settings_norms`, River's job tables

Money is stored in paise as whole numbers (`bigint`), never as decimals in floating point. Quantities use `numeric` with the unit stored beside them.

**Photos, audio, drawings**
1. The app asks the API for a signed upload link (`POST /v1/media/upload-url`).
2. The phone uploads the compressed file straight to S3 (Mumbai), in the background.
3. S3 tells the server the file arrived; a job makes a thumbnail and marks the `media` row uploaded.
4. Files are read back through short-lived signed links on a CDN. No file is public.

### Outside services
Each outside service sits behind a small Go interface in `platform/` or `notify/`, so a provider can be swapped without touching business code.

| Need | First choice | Backup | Notes |
| :--- | :--- | :--- | :--- |
| OTP by SMS | MSG91 | WhatsApp OTP | Indian DLT-registered templates needed for SMS |
| WhatsApp messages | Meta WhatsApp Cloud API | SMS | Invites, approvals, supplier “2 bags short” notes |
| Push | Firebase Cloud Messaging | — | Also used to trigger a sync pull |
| UPI payments | Razorpay or Cashfree | — | Owner pays bills and stages; webhooks confirm |
| Speech-to-text | Sarvam AI | Google Speech, Bhashini | 22 Indian languages plus English, with translate-to-English mode. |
| Reading bills, quotes, work orders | A vision-capable AI model | Manual entry | Output as fields with confidence; a person confirms; no personal data sent beyond the document |
| Maps | Google Maps SDK | — | Site pin and distance checks only |

### Security, privacy and anti-fraud
The app holds people's money records, workers' details and site photos, so trust is the product: every proof must be hard to fake and every change traceable.

**Access**
* Every API call checks the user's role *in that project* on the server; the app hiding a button is never the only guard.
* Each project's data is fetched only with its `project_id` filter. Tests try to read other projects' data on purpose.
* Tokens are short-lived; refresh tokens are tied to a device and can be revoked from the web (“log out lost phone”).

**Data**
* All data and files stay in AWS Mumbai. Encrypted on disk and over the network.
* Collect only what a feature needs. Workers' phone numbers and IDs are visible only to the engineer and owner.
* Plan for India's data protection law (DPDP Act): consent at sign-up, a way to export and delete personal data, and a breach plan. Get legal review before launch.

**Anti-fraud (proof that site data is real)**
| Risk | Guard |
| :--- | :--- |
| Old or downloaded photo | In-app camera only; server time stamped; photo hash stored to catch re-use |
| Fake GPS | Android's mock-location flag checked on every fix; impossible speed between fixes flagged |
| Wrong person marking attendance | Group photo; optional face check later |
| Edited records after approval | Approved records lock; any change makes a new version and an audit-log row |
| Rooted phones and emulators | Device integrity check (Play Integrity on Android); flagged, not blocked, at first |

**Audit log:** who changed what, when, before and after, for every money and approval action. It is append-only and kept for the life of the project plus the 5-year defect period.

### Infrastructure, deployment and monitoring
Keep the first setup small and managed: two containers, one database, one bucket, all in AWS Mumbai.

| Piece | Service | Start with |
| :--- | :--- | :--- |
| API and workers | Containers on AWS ECS Fargate behind a load balancer | 2 small API tasks, 1 worker task |
| Database | Amazon RDS for PostgreSQL | One instance, automatic backups with point-in-time restore; add a standby when paying customers arrive |
| Files | S3 + CloudFront | Versioning on; old originals move to cheaper storage after 90 days |
| Secrets | AWS Secrets Manager | API keys for SMS, WhatsApp, payments, speech |
| Environments| dev, staging, production | Staging uses copies of real flows with test data |

**Release pipeline**
* **Backend:** GitHub Actions runs tests and `sqlc` checks, builds a container and deploys to staging, then to production after approval. Database migrations run before the new code.
* **App:** GitHub Actions or Codemagic builds Android and iOS. Releases go to the Play Store internal track first, then 10% of users, then everyone. Old app versions keep working because the API is versioned (`/v1`).

**Monitoring:** Sentry for crashes (app and server), plus OpenTelemetry traces and metrics to Grafana or CloudWatch. Watch the numbers that show site pain:
* Outbox size per phone, and time from save to synced
* Failed photo uploads
* Voice transcription time and failure rate
* Daily reports submitted vs expected
* Smart-check flags raised and resolved

### Build phases and team
Build the base that every feature stands on first — login, roles, sync and photo upload — then the supervisor's daily loop. Move to the next phase only when the gate is met on real sites.

| Phase | Backend (Go) | App (Flutter) | Gate to move on |
| :--- | :--- | :--- | :--- |
| 0 · Foundation | OTP login, projects, members, roles, sync push/pull, change log, media upload links, audit log | Login, language, project switch, Drift schema, outbox, background upload | A test phone works a full day offline and syncs with no lost or doubled records |
| 1 · Daily site loop | Attendance, safety, material request and receipt, daily report + transcription, petty cash, drawings, issues | Supervisor Today and its 9 actions; owner home and updates | 3–5 pilot sites send a daily report for 4 weeks |
| 2 · Money and quality | BOQ import, purchase orders, supplier and running bills, labour bills, quality checks and tests, change orders, UPI payments, rule-only smart checks | Engineer screens on phone; owner approvals and payments | Owners approve and pay stages in the app |
| 3 · Smart and compliance | Bill, quote and work-order reading; bill vs order vs received; registers export; RERA pack; snags, handover, warranty | Snags, handover folder; engineer web dashboard | First paying builders using the checks |

**Suggested team for phases 0–2:** 2 Flutter developers, 2 Go developers (one strong in Postgres and sync), a part-time designer, and one tester who visits sites weekly with a cheap Android phone. The founder acts as product owner and talks to supervisors every week.
