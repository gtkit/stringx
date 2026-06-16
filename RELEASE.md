# Release Notes

## Next

This release is ready for rollout after the production-safety review focused on performance, memory safety, concurrency safety, and API ergonomics.

### Compatibility

- The public API remains source-compatible.
- `RandId` and `Randn` are kept for compatibility.
- The root package continues to act as the compatibility layer, while feature subpackages remain the recommended import path for new code.
- `.golangci.yml` has been migrated to golangci-lint v2 format. CI must use golangci-lint v2.x.

### Security

- Upgraded `golang.org/x/net` to `v0.55.0` to address `GO-2026-5026` in `golang.org/x/net/idna`.
- Upgraded `golang.org/x/text` to `v0.37.0`.
- `govulncheck ./...` reports no vulnerabilities.
- `gosec` passes with project-intentional exclusions for explicit unsafe helpers, non-cryptographic random helpers, and reviewed integer range conversions.

### Added

- `RandomFromCharset(length int, charset []byte) string` in root package and `randomx`.
- `CloneStringBytes(s string) []byte` for mutable byte copies.
- `CloneBytesString(b []byte) string` for stable string snapshots.
- Regression tests for large integer random ranges, URL port validation, `NUL` rune translation, UTF-8 byte truncation, safe byte/string copies, and literal inflection custom words.

### Changed

- Root random helpers now reuse `randomx` implementation to reduce duplication.
- `NormalizeSpace` now has normalized-string and ASCII fast paths.
- `ValidateHTTPURL` accepts case-insensitive HTTP/HTTPS schemes and rejects invalid ports.
- `inflectx` treats custom uncountable and irregular words as literals instead of regular-expression fragments.

### Integration Checks

- `github.com/gtkit/migrate` passed `go test ./...` with `github.com/gtkit/stringx` replaced by this local checkout.
- `github.com/gtkit/migrate/v2` compiled and ran with the same local replacement, but its own `make` package tests failed because expected generated files differ from its current generator output (`repository.go` expected, `user_i.go`/`user_util.go` generated). This appears unrelated to `stringx`.

### Verification

- `go test ./...`
- `go test -race ./...`
- `go vet ./...`
- `go test ./... -count=5`
- `go test -cover ./...`
- `golangci-lint config verify -c .golangci.yml`
- `golangci-lint run ./...`
- `govulncheck ./...`
- `gosec -exclude=G103,G104,G115,G304,G404 ./...`
- `go test ./... -run '^$' -bench 'Benchmark(ToCamel|ToSnake|NormalizeSpace|BuilderJoin|IsChinaIDCard|IsBankCard)' -benchmem`
