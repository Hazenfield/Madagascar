// Package authz decides who may do what.
//
// Roles rather than individuals, scoped by territory, with sensitive acts drafted
// at one level and made final at a higher one. Whoever operates a node holds no
// role here: technical administration is deliberately not a business authority.
//
// This sits above the integrity guarantees and varies by procedure, where those
// guarantees do not.
//
// See docs/architecture/authorization.md.
package authz
