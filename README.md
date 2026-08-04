# Madagascar

Software for the administrative work Malagasy people dread.

Getting a copie d'acte de naissance takes days, and it has to come from the
commune you were born in — so if you were born in Toliara and live in Diego, a
sheet of paper travels 2,000 km to reach you. A CSJ tells you who owns a plot of
land; it should be a lookup, and instead it takes months.

These documents are lookups, not productions. I'm building the system that
treats them that way: one registry every commune can read and write, answering
instantly, with the data provably untampered.

## The podcast

I build this on camera, screen shared, in Malagasy. Two reasons beyond the
software itself:

- **Teaching** — Malagasy developers can watch the whole thing get built,
  decisions and mistakes included.
- **Advocacy** — I would like to see this adopted officially, and the most
  honest case I can make for it is a working system, out in the open, that
  people already find useful. Building it in public means anyone can judge it
  on its merits.

## Open source

All of it, under [AGPL-3.0](LICENSE). For a system that records who owns a plot
of land or who was born where, being able to read every line is part of why
anyone should accept it. Anyone can run it themselves; anyone who runs a
modified version as a service has to publish their changes.

## Layout

| | |
|---|---|
| [`backend/`](backend/) | The register itself — Go, PostgreSQL, one node per commune or service |
| [`frontend/`](frontend/) | The desk application for registrars — React, TypeScript |
| [`mobile_consumer/`](mobile_consumer/) | For Malagasy people: their own records, and authorising anything done in their name |
| [`mobile_admin/`](mobile_admin/) | For registrars, for acts that happen in a room rather than at a desk |
| [`docs/`](docs/) | Design, one administrative domain at a time |

## Language

Everything a user reads is written in **Malagasy first**, then translated into
French and English. Malagasy is the source, not a localisation of something
written in another language. Code, comments, commit messages and these documents
are in English.

## Status

Early. The structure is in place and nothing real is built yet.

I work through the administrative domains one at a time, and each gets designed
before anything is written. The first two bodies and three procedures are
specified in [`docs/prd/`](docs/prd/), one versioned PRD per administrative
body; how records are kept honest is in
[`docs/architecture/`](docs/architecture/); and every decision taken so far,
with its reasoning, is in [`docs/decisions.md`](docs/decisions.md).
