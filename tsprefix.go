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
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Printf("%s %s\n", time.Now().Format(*layout), scanner.Bytes())
	}
	// Don't bother checking scanner.Err().
}
