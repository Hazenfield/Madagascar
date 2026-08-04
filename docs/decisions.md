# Decisions

A running log of the decisions that shape this project, newest last. Each entry
records what was decided and why, so the reasoning survives the conversation it
came from.

---

## 2026-07-31 — Completely open source

The whole system is open source.

For a registry that decides who owns a plot of land or who was born where, being
able to read every line is not a courtesy — it is part of why anyone should
accept the system at all. The integrity argument only closes if all three pieces
hold together: the code is public, the record is public and append-only, and the
witnesses are independent. Closed source would quietly undercut the other two,
because "trust the operator's binary" is exactly the assumption the whole design
exists to remove.

It also makes the case for adoption a straightforward one. The system is free,
it can be inspected in full, and any institution that would rather not have me
running it can run it themselves. Nobody has to take my word for anything.

## 2026-07-31 — AGPL-3.0

Licensed AGPL-3.0.

Anyone may use, modify and self-host it. Anyone who runs it as a service must
publish their modifications. It is the only widely-used license that binds the
*operator* rather than only the distributor, which makes it the legal echo of
what the transparency log does technically — you cannot run a modified version
of this in secret.

It also closes the realistic bad ending: a firm forks the code, keeps its changes
private, and sells the result as a government contract.

Trade-off accepted: some organisations refuse AGPL software on policy grounds,
so this costs a little adoption. Worth it here.

## 2026-07-31 — Public repository, PR-only branches

`Hazenfield/Madagascar`, public from the first commit.

`main`, `staging` and `dev` accept pull requests only. Nobody pushes to them
directly except repository admins, and every merge needs an approving review
from a code owner, resolved conversations, a linear history and cryptographically
signed commits. Force pushes and branch deletion are blocked outright.

The signing requirement is not ceremony. A project whose central claim is that
its records cannot be altered without detection cannot have a source history that
can be. Anyone should be able to check that the code running a register is the
code that was reviewed.

Admins keep a bypass for now, while I am effectively working alone and the
overhead of a pull request for a one-line documentation fix is not yet worth it.
That bypass should be removed once anyone else is contributing — the rules are
worth more when they bind me too.

One org-level detail worth knowing: the Hazenfield organisation grants members
write access by default, and GitHub does not allow a repository to grant less
than the organisation's baseline. So branch rules, not repository permissions,
are what actually keep the protected branches safe here.

Contributions from forks cannot read secrets, and their workflows require
explicit approval before they run. Secret scanning and push protection are on.

## 2026-07-31 — Integrity as a general layer, not a per-procedure feature

Records are protected by an append-only, hash-chained log whose checkpoints are
co-signed by independent witnesses, with every issued document carrying a proof
that citizens keep. Full design in
[architecture/integrity.md](architecture/integrity.md).

Two things settled here.

**The claim is detection, not prevention.** Anyone with root can change bytes;
pretending otherwise would be dishonest and would collapse the first time someone
competent looked. What the design guarantees is that a change is either rejected
or provable by anyone, without trusting whoever runs the system — me included.
That converts a single corrupt administrator into a conspiracy among parties who
have no reason to cooperate.

**Integrity is universal; authorization is per-procedure.** The integrity layer
sees signed entries and knows nothing about what a birth or a parcel is. Rules
about who may write what, and how many signatures an act needs, sit in a separate
layer above it. Adding a new administrative domain must never mean touching the
security core.

Witnesses are chosen for divergent interests: other communes, banks and notaries
who lose money if the register lies, universities and the press, plus a public
blockchain used only to publish a hash as a timestamp that cannot be retracted.
No token, no smart contract, and nothing about running the system depends on it.
It is one signature in a quorum, and it is the one that survives every domestic
witness being pressured at once.

Not built on a blockchain otherwise: it solves none of the actual failure modes,
and it needs connectivity and running costs that communes with power cuts do not
have.

**Roles before keys.** Writes are attributed to accounts with roles, with
sensitive acts validated at a higher level — see
[architecture/authorization.md](architecture/authorization.md) — rather than
signed by a personal key per officer. Measured against a paper register where a
registrar can enter a person who does not exist and nothing shows it, roles
and a validation step are not a modest control; they are the first one. Personal
keys remain an upgrade the log format already accommodates, worth adding for
high-value acts later. What this concedes is attribution — a technical
administrator could act as another user — and not history, which stays
append-only and witnessed.

