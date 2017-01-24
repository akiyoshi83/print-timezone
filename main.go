package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/akiyoshi83/print-timezone/ptz"
)

const (
	confFileName = ".print-timezone.yml"
)

var (
	defaultConfPath string
	confPath        string
	inputTime       string
	pptz            *ptz.Ptz
)

func init() {
	defaultConfPath = filepath.Join(homeDir(), confFileName)
	pptz = ptz.NewPtz()
}

func main() {
	parseArgs()
	loadConfig(confPath)

	t, err := pptz.TryParseTime(inputTime)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse error %v\n", inputTime)
		fmt.Fprintf(os.Stderr, "\n---SUPPORTED FORMATS---\n%v\n\n", strings.Join(pptz.InputFormats(), "\n"))
		os.Exit(1)
	}

	pptz.PrintWithTimezone(t)
}

func parseArgs() {
	flag.StringVar(&confPath, "f", "", "configuration file path")
	flag.Parse()
	if flag.NArg() < 1 {
		inputTime = time.Now().Format(pptz.InputFormats()[0])
	} else {
		inputTime = strings.Join(flag.Args()[:], " ")
	}
}

func loadConfig(confPath string) error {
	var fpath string
	var err error

	if confPath != "" {
		if !exists(confPath) {
			return fmt.Errorf("%s does not exist", confPath)
		}
		fpath = confPath
	} else {
		if !exists(defaultConfPath) {
			return nil
		}
		fpath = defaultConfPath
	}

	data, err := ioutil.ReadFile(fpath)
	if err != nil {
		return err
	}

	pptz.LoadFromYaml(data)
	return nil
}

func homeDir() string {
	home := os.Getenv("HOME")
	if home == "" && runtime.GOOS == "windows" {
		home = os.Getenv("USERPROFILE")
	}
	return home
}

func exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
