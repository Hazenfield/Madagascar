# Contributing

Contributions are welcome. This project is meant to be run by whoever ends up
running it, so the code has to be readable and reviewable by people who weren't
in the room when it was written.

## How to propose a change

Nobody pushes to `main`, `staging` or `dev` — those branches only accept pull
requests.

1. Fork the repository (or, if you have write access, push a branch).
2. Make your change on a branch.
3. Open a pull request.
4. A review and a passing check are required before it merges.

## Sign your commits off (DCO)

Every commit must carry a `Signed-off-by:` line. Sign off with:

```bash
git commit -s -m "your message"
```

which appends:

```
Signed-off-by: Your Name <your.email@example.com>
```

By signing off you certify the [Developer Certificate of Origin](DCO) — in
short, that you wrote the change or otherwise have the right to submit it under
this project's license, and that your contribution and sign-off become part of
a permanent public record.

Use your real name. Anonymous or pseudonymous sign-offs don't certify anything.

## Cryptographic signing

Commits on protected branches must also be cryptographically signed, so that
authorship is verifiable rather than merely claimed. SSH signing is the simplest
route:

```bash
git config --global gpg.format ssh
git config --global user.signingkey ~/.ssh/id_ed25519.pub
git config --global commit.gpgsign true
```

Then add that same public key to your GitHub account **a second time**, with key
type **Signing Key** — GitHub tracks authentication keys and signing keys
separately, and a key registered only for authentication will still show your
commits as unverified.

## License

This project is licensed [AGPL-3.0](LICENSE). Contributions are accepted under
the same license. Note that AGPL obliges anyone who runs a modified version as a
network service to publish their modifications — that is deliberate, and it is
part of why this system can be trusted to hold public records.

## What makes a good contribution here

This software is intended to hold civil records. Two things matter more than
they would elsewhere:

- **Nothing is deleted or overwritten.** Corrections are new entries that
  supersede old ones. If a change makes history mutable, it will be rejected
  regardless of how convenient it is.
- **It has to work offline, on bad connections and through power cuts.** A
  feature that assumes a reliable link to a central server is not usable in most
  of the country.
