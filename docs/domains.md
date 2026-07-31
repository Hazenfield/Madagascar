# Administrative domains

The administrative bodies this project covers, and the papers handled for each.
I take them one at a time, and each gets designed before anything is built.

Deliberately small to begin with: two bodies, three procedures — the three that
cost me the most time and patience.

| # | Body | Procedure | Kind | Design |
|---|------|-----------|------|--------|
| 1 | Commune — état civil | Acte de naissance | read | [commune.md](domains/commune.md) |
| 2 | Service des Domaines | Certificat de Situation Juridique (CSJ) | read | [service-des-domaines.md](domains/service-des-domaines.md) |
| 3 | Service des Domaines | Mutation d'une parcelle | **write** | [service-des-domaines.md](domains/service-des-domaines.md) |

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
- Certificats fonciers issued by the guichets fonciers communaux — see the
  scope note in [service-des-domaines.md](domains/service-des-domaines.md).
