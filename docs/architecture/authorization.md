# Authorization

Who may do what. This sits above the [integrity layer](integrity.md) and varies
by procedure, where integrity does not.

Main lines only.

## What this replaces

Today a registrar can enter a person who does not exist, and a paper
register makes that invisible. There is no attribution, no second pair of eyes,
and no way for anyone outside the office to look.

So roles and a validation step are not a modest improvement here. They are the
first control of any kind.

## Shape

Roles and access levels, in the way an ERP does it — familiar, unglamorous, and
sufficient for the abuses that actually happen:

- **Rights attach to roles, not to people.** Who holds a role is itself a
  recorded, privileged act.
- **Entry and validation are separate.** Sensitive acts are drafted at one level
  and only become final when validated at a higher one. The draft is already in
  the log; validation is a second entry, not an edit of the first.
- **Scope by territory.** An agent of one commune works on that commune's
  records. Reach beyond it is an explicit grant, and it is visible.
- **Technical administration is not a business authority.** Whoever operates the
  servers must hold no role that can validate a civil-status entry or a transfer.
  Separating those two costs nothing and removes the most obvious single point of
  abuse.
- **Reads are logged too.** Who consulted which record, and when. In an
  environment where information itself is sold, a consultation trail is a control
  in its own right, not merely an audit nicety.

## First sketch of the roles

To be refined per domain, not settled here.

| Role | Can do |
|---|---|
| Agent | Enter and draft. Cannot make anything final. |
| Officer in charge | Validate what agents draft, within their scope. |
| Auditor | Read everything, write nothing. |
| Technical administrator | Operate the system. No business authority at all. |

## Deferred

- Per-procedure rules — in particular what a mutation requires, which is a
  different question from what a birth registration requires.
- Personal cryptographic keys for high-value acts.

## Open

- Which acts specifically need validation at a higher level, and which do not.
  Requiring it everywhere would reproduce the delays this project exists to
  remove.
- How role assignment is itself controlled — granting someone the right to
  validate transfers is a more dangerous act than most transfers.
- What an officer sees when they are about to do something unusual. A warning at
  the moment of the act is worth more than a report nobody reads.
