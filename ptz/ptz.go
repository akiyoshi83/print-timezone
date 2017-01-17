package ptz

import (
	"fmt"
	"os"
	"time"
)

// Conf is application configuration
type Conf struct {
	Locations []string
}

// NewConf return new ptz.Conf pointer
func NewConf() *Conf {
	c := new(Conf)
	c.SetDefaultLocation()
	return c
}

// SetDefaultLocation set default timezone locations
func (c *Conf) SetDefaultLocation() {
	c.Locations = append(c.Locations, DefaultLocations...)
}

const (
	// ConfFileName is application configuration file name
	ConfFileName = ".print-timezone.yml"
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

// DefaultLocations is default supported timezone locations
var DefaultLocations = []string{
	"UTC",
	"America/Los_Angeles",
	"Europe/London",
	"Asia/Tokyo",
	"Australia/Sydney",
}

// TryParseTime parse `s` by `formats`
func TryParseTime(formats []string, s string) (time.Time, error) {
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

// FormatWithTimezone format `t` in `loc` and return formated string
func FormatWithTimezone(t time.Time, loc string) (string, error) {
	l, err := time.LoadLocation(loc)
	if err != nil {
		return "", fmt.Errorf("FormatWithTimezone %v", err)
	}
	lt := t.In(l)
	return lt.Format(OutputFormat), nil
}

// PrintWithTimezone write formated string to stdout
func PrintWithTimezone(t time.Time, locations []string) {
	for _, loc := range locations {
		str, err := FormatWithTimezone(t, loc)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
		}
		fmt.Printf("%s\n", str)
	}
}
