package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	// InputFormat is format for input time
	InputFormat = "2006-01-02 15:04 MST"
	// OutputFormat is format for output time
	OutputFormat = "2006-01-02 15:04 MST -0700"
)

var defaultLocations = [...]string{
	"UTC",
	"America/Los_Angeles",
	"Europe/London",
	"Asia/Tokyo",
}

func main() {
	var src string
	if len(os.Args) < 2 {
		src = time.Now().Format(InputFormat)
	} else {
		src = strings.Join(os.Args[1:], " ")
	}

	t, err := time.Parse(InputFormat, src)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse error %v", src)
		os.Exit(1)
	}

	locations := defaultLocations

	for _, loc := range locations {
		str, err := formatWithTimezone(t, loc)
		if err != nil {
			fmt.Fprintf(os.Stderr, "convert error %v", err)
		}
		fmt.Printf("%s\n", str)
	}
}

func formatWithTimezone(t time.Time, loc string) (string, error) {
	l, err := time.LoadLocation(loc)
	if err != nil {
		return "", err
	}
	lt := t.In(l)
	return lt.Format(OutputFormat), nil
}
