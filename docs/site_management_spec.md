# SiteSathi — Site Management Spec

## Scope
This spec lists everything SiteSathi must handle to run one construction site, from setup to handover: who enters what, what is saved, and what the app checks automatically.

* **In scope:** the 17 site processes for 4 roles, broken into 33 steps below — supervisor, civil engineer, architect, owner. These are the same processes as the design canvas.
* **Deferred:** remote site visits by local firms (the freelancer network). It returns after the core app is in use.
* **New layer, inspired by FlowManual:** *smart checks*. FlowManual is a Y Combinator start-up that reads contracts, quotes and invoices to catch money leaks for US contractors. SiteSathi does the same with data the site already captures — bills against orders and deliveries, contractor bills against work orders, and material use against standard norms.
* **Guiding rule:** one entry, many uses. A supervisor's photo, voice note or count fills the paper register it replaces, the owner's update and the checks. Nobody types the same thing twice.

Each process section below uses the same columns: what the record is, who enters it, what is captured, and the rules and checks that apply.

## The paper registers this replaces
CPWD's guidelines list 15 records to keep at a site. Private sites keep fewer, often in notebooks and WhatsApp. SiteSathi builds each one from daily entries, so the register exists without anyone writing it.

| Register (paper today) | What it records | Filled in SiteSathi by |
| :--- | :--- | :--- |
| Daily progress record | Work done each day | Supervisor's voice daily report (DPR) |
| Site engineer's field book | Soil, material and workmanship notes | Engineer's comments on reports and checks |
| Site order book | Instructions to the contractor, defects found | Engineer and architect instructions; quality “Fix” results |
| Hindrance register | What stopped work, and for how long | Problems marked “work stopped”, rain days from reports |
| Cement register | Cement received and used vs theoretical need | Material receipts + work done × consumption norm |
| Steel register | Steel used vs drawing requirement | Steel receipts + bar-bending schedule from drawings |
| Measurement book (MB) | Measured work for contractor bills | Engineer's measurements on running bills |
| Tests record | Cube tests, material tests | Quality module test log |
| Non-conformance record | Sub-standard work and its fix | Quality checks marked “Fix” and their closing photos |
| Periodic inspection record| Inspections against drawings and specs | Engineer checks and architect site visits |
| Drawings record | Which Good-for-Construction (GFC) drawing is current | Drawing register with revisions and “seen by site” |
| Coordination meeting minutes | Decisions and actions | Meeting note (voice) with action items |
| Labour muster roll and wage register | Attendance and wages | Attendance + labour bills (also required by the labour codes) |
| Guarantee bonds and insurance | Bank guarantees, contractor insurance | Document vault with expiry reminders |
| Paint and chemicals record | Use of paint, waterproofing chemicals | Material module, same as cement |

---

## A · Before work starts
Setup decides how good every later check is: a BOQ, work orders and payment stages entered once let the app check every bill and delivery that follows.

| Process | Who enters | What is captured | Rules and checks |
| :--- | :--- | :--- | :--- |
| **1 · Project and team** | Owner or engineer | Name, address and location pin, plot size, floors, type, budget, start date; team invited by phone with a role each | One phone number = one person across all projects; roles decide what each person sees and can approve |
| **2 · Site survey** | Supervisor, engineer | Checklist photos (plot corners, neighbour walls, road, trees), levels, soil test, bench mark, water and power points | Photos dated and located; becomes the “before” record for disputes |
| **3 · Drawings** | Architect, engineer | Sheet number, title, discipline (architectural, structural, plumbing, electrical), revision, status, what changed | Site sees only the latest GFC revision; old ones are hidden; site must tap “seen” on each new revision |
| **4 · BOQ and budget** | Engineer (Excel import or template) | Item, unit, quantity, rate, amount, floor or area, linked spec from architect | Every later material request, bill and change links to a BOQ item; owner sees it grouped into simple stages |
| **5 · Schedule** | Engineer | Stages, start and end dates, dependencies, owner decision dates (tile choice), architect drawing dates | Delay reasons pulled from the hindrance register; owner sees one forecast date |
| **6 · Work orders** | Engineer | Contractor (thekedar or trade), scope, BOQ items covered, rates, advance, retention %, TDS applicable, payment terms | Every running bill is checked against its work order (see *Smart checks*) |
| **7 · Payment stages** | Owner and engineer | Stage name, amount, what must be done and checked before payment | Payment unlocks only after the stage is marked done and checked; matches bank loan stages if the owner has a home loan |
| **8 · Vendors** | Engineer | Supplier name, GSTIN, phone, items supplied, agreed rates | Rates stored here are what bills are compared against |

