// Package store persists a node's data in PostgreSQL.
//
// Two different things live here and must not be confused: the append-only log,
// which is written once and never modified, and the projections derived from it,
// which are caches and may be dropped and rebuilt at any time.
//
// If a projection ever disagrees with the log, the log is right.
package store
