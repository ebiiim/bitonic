bits = 26

bench:
	BITS=${bits} go test -bench '^(BenchmarkSortInt)$$' -benchmem

test:
	go test -race -cover