---

## B · Every day on site
The supervisor's day is four must-dos — safety talk, attendance, receiving material, daily report — and every other register is built from them.

| Process | Who enters | What is captured | Rules and checks |
| :--- | :--- | :--- | :--- |
| **9 · Morning safety talk** | Supervisor | Topic played (local language), gear check (helmet, shoes, belt, gloves), group photo | Group photo also counts as attendance proof; missing talk shows on engineer's safety summary |
| **10 · Attendance and wages** | Supervisor | Count per trade per contractor, overtime, photo, location, time | Builds muster roll and wage register; weekly labour bill = days × work-order rate − advances; days without a photo are flagged before the engineer certifies |
| **11 · Material request** | Supervisor | BOQ item, quantity, needed-by date, voice note | Shows quantity left in BOQ; above-plan requests need engineer approval, large ones need owner approval |
| **12 · Purchase order** | Engineer | Vendor, item, quantity, agreed rate, GST, delivery date | Rate checked against vendor's agreed rate and last 3 purchases |
| **13 · Receive material (GRN)** | Supervisor | Quantity received, condition (good, torn, wet), challan photo, truck and number plate photo, e-way bill number if any | Short or damaged delivery tells the engineer and supplier at once; received quantity is what the supplier bill is checked against |
| **14 · Material use and stock** | Auto + supervisor | Stock = received − used; used = work done × consumption norm; weekly physical count by supervisor | Cement and steel registers; gap between expected and counted stock is flagged |
| **15 · Daily report (DPR)** | Supervisor (voice) | Work done with quantity and location, workers (auto), material in (auto), machines (auto), problems, weather, photos | Due by a set time with reminder; engineer corrects quantities and links them to BOQ items |
| **16 · Petty cash** | Supervisor | Bill photo, amount, category, voice note | Balance shown; engineer approves; owner tops up by UPI |
| **17 · Quality check** | Supervisor asks, engineer decides | Stage (slab steel, column steel, pipes before plaster), one photo per checklist point with tape in frame, optional video call | Next activity (pour, plaster) is blocked until engineer marks OK; “Fix” items go to the site order book |
| **18 · Tests** | Engineer, supervisor | Cube casting date, 7- and 28-day results, other material tests | Reminders on test days; failed result raises a non-conformance record |
| **19 · Site problems** | Supervisor | Type (drawing doubt, work problem, owner wants a change, safety), photo, voice, pin on drawing, “work stopped” flag | Sent to the right person by type; “work stopped” enters the hindrance register with time lost |
| **20 · Machines and fuel** | Supervisor | Start/stop with meter photos, breakdown time, diesel given with slip photo | Hired machine bill = hours × rate; checked against meter photos |
| **21 · Instructions and meetings** | Engineer, architect | Instruction text or voice, who must act, due date; meeting notes with actions | Site order book; open instructions shown on supervisor's Today |

---

## C · Money and changes
No rupee leaves without proof: every payment shows the photos, counts or measurements behind it, and the engineer's check, before the owner taps Pay.

