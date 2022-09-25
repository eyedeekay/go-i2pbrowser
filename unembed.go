//go:build !generate
// +build !generate

package goi2pbrowser

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/artdarek/go-unzip/pkg/unzip"
)

func existsAlready(profileDir string) bool {
	if _, err := os.Stat(filepath.Join(profileDir, "user.js")); err == nil {
		return true
	}
	return false
}

func UnpackBase(profileDir string) error {
	if existsAlready(profileDir) {
		log.Println(profileDir, "exists already")
		return nil
	}
	os.MkdirAll(filepath.Dir(profileDir), 0755)
	zipFile := filepath.Join(filepath.Dir(profileDir), "i2p.firefox.base.profile.zip")
	err := ioutil.WriteFile(zipFile, baseProfile, 0644)
	if err != nil {
		return err
	}
	uz := unzip.New()
	_, err = uz.Extract(zipFile, profileDir)
	if err != nil {
		return err
	}
	return nil
}

func UnpackUsability(profileDir string) error {
	if existsAlready(profileDir) {
		log.Println(profileDir, "exists already")
		return nil
	}
	os.MkdirAll(filepath.Dir(profileDir), 0755)
	zipFile := filepath.Join(filepath.Dir(profileDir), "i2p.firefox.usability.profile.zip")
	err := ioutil.WriteFile(zipFile, usabilityProfile, 0644)
	if err != nil {
		return err
	}
	uz := unzip.New()
	_, err = uz.Extract(zipFile, profileDir)
	if err != nil {
		return err
	}
	return nil
}
