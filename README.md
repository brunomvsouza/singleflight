# Singleflight with Generics!

[![GoDoc](https://godoc.org/github.com/brunomvsouza/singleflight?status.svg)](https://godoc.org/github.com/brunomvsouza/singleflight)

> Package singleflight provides a duplicate function call suppression mechanism.

This fork adds generics support (`Group[K comparable, V any]` and `Result[V any]`) to the original [x/sync/singleflight](https://golang.org/x/sync/singleflight) package.

### Updates & Versioning

- I will keep this package up-to-date with the original one, at least until `x/sync/singleflight` adds support for generics. **If you notice an update before I do, please open an issue or submit a pull request**.
- Versions will be tagged to align with the same versioning as the `x/sync/singleflight` package.

### Usage

For example usage, see:
- [Group.Do](examples/Do/main.go)
- [Group.DoChan](examples/DoChan/main.go)
- [Group.DoForget](examples/Forget/main.go)
