package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	layout := flag.String("format", time.Stamp, "Time format string (see http://golang.org/pkg/time/), or else ns/us/ms/s with -elapsed/-delta")
	utc := flag.Bool("utc", false, "Use UTC instead of local time")
	elapsed := flag.Bool("elapsed", false, "Instead of timestamp, show elapsed time since beginning")
	delta := flag.Bool("delta", false, "Instead of timestamp, show delta time since previous line")
	flag.Parse()

	if *elapsed && *delta {
		log.Fatal("-elapsed and -delta are mutually exclusive")
	}
	if *elapsed || *delta {
		if *utc {
			log.Fatal("-utc doesn't make sense in -elapsed/-delta mode")
		}
		if *layout == time.Stamp {
			*layout = "ms"
		}
		if _, err := formatDuration(0, *layout); err != nil {
			log.Fatalf("Bad duration unit %q", *layout)
		}
	}

	var a annotator
	switch {
	case *elapsed:
		a = byElapsed{time.Now(), *layout}
	case *delta:
		a = &byDelta{time.Now(), *layout}
	default:
		a = byTime{*layout, *utc}
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Printf("%s %s\n", a.annotate(), scanner.Bytes())
	}
	// Don't bother checking scanner.Err().
}

type annotator interface {
	annotate() string
}

type byTime struct {
	layout string
	utc    bool
}

func (a byTime) annotate() string {
	t := time.Now()
	if a.utc {
		t = t.UTC()
	}
	return t.Format(a.layout)
}

type byElapsed struct {
	start time.Time
	unit  string
}

func (a byElapsed) annotate() string {
	s, err := formatDuration(time.Since(a.start), a.unit)
	if err != nil {
		panic(err)
	}
	return s
}

type byDelta struct {
	last time.Time
	unit string
}

func (a *byDelta) annotate() string {
	t := time.Now()
	s, err := formatDuration(t.Sub(a.last), a.unit)
	if err != nil {
		panic(err)
	}
	a.last = t
	return s
}

func formatDuration(d time.Duration, unit string) (string, error) {
	switch unit {
	case "ns":
		return fmt.Sprintf("+%12dns", d.Nanoseconds()), nil
	case "us", "μs":
		return fmt.Sprintf("+%10dμs", d.Microseconds()), nil
	case "ms":
		return fmt.Sprintf("+%8dms", d.Milliseconds()), nil
	case "s":
		return fmt.Sprintf("+%10.3fs", d.Seconds()), nil
	default:
		return "", fmt.Errorf("unrecognized unit %q", unit)
	}
}
