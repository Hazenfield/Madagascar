// Package translog is the append-only log that holds the record.
//
// Nothing here updates anything. Entries are appended, each committing to
// everything appended before it, and a correction is a new entry superseding an
// older one rather than an edit of it. Current state - who holds a parcel, what
// a birth record says today - is a projection rebuilt from this log, never a
// source of truth in its own right.
//
// This package also owns checkpoints: the signed head of the log, published
// where it cannot be retracted, and co-signed by witnesses.
//
// See docs/architecture/integrity.md.
package translog
