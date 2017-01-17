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
	conf            *ptz.Conf
)

func init() {
	defaultConfPath = filepath.Join(homeDir(), confFileName)
	conf = ptz.NewConf()
}

func main() {
	parseArgs()
	loadConfig(confPath)

	t, err := ptz.TryParseTime(ptz.InputFormats, inputTime)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse error %v\n", inputTime)
		fmt.Fprintf(os.Stderr, "\n---SUPPORTED FORMATS---\n%v\n\n", strings.Join(ptz.InputFormats, "\n"))
		os.Exit(1)
	}

	ptz.PrintWithTimezone(t, conf.Locations)
}

func parseArgs() {
	flag.StringVar(&confPath, "f", "", "configuration file path")
	flag.Parse()
	if flag.NArg() < 2 {
		inputTime = time.Now().Format(ptz.InputFormats[0])
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

	conf.LoadFromYaml(data)
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
