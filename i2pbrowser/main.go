package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	goi2pbrowser "github.com/eyedeekay/go-i2pbrowser"
)

var (
	u = flag.Bool("usability", false, "Launch in usability mode")
	a = flag.Bool("application", false, "Launch in app mode")
	d = flag.String("directory", defaultDir(), "Directory to store profiles in")
)

func defaultDir() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	ret := filepath.Join(wd, "i2p-profiles")
	userHome, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	if wd == userHome {
		ret = filepath.Join(userHome, ".i2p/plugins/i2pbrowser")
	}
	return ret
}

func main() {
	flag.Parse()
	url := "http://127.0.0.1:7657"
	if len(flag.Args()) > 0 {
		url = flag.Arg(0)
	}
	application := *a
	if application {
		goi2pbrowser.BrowseApp(*d, url)
		return
	}
	usability := *u
	switch usability {
	case true:
		goi2pbrowser.BrowseUsability(*d, url)
	default:
		goi2pbrowser.BrowseStrict(*d, url)
	}
}
