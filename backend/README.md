# backend

A register node. One runs per commune or service.

Go, PostgreSQL, no framework. It builds to a single static binary, which is the
point: a commune node has to run on whatever hardware is in the room, survive
losing power mid-request, and keep working for hours without a network.

## Running it

```bash
make db-up     # PostgreSQL in Docker
make run       # the server, on :8080
make check     # format, vet, test
make help      # everything else
```

`DATABASE_URL` must be set. The Makefile provides a local default; the server
refuses to start without one rather than guessing, because a node silently
running against the wrong database is worse than a node that does not start.

## Layout

```
cmd/server/        the binary
internal/
  config/          settings, read from the environment
  httpapi/         HTTP handlers - no domain logic lives here
  translog/        the append-only log and its checkpoints
  registry/        what records mean: persons, births, parcels, holders
  authz/           who may do what
  store/           PostgreSQL: the log, and the projections built from it
```

The dependency direction is one way. `httpapi` decodes and delegates; `registry`
and `authz` know nothing about HTTP; `store` knows nothing about either.

## The one rule that shapes everything

Nothing is updated. Entries are appended, each committing to everything before
it, and a correction is a new entry superseding an old one. Anything that looks
like current state — who holds a parcel, what a birth record says today — is a
projection that can be dropped and rebuilt from the log.

If a projection ever disagrees with the log, the log is right.

See [docs/architecture/integrity.md](../docs/architecture/integrity.md).