**A witness that goes quiet is not the same as a witness that refuses.** Refusal
means it looked at a checkpoint and would not sign it, which is the alarm working
rather than a fault, so the two are recorded differently and a witness can
publish a refusal with its reason. When quorum degrades the system keeps issuing
— withholding someone's birth certificate because a university's server is down
would punish citizens for somebody else's infrastructure — but the degraded state
is public, documents issued during it say so, and an inspection re-verifies the
log against every checkpoint still in circulation before the record is re-anchored
to the witness set. The conclusion is itself an entry, so the record carries its
own audit history.

What this does not solve is stated in the design and should be repeated whenever
the system is described — it guarantees the record, not the truth.

## 2026-08-01 — Go for the register, TypeScript for the clients

`backend/` in Go with PostgreSQL. `frontend/` in React and TypeScript.
`mobile_consumer/` and `mobile_admin/` in React Native.

Go because of where this has to run. A register node belongs in a commune, on
whatever hardware is in the room, surviving power cuts and hours without a
network. Go builds a single static binary with no runtime to install and no
dependency tree to keep alive on a machine nobody will maintain. The reference
implementations of transparency logs are also Go, which matters for the part of
this that must be got right rather than invented.

The cost is real and accepted: Go has the smallest pool of developers in
Madagascar of the options considered, it is slower for me to write than Python,
and it is the least familiar thing to show on camera. Two languages across four
applications is the ceiling I am willing to carry.

React Native for both mobile applications so they share a language with the web
client. Four applications maintained by a very small number of people is the
binding constraint, not raw performance.

## 2026-08-01 — One versioned PRD per administrative body

Each administrative body gets one PRD, in [`prd/`](prd/), versioned on its own
line from v0.1. The design notes that were in `domains/` became the v0.1 of their
body; they are not kept alongside.

The reason is traceability. Once code exists, every claim that something is
missing — a review finding, a bug report, my own mid-task observation — has to
trace back to a requirement, or it is an opinion, and building an opinion either
spends review time on work nobody asked for or silently changes agreed behaviour.
A prose paragraph cannot be traced to. `ETC-002` can. So the delta between the
old design notes and a PRD is exactly this: stable identifiers, explicit scope,
and a statement of what done means.

Identifiers are never reused and never renumbered. A dropped requirement is
marked withdrawn in place with its reason. If `ETC-004` means one thing in v0.1
and another in v0.3, every citation of it rots silently, which is the failure the
identifiers exist to prevent.

Versions are per body, with superseded files kept on disk and marked as
superseded at the top. The commune reaching v0.3 while the Domaines is still at
v0.1 is the honest picture of working one domain at a time; a synchronised
version would either misrepresent that or generate copies of documents that did
not change. A version file is immutable once superseded — it records what I
believed at the time, and editing it destroys the only thing it is for.

Cross-references between PRDs name a specific version rather than "current". A
PRD version is a point-in-time statement, and one that cited commune v0.1 stays
accurate forever — the same reasoning that makes a CSJ answerable as of a date.

Open questions get no identifier. Consent in a mutation, offline reconciliation,
identity resolution across spelling variants: named problems, not settled
requirements. Numbering them would dress an unanswered question as an agreed one.

## 2026-08-01 — Registrar, not fonctionnaire

The English word throughout is **registrar**.

It names the function rather than the employment status, which is what these
documents are about: a person whose job is the register. It is also the standard
English term in both domains covered here — Registrar of Births, Deaths and
Marriages, and Land Registrar. "Civil servant" points at a payroll instead of a
register, and sits badly next to "civil register".

It does not collide with the roles in
[architecture/authorization.md](architecture/authorization.md), which are Agent
and Officer in charge. Registrar is the generic word above those, exactly where
fonctionnaire sat.

## 2026-08-01 — Malagasy is the source language

Everything a user reads is written in Malagasy first and translated into French
and English. Code, comments, commit messages and these documents are in English.

Malagasy is the source, not a target. A string written in English and then
localised carries the shape of English, and the people this is built for are
Malagasy — the same reason the podcast is in Malagasy. Where a translation cannot
carry what the Malagasy says, the Malagasy is right and the translation gets
reworked.

French and English are carried because administrations, courts and any
institution assessing this system work in French, and because an international
reader should be able to audit a system that holds public records. Neither is a
reason to write in them first.

The split at the code boundary is deliberate. Making the interface Malagasy costs
nothing in reviewability; making the *code* Malagasy would exclude every reviewer
who does not speak it, and this project's argument depends on being reviewable by
people who were not in the room. The contributor rules are in
[CONTRIBUTING.md](../CONTRIBUTING.md).

---

## Open questions

- **Who operates the deployed system.** I would be paid to run and maintain it at
  ordinary hosting cost, but that is not a requirement and the government can
  operate it themselves. Open source means I cannot control who runs it — what
  keeps me the natural operator is authorship, keys the witness network already
  trusts, and competence, not the license.
