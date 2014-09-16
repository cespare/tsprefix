# tsprefix

A tiny command-line tool for prepending timestamps to each line of output. It operates on stdin and writes to
stdout.

```
$ (echo hello; sleep 2; echo world) | tsprefix
Sep 15 21:55:19 hello
Sep 15 21:55:21 world
```

The timestamp format may be changed with the `-format` flag. See the [Go time package
documentation](http://golang.org/pkg/time/) for details on the format.

## Installation

You need to have Go installed.

    $ go get github.com/cespare/tsprefix
