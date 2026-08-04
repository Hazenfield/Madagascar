// Package registry holds what the records mean.
//
// Persons, birth events and the annotations attached to them; parcels, the
// holders of a parcel over time, and what encumbers it. This package answers
// questions by reading projections of the log - it never writes state directly.
//
// It knows nothing about HTTP, and nothing about who is allowed to ask.
//
// See docs/prd/.
package registry
