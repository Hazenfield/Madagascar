# Commune — état civil

**Procedure covered:** acte de naissance (obtaining a copy).

Main lines only. Each section here gets its own pass later.

## Why this one first

The civil register is the database of Malagasy people. School enrolment, exams,
the CIN, a passport, marriage, inheritance, a land transfer, the electoral roll —
each of them starts by asserting that a particular person exists. That assertion
is only worth anything if it resolves to exactly one authoritative record.

So this is not one document among many. It is the layer everything else stands
on, including both foncier procedures in this project.

## What happens today

A birth is registered at the commune where it happened, in that commune's
register. Copies are issued by that same commune, and only by it.

Consequences:

- If you were born in Toliara and live in Diego, the paper has to make the
  journey. Roughly 2,000 km for a document that is a lookup.
- A copy takes days even when you are standing in the right building.
- Administrations routinely demand a *recent* copy, so the same person requests
  the same unchanged fact over and over.
- Someone whose birth was never registered has no route through this at all,
  only a court procedure.

None of that is caused by the record being hard to find. It is caused by the
register being paper, local, and singular.

## What it should be

Give a name and a date of birth at any commune, or from a phone, and get the
record. Immediately. Wherever you are.

The commune of birth remains the authority that registered the event — that does
not change and should not. What changes is that its register stops being the only
place the fact can be read.

## What the system has to hold

At the level of main lines:

- **The person** — an identity that survives spelling variants and can be pointed
  at by other registers.
- **The birth event** — when, where, declared by whom, recorded by which officer,
  in which register at which page.
- **The parents**, as recorded.
- **Mentions marginales** — marriage, divorce, death, recognition, adoption,
  court judgments. These are annotations added to a birth record years after the
  fact, and they are why a copy is not a static photocopy: an acte de naissance
  is a record plus everything later attached to it. Any design that treats it as
  a fixed document is wrong.
- **Provenance** — which paper register each entry came from, transcribed by
  whom, and when. A digitized record whose origin cannot be traced is not
  evidence.

Corrections are new entries that supersede earlier ones. Nothing is overwritten.

## The hard parts

1. **Identity resolution.** Matching a person across communes with spelling
   variants, transcription errors, approximate or missing dates of birth, and
   homonyms. Malagasy naming does not give the clean surname/given-name split
   most matching algorithms assume. This is the hardest problem in the whole
   project and everything downstream rests on it.
2. **Retro-digitization.** Decades of handwritten registers, in varying
   condition, held in ~1,500 communes. This is a logistics and transcription
   problem far more than a software one.
3. **Unregistered births.** People with no entry at all. The system must
   represent their absence honestly rather than quietly inventing them.
4. **Who may read whose record**, and what is shown to them. An open register of
   every citizen's parentage is not an acceptable outcome.
5. **Working offline.** Communes with intermittent power and no reliable link.
   Anything that assumes a live connection to a central server is unusable in
   most of the country.

## To confirm

Points where I need to check the texts and current practice rather than assume:

- The exact legal status and required form of a copie intégrale versus an
  extrait, and which administrations accept which.
- Whether anything obliges the issuing commune to be the commune of birth, or
  whether that is only how the paper register makes it work.
- The current state of any existing état civil digitization programme, so this
  builds on what exists instead of ignoring it.
- The legal fee, and what a copy is actually supposed to cost.
