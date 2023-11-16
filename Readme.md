Json vs Proto marshaling/unmarshaling benchmark
----

```
$ > go test -bench=. -benchtime=10s -benchmem
goos: darwin
goarch: arm64
pkg: json_vs_proto
BenchmarkMarshalJson-12          1902319              6075 ns/op            2307 B/op         27 allocs/op
BenchmarkMarshalProto-12         8690508              1220 ns/op            1736 B/op         22 allocs/op
```