# Service des Domaines — foncier

**Procedures covered:** Certificat de Situation Juridique (CSJ), and mutation of
a parcel from one holder to another.

Main lines only. Each section here gets its own pass later.

## Scope note

This covers **titled land** — parcels with a titre foncier held at the
Conservation Foncière. Certificats fonciers issued by the guichets fonciers
communaux are a separate regime and are out of scope for now. Worth revisiting,
but adding a second land regime before the first one works would double the
difficulty for no gain.

## Depends on the commune

Both procedures name people. A parcel belonging to "RAKOTO Jean" is only
unambiguous once that name resolves to a single person. The identity work in
[commune.md](commune.md) is a prerequisite here, not a parallel track.

---

# CSJ — Certificat de Situation Juridique

## What it is

A statement of the current legal situation of a parcel: who holds it, and what
is attached to it — mortgages, oppositions, prénotations, servitudes, anything
that constrains what the holder can do with it.

It is the document you need before buying land, lending against it, or trusting
anyone's claim to it.

## What happens today

You give a parcel reference and you wait months for an answer that already
exists in a register. In the meantime, a sale either stalls or proceeds on
trust — and land bought on trust is how disputes are made.

## What it should be

Give the parcel reference, get its situation. Immediately. Verifiable by whoever
receives it without having to trust the person who handed it over.

## What the system has to hold

- **The parcel** — its title, references and identifying details.
- **The chain of holders** over time, not merely the current one.
- **Encumbrances**, each with when it was inscribed and when it was lifted.
- **The ability to answer as of a date**, not only as of today. A CSJ is a
  point-in-time statement, and a transaction that relied on one must remain
  checkable years later. This means the history is the record and the current
  state is derived from it — never the reverse.

## The hard parts

- Land is contested. The register must be able to say "this is disputed" rather
  than being forced to pick a side it has no business picking.
- Parcels change shape — divided, merged, re-surveyed. Identity of a *parcel* is
  its own problem, separate from identity of a person.
- Existing records are inconsistent, incomplete and sometimes contradictory.
  Digitizing them faithfully means digitizing the contradictions too, and saying
  so, rather than silently choosing a winner.

---

# Mutation — transferring a parcel

## Why this one is different

The other two procedures are reads. This one is a **write**: it changes who owns
something, and it is the point where money, fraud and family disputes all meet.
Making a read faster helps people. Making a write easy without getting it right
would do real harm — a smooth, fast, wrong transfer is worse than a slow one.

Everything about how this is designed should follow from that.

## What happens today

A transfer follows a sale, an inheritance, a donation. It runs through the
Domaines, takes a long time, involves several people who each have an
opportunity to slow it down, and finishes with the register being updated.

The delay is not only inconvenient — while the register is stale, it says
someone owns land they have already sold.

## What it should be

A single event, in person. Seller and buyer both present, each authorising from
their own phone, witnesses attending as they already do, and two fonctionnaires
co-signing. The transfer completes immediately, leaving nothing pending in an
office afterwards — which is exactly where the delay, and the leverage that comes
with it, lives today.

Consent is the hardest thing to prove about a land sale, and every attempt to
prove it on paper is an attempt to reconstruct after the fact something nobody
watched. In a ceremony it is not reconstructed. It is observed, by people and by
the register at the same moment.

See [the ceremony](../architecture/apps.md#the-mutation-ceremony).

## The hard parts

These are the open questions, not answers:

1. **Authorization** — who is entitled to start a mutation on a given parcel.
2. **Consent** — how the seller's agreement is proven, and how that proof is
   verified rather than asserted.
3. **Identity of both parties** — depends directly on the état civil work.
4. **Concurrency** — two mutations touching one parcel, or a sale of land that
   is already under an opposition. The register must be able to refuse.
5. **Fraud and reversal** — what happens when a completed transfer is later
   found to have been fraudulent. The entry cannot be deleted; it has to be
   superseded, with the reason visible. History stays readable.
6. **Money** — droits, taxes and fees attached to a transfer, and the fact that
   payment is part of the procedure rather than separate from it.
7. **Disputes** — a mutation blocked by a court case has to be representable as
   blocked, without the system pretending to adjudicate.

## To confirm

- The exact steps a mutation goes through today, and which office does what.
- Where the delays actually accumulate, as opposed to where people assume.
- What proof of consent is legally sufficient.
- What the fees legally are.
- Whether anything about the current procedure is required by law, or is only
  how the paper process has always been run.
