# Administrative domains

The administrative bodies this project covers, and the papers handled for each.
I take them one at a time, and each gets designed before anything is built.

Deliberately small to begin with: two bodies, three procedures — the three that
cost me the most time and patience.

| Body | Procedures | Kind | Current |
|---|---|---|---|
| Commune — état civil | Acte de naissance | read | [v0.1](commune/v0.1.md) |
| Service des Domaines | CSJ · Mutation d'une parcelle | read · **write** | [v0.1](service-des-domaines/v0.1.md) |

## How these documents work

One PRD per body, versioned on its own line starting at v0.1. Superseded
versions stay on disk and say so at the top, because a version records what I
believed at the time and editing it destroys the only thing it is for. Bumps are
per body — the commune can reach v0.3 while the Domaines is still at v0.1, which
is what working one domain at a time actually produces.

Requirements carry identifiers — `ETC-001`, `DOM-001` — that are never reused
and never renumbered. That is what lets a later claim that something is missing
be checked against something citable, rather than remaining an opinion.

Open questions get no identifier. A named problem is not a settled requirement,
and giving it a number would dress it as one.

Reasoning behind decisions is in [decisions.md](../decisions.md); how records
are kept honest is in [architecture/](../architecture/).

## Two kinds of operation, not three procedures

An acte de naissance and a CSJ are **reads**. The record already exists — it is
simply trapped in one building, and the wait is the cost of reaching it. Making
them instant changes nobody's legal powers; it only removes distance.

A mutation is a **write**. It changes who owns something. Every hard problem
lives here: who may start one, how consent is proven, what happens when two
mutations touch the same parcel, and what happens when one turns out to have
been fraudulent. It is not a third document. It is a different category, and it
carries the risk that the other two do not.

## Why the commune comes first

The civil register is the database of Malagasy people. Every other
administration ultimately names a person, and a name only identifies someone if
there is exactly one authoritative record behind it.

A parcel belonging to "RAKOTO Jean" is unambiguous only if that person can be
resolved to a single record. So the état civil is not first merely because it is
the most painful — the land work is built on top of it. If person identity is
right, land becomes tractable. If it is wrong, every land record inherits the
ambiguity, and a register that confidently records the wrong owner is worse than
no register at all.

## Deliberately out of scope for now

- Every other paper the communes and the Domaines issue.
- Every other administrative body.
- Registering a new birth — see the scope section in
  [commune/v0.1.md](commune/v0.1.md).
- Succession, and certificats fonciers issued by the guichets fonciers
  communaux — see the scope section in
  [service-des-domaines/v0.1.md](service-des-domaines/v0.1.md).
