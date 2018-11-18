[![GoDoc](https://godoc.org/github.com/fishy/wrapreader?status.svg)](https://godoc.org/github.com/fishy/wrapreader)
[![Go Report Card](https://goreportcard.com/badge/github.com/fishy/wrapreader)](https://goreportcard.com/report/github.com/fishy/wrapreader)

# WrapReader

WrapReader is a [Go](https://golang.org)
[library](https://godoc.org/github.com/fishy/wrapreader).
It provides an [`io.ReadCloser`](https://godoc.org/io#ReadCloser)
that wraps an [`io.Reader`](https://godoc.org/io#Reader) and
an [`io.Closer`](https://godoc.org/io#Closer) together.

It's useful when dealing with `io.Reader` implementations that wraps another
`io.ReadCloser`, but will not close the underlying reader, such as
[`gzip.Reader`](https://godoc.org/compress/gzip#Reader).

([Example code on godoc](https://godoc.org/github.com/fishy/wrapreader#example-Wrap))
