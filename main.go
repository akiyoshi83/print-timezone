package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	// OutputFormat is format for output time
	OutputFormat = "2006-01-02 15:04 MST -0700"
)

// InputFormats are to parse input time string
var InputFormats = []string{
	"2006-01-02 15:04 MST",
	"2006/01/02 15:04 MST",
	"2006-01-02 15:04:05 MST",
	"2006/01/02 15:04:05 MST",
	"MST 2006-01-02 15:04",
	"MST 2006/01/02 15:04",
	"MST 2006-01-02 15:04:05",
	"MST 2006/01/02 15:04:05",
}

var defaultLocations = []string{
	"UTC",
	"America/Los_Angeles",
	"Europe/London",
	"Asia/Tokyo",
	"Australia/Sydney",
}

var (
	confPath  string
	inputTime string
)

func main() {
	parseArgs()

	t, err := tryParseTime(InputFormats, inputTime)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse error %v\n", inputTime)
		fmt.Fprintf(os.Stderr, "\n---SUPPORTED FORMATS---\n%v\n\n", strings.Join(InputFormats, "\n"))
		os.Exit(1)
	}

	printWithTimezone(t, defaultLocations)
}

func parseArgs() {
	flag.StringVar(&confPath, "f", "", "configuration file path")
	flag.Parse()
	if flag.NArg() < 2 {
		inputTime = time.Now().Format(InputFormats[0])
	} else {
		inputTime = strings.Join(flag.Args()[:], " ")
	}
}

func tryParseTime(formats []string, s string) (time.Time, error) {
	var parsed time.Time
	var err error
	for _, format := range formats {
		var t time.Time
		t, err = time.Parse(format, s)
		if err == nil {
			parsed = t
			break
		}
	}
	if err != nil {
		return parsed, err
	}
	return parsed, nil
}

func formatWithTimezone(t time.Time, loc string) (string, error) {
	l, err := time.LoadLocation(loc)
	if err != nil {
		return "", fmt.Errorf("formatWithTimezone %v", err)
	}
	lt := t.In(l)
	return lt.Format(OutputFormat), nil
}

func printWithTimezone(t time.Time, locations []string) {
	for _, loc := range locations {
		str, err := formatWithTimezone(t, loc)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
		}
		fmt.Printf("%s\n", str)
	}
}
