package ptz

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

// Ptz print time in multiple timezone
type Ptz struct {
	locations    []string
	inputFormats []string
	outputFormat string
}

var defaultLocations = []string{
	"UTC",
	"America/Los_Angeles",
	"Europe/London",
	"Asia/Tokyo",
	"Australia/Sydney",
}

var defaultInputFormats = []string{
	"2006-01-02 15:04 MST",
	"2006/01/02 15:04 MST",
	"2006-01-02 15:04:05 MST",
	"2006/01/02 15:04:05 MST",
	"MST 2006-01-02 15:04",
	"MST 2006/01/02 15:04",
	"MST 2006-01-02 15:04:05",
	"MST 2006/01/02 15:04:05",
}

const defaultOutputFormat = "2006-01-02 15:04 MST -0700"

// NewPtz create New Ptz object pointer
func NewPtz() *Ptz {
	p := new(Ptz)
	return p
}

// LoadFromYaml loads ptz.Ptz configuration from yaml
func (p *Ptz) LoadFromYaml(data []byte) {
	yaml.Unmarshal(data, p)
}

// Locations are supported timezone
func (p *Ptz) Locations() []string {
	if len(p.locations) == 0 {
		return defaultLocations
	}
	return p.locations
}

// InputFormats are used for parse input time from string
func (p *Ptz) InputFormats() []string {
	if len(p.inputFormats) == 0 {
		return defaultInputFormats
	}
	return p.inputFormats
}

// OutputFormat is parse output time to string
func (p *Ptz) OutputFormat() string {
	if len(p.outputFormat) == 0 {
		return defaultOutputFormat
	}
	return p.outputFormat
}

// TryParseTime parse `s` by `formats`
func (p *Ptz) TryParseTime(s string) (time.Time, error) {
	var parsed time.Time
	var err error
	for _, format := range p.InputFormats() {
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
func (p *Ptz) FormatWithTimezone(t time.Time, loc string) (string, error) {
	l, err := time.LoadLocation(loc)
	if err != nil {
		return "", fmt.Errorf("FormatWithTimezone %v", err)
	}
	lt := t.In(l)
	return lt.Format(p.OutputFormat()), nil
}

// PrintWithTimezone write formated string to stdout
func (p *Ptz) PrintWithTimezone(t time.Time) {
	for _, loc := range p.Locations() {
		str, err := p.FormatWithTimezone(t, loc)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
		}
		fmt.Printf("%s\n", str)
	}
}
