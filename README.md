# BITONIC: A bitonic sorter implemented in Go.

[![GoDoc](https://godoc.org/github.com/ebiiim/bitonic?status.svg)](https://godoc.org/github.com/ebiiim/bitonic)
[![Build Status](https://travis-ci.org/ebiiim/bitonic.svg?branch=master)](https://travis-ci.org/ebiiim/bitonic)
[![Go Report Card](https://goreportcard.com/badge/github.com/ebiiim/bitonic)](https://goreportcard.com/report/github.com/ebiiim/bitonic)

```go
x := []int{8, 7, 6, 5, 4, 3, 2, 1}
bitonic.SortInts(x, bitonic.Ascending)
fmt.Print(x) // [1 2 3 4 5 6 7 8]
```
