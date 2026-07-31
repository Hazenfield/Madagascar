# Integrity

How the system prevents a record from being altered — including by whoever runs
it, and including by me.

This is a general layer that applies to the whole application. It is not
specific to land, or to births, or to any procedure added later.

## The claim, narrowly stated

Someone with root on a server can change bytes. That is true of every system,
and any design claiming otherwise is dishonest. So the goal is not prevention:

> **Any modification is either rejected, or provable by anyone — without having
> to trust the operator, including me.**

That is a stronger position than it sounds. It turns "one corrupt person with
database access" into "a conspiracy between many independent parties, at the same
moment". Conspiracies are expensive, slow, and they leak.

## Two layers, kept apart

**The integrity layer** — universal. It sees signed entries and guarantees they
cannot be altered without detection. It does not know what an acte de naissance
is, or what a parcel is.

**The authorization layer** — per-procedure. Who may write what, how many
signatures a given act requires, what evidence must accompany it. This sits on
top and varies by domain.

Keeping them apart is deliberate: adding a new administrative domain must never
require touching the security core. The domain describes its own rules; the
integrity guarantees come for free and stay identical everywhere.

## The integrity layer

**1. The log is the record. The database is not.**
Ownership is not a column that gets updated. It is the result of replaying
entries. Current state is a **projection** — a cache that can be rebuilt from the
log at any time. Someone who edits a projection table directly has poisoned a
cache: it disappears on the next rebuild and contradicts the log immediately. To
forge anything durable they have to attack the log itself.

**2. The log is hash-chained.**
Every entry commits to everything before it. Altering an old entry changes every
hash after it. Tampering is not a surgical edit; it is a rewrite of all
subsequent history.

**3. Checkpoints are signed and published beyond retrieval.**
On a schedule, the head of the log is signed and pushed out — to every commune
node, to public mirrors, anywhere it leaves the operator's control. A rewritten
history then contradicts checkpoints already in other people's hands.

**4. Independent witnesses co-sign.**
Steps 1–3 still leave one hole: if the operator alone publishes checkpoints, the
operator could show one history to Toliara and a different one to Diego. So a
checkpoint is only valid carrying a quorum of signatures from independent
witnesses. This is the layer that actually carries the guarantee.

**5. Every citizen holds a receipt.**
Each issued document carries the record, a proof that this exact record was in
the log at a given position, and the signed checkpoint covering it — printed as a
QR code. Someone holding a paper certificate from 2027 can still prove in 2035
what the register said. Millions of people become witnesses without ever being
recruited, and history cannot be rewritten out from under receipts already in
circulation.

**6. Entries are signed where they are made.**
Each entry carries the signature of the officer who made it. Database access is
not enough to append: it requires a key, which is attributable, revocable, and
noisy when stolen.

## Witnesses

Chosen for divergent interests — the guarantee is only as strong as the
witnesses' independence from each other and from me.

- **Other communes.** Each node witnesses the others. Costs nothing, needs no
  recruitment, grows with adoption. Weak on its own, since all communes sit under
  the same state.
- **Banks and notaries.** They lend against land and certify sales, so they lose
  real money if the register lies. The strongest incentive alignment available.
- **Universities and the press.** Independent of both the state and commerce,
  with a public-interest reason to keep a copy and to look at it.
- **A public blockchain, as one witness among many.** Only a hash is published,
  purely as a timestamp nobody can retract. No token, no smart contract, and the
  system does not depend on it to operate — it is one signature in the quorum,
  and it survives every domestic witness being pressured at once.

Public mirrors run by anyone, including from the diaspora, remain available to
add later.

## Why not build this on a blockchain

It does not address the real failure modes — bad data entered in good faith, and
who was authorized to write. It needs connectivity and running costs that
communes with power cuts do not have. Public chains leak data; private chains are
a slow database with the same trusted operators and extra steps.

The one thing a chain genuinely provides — an unretractable public timestamp — is
what a transparency log with witnesses already provides. That is the same design
that underpins certificate transparency for the whole web's TLS: proven at
planetary scale, and it works offline.

## What this does not solve

To be stated plainly whenever the system is described, because overselling it is
the fastest way to lose the credibility it needs.

- **Bad input.** If an officer signs a transfer backed by forged consent papers,
  the log faithfully records a fraud. This guarantees the *record*, not the
  *truth*.
- **The paper import.** Digitizing decades of registers is exactly where wrong
  data enters. No hash chain makes a wrong register right.
- **Destruction and refusal.** Root can delete everything, or simply decline to
  serve. That is answered by replication and governance, not by tamper-evidence.

## Deferred

- **Per-procedure authorization rules**, including how many signatures a mutation
  requires and what evidence must accompany it. These belong to the authorization
  layer, not here.

## Open

- Quorum size, and how a witness is onboarded or removed.
- How often checkpoints are published.
- How officer keys are issued, stored and revoked.
- What happens when a witness goes quiet — and how that is made visible rather
  than silently tolerated.
