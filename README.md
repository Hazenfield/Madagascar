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
- **Pressure** — once the tool exists, works, and people want it, not adopting
  it becomes the thing that needs explaining.

## Open source

All of it, under [AGPL-3.0](LICENSE). For a system that records who owns a plot
of land or who was born where, being able to read every line is part of why
anyone should accept it. Anyone can run it themselves; anyone who runs a
modified version as a service has to publish their changes.

## Status

Early. No code yet. I'm working through the administrative domains one at a
time, and each one gets designed before anything gets built. That work lives in
[`docs/`](docs/) — including a log of the [decisions](docs/decisions.md) taken
so far.
