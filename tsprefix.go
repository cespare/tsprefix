package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	layout := flag.String("format", time.Stamp, "Time format string (see http://golang.org/pkg/time/)")
	utc := flag.Bool("utc", false, "Use UTC instead of local time")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		t := time.Now()
		if *utc {
			t = t.UTC()
		}
		fmt.Printf("%s %s\n", t.Format(*layout), scanner.Bytes())
	}
	// Don't bother checking scanner.Err().
}
