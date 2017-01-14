package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	//Format is input time format
	Format = "2006-01-02 15:04 MST"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "%s <timestring>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	src := strings.Join(os.Args[1:], " ")
	t, err := time.Parse(Format, src)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse error %v", src)
		os.Exit(1)
	}
	utc, _ := time.LoadLocation("UTC")
	pst, _ := time.LoadLocation("America/Los_Angeles")
	utcT := t.In(utc)
	pstT := t.In(pst)
	fmt.Printf("%s\n", t.Format(Format))
	fmt.Printf("%s\n", utcT.Format(Format))
	fmt.Printf("%s\n", pstT.Format(Format))
}
