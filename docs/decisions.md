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

---

## Open questions

- **Who operates the deployed system.** I would be paid to run and maintain it at
  ordinary hosting cost, but that is not a requirement and the government can
  operate it themselves. Open source means I cannot control who runs it — what
  keeps me the natural operator is authorship, keys the witness network already
  trusts, and competence, not the license.