| Process | Who enters | What is captured | Rules and checks |
| :--- | :--- | :--- | :--- |
| **22 · Supplier bills** | Engineer or owner (photo of GST invoice) | Supplier GSTIN, invoice number and date, items, HSN, quantity, rate, GST, total | Matched to purchase order and goods received (see *Smart checks*); duplicate invoice numbers blocked |
| **23 · Contractor running bills** | Contractor submits, engineer measures | Measurements (length × height × thickness, or count) by location, BOQ item, rate, previous bills, advance recovery, retention, TDS | Checked against work order rates and BOQ quantities; total-to-date can't exceed BOQ without a change order |
| **24 · Labour bills** | Auto from attendance | Days per trade, rate, overtime, advances | Certified by engineer; proof of wage payment kept (principal employer liability) |
| **25 · Machine and other bills** | Auto from machine log | Hours, rate, fuel | Compared with meter photos |
| **26 · Change orders** | Supervisor flags, engineer prices, architect revises drawing, owner approves | Description, reason, cost lines from BOQ rates, extra days, new budget, new finish date, linked drawing revision | Work can't start before owner approval; approved changes update BOQ, schedule and the contractor's work order |
| **27 · Owner stage payments** | Engineer marks stage done, owner pays | Stage, amount, proof (photos, checks), UPI reference | Unlocked only after the stage's quality check is OK |
| **28 · Compliance packs** | Auto, engineer and architect sign | RERA quarterly pack (floor-wise photos, % complete, Form 1 and 2 drafts), labour registers, cess estimate | Reminders before due dates; for RERA-registered projects only |
| **29 · Cash flow view** | Auto | Paid, due this week, due at next stage, retention held | Owner and engineer see the same numbers |

---

## D · Finishing, handover and after
The site record doesn't end at handover: it becomes the owner's house folder and the evidence for the 5-year defect period.

| Process | Who enters | What is captured | Rules and checks |
| :--- | :--- | :--- | :--- |
| **30 · Snag list** | Architect, engineer, owner raise; supervisor fixes | Defect, room, pin on plan, before photo, assigned trade, after photo | Contractor's retention is held until snags close; “fixed” needs an after photo and a recheck |
| **31 · Handover** | Engineer, owner | Final walk report, as-built drawings, warranties (tiles, pumps, paint, waterproofing), all bills, stage photos, test results, completion certificate | Owner accepts in the app; first half of retention can be released |
| **32 · Defect liability (after move-in)** | Owner or buyer raises; builder routes | Complaint photo, location, category; linked to the trade and stage photos that built it | For RERA projects the builder must fix structural and workmanship defects free within 30 days, for 5 years after possession; the app runs the 30-day timer |
| **33 · Retention release** | Engineer | Defect period end date, open complaints | Reminder on the release date; the contractor must claim it in writing, so the app drafts the claim |

---

## Smart checks (inspired by FlowManual)
FlowManual reads contracts, quotes and invoices to flag scope changes and invoice drift. SiteSathi can go further, because it also knows what physically arrived and what was built. Most checks are plain rules on data already captured. Only reading photos and PDFs of bills, quotes and work orders needs AI.

*The bill was ₹19,000 (50 × ₹380); only ₹17,520 matches what was ordered and received, so the gap is held until the supplier explains it.*

| Check | Compares | Example flag | Needs AI? |
| :--- | :--- | :--- | :--- |
| **Bill vs order vs received** | Supplier invoice ↔ purchase order ↔ goods received note | “Billed 50 bags, received 48 — hold ₹760”; “Rate ₹380, agreed ₹365”; same invoice number twice | Reading the invoice photo |
| **Running bill vs work order** | Contractor's bill ↔ work-order rates ↔ BOQ quantity | “Brickwork to date 405 m³, BOQ 384 m³ — needs a change order”; item not in work order; retention or advance not deducted | Reading the work order once |
| **Labour bill vs proof** | Days billed ↔ attendance photos ↔ work done | “Saturday has no group photo”; “Bricks laid per mason fell 40% this week” | No |
| **Material use vs norm** | Received − stock ↔ work done × standard consumption | M20 concrete needs about 8 bags/m³ and M15 about 6.3–6.5. Flag use above norm + 2–5% wastage | No |
| **Quote comparison** | 3 supplier or contractor quotes, usually sent on WhatsApp | Lines up rate per unit with GST, transport, unloading, credit days and delivery date; picks the lowest true cost | Reading quote photos and PDFs |
| **Drawing change impact** | New revision ↔ old revision ↔ work already done | “Window W3 changed after its lintel was cast — likely change order” | Comparing drawings (later phase) |
| **Work order scope** | Work order text ↔ bills | “Scaffolding is excluded in the work order but billed”; missing retention or defect clause | Reading the work order |
| **Machine bill vs meter** | Hours billed ↔ start/stop meter photos | “Billed 12 h, meter shows 9.5 h” | No |
| **Petty cash patterns** | Spends ↔ bills ↔ history | Same amount every day, spends with no bill photo, spends on holidays | No |
| **Order on time** | Needed-by date ↔ vendor lead time ↔ schedule | “Order steel today or the slab slips 2 days” | No |

