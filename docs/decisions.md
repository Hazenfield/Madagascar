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

What this does not solve is stated in the design and should be repeated whenever
the system is described — it guarantees the record, not the truth.

---

## Open questions

- **Who operates the deployed system.** I would be paid to run and maintain it at
  ordinary hosting cost, but that is not a requirement and the government can
  operate it themselves. Open source means I cannot control who runs it — what
  keeps me the natural operator is authorship, keys the witness network already
  trusts, and competence, not the license.