**Build order:** the rule-only checks first (labour, material norms, machines, petty cash, order timing), then bill vs order vs received, then reading work orders and quotes.

---

## Indian tax and legal rules the app must follow
These rules change; build them as settings the team can update, not as fixed code, and have a CA review before launch.

| Rule | What it means for the app |
| :--- | :--- |
| **TDS on contractor payments** | From 1 April 2026, old Section 194C sits under Section 393(1) of the Income-tax Act 2025. Rate 1% for individual/HUF contractors, 2% for others, 20% without PAN. Applies when one payment exceeds ₹30,000 or the year's total exceeds ₹1,00,000. The app deducts it on running bills and keeps PAN per contractor |
| **GST invoices** | Store supplier GSTIN, invoice number, HSN, taxable value and GST per line; block duplicates; keep for input credit |
| **E-way bill** | Needed when goods worth over ₹50,000 move (₹1,00,000 intra-state in Delhi and West Bengal, ₹2,00,000 in Bihar). Capture the e-way bill number at goods receipt for large deliveries |
| **Retention** | Usually 5–10% of each running bill (5% common in CPWD contracts), often half released at handover and half after the defect period |
| **Labour codes** | In force since 21 November 2025; registers cut from 84 to 8; the principal employer pays if the contractor doesn't. App keeps muster roll, wage, deduction and overtime registers and wage-payment proof |
| **Building workers' cess** | 1–2% of construction cost; workers to be registered with the welfare board |
| **RERA quarterly update** | Registered projects upload progress with floor-wise photos each quarter; many states also need architect, engineer and CA certificates |
| **RERA defect liability** | 5 years from possession; fix within 30 days, free |

---

## Core data the app stores
About 25 records cover the whole site; every one is created on the phone with its own ID so it can be saved offline and synced later without duplicates.

| Record | Key fields | Links to |
| :--- | :--- | :--- |
| **Person** | Phone, name, language | Projects (with a role in each) |
| **Project** | Name, address, pin, type, budget, dates | Everything below |
| **Drawing revision** | Sheet, revision, status, file, what changed, seen-by | Issues, change orders |
| **BOQ item** | Code, description, unit, quantity, rate, floor or area, spec | Requests, work orders, measurements, changes |
| **Schedule stage** | Name, planned and actual dates, depends on | Payment stages, delays |
| **Vendor / Contractor** | Name, GSTIN or PAN, phone, trade, agreed rates | Purchase orders, work orders |
| **Work order** | Contractor, BOQ items, rates, advance, retention %, TDS | Running bills, attendance |
| **Material request → Purchase order → Goods receipt → Supplier invoice** | Item, quantity, rate, dates, photos, GST lines | Each to the next; BOQ item; checks |
| **Stock count** | Item, counted quantity, date | Cement and steel registers |
| **Attendance** | Date, contractor, trade, count, photo | Labour bill, muster roll |
| **Daily report** | Date, work done lines, weather, voice and transcript | Attendance, receipts, machine logs, issues, photos |
| **Photo** | File, time, GPS, device, mock-location flag | Any record above |
| **Issue / Instruction** | Type, text or voice, pin on drawing, owner, due, work-stopped time | Hindrance register, change orders |
| **Quality check / Test** | Stage, checklist results, photos, decision; test dates and results | Payment stages, non-conformance |
| **Machine log / Fuel** | Machine, start, stop, meter photos, litres | Machine bills |
| **Petty cash entry** | Amount, category, bill photo, approval | Cash book |
| **Running bill** | Measurements, quantities, deductions, TDS, status | Work order, BOQ items, payments |
| **Change order** | Cost lines, days, approvals | BOQ, schedule, work order, drawing |
| **Payment** | Type (stage, bill, labour, top-up), amount, UPI reference | Bills, stages |
| **Snag / Warranty complaint** | Room, pin, photos, trade, timer | Handover, retention |
| **Check flag** | Check type, records compared, amount at risk, status | The records it flags |
| **Audit log** | Who, what, when, before and after | Every change, for disputes |
