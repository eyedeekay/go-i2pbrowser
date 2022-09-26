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

// UnpackBase unpacks a "Strict" mode profile into the "profileDir" and returns the
// path to the profile and possibly, an error if something goes wrong. If everything
// works, the error will be nil
func UnpackBase(profileDir string) (string, error) {
	if existsAlready(profileDir) {
		log.Println(profileDir, "exists already")
		return filepath.Join(profileDir, "i2p.firefox.base.profile"), nil
	}
	os.MkdirAll(filepath.Dir(profileDir), 0755)
	zipFile := filepath.Join(filepath.Dir(profileDir), "i2p.firefox.base.profile.zip")
	err := ioutil.WriteFile(zipFile, BaseProfile, 0644)
	if err != nil {
		return filepath.Join(profileDir, "i2p.firefox.base.profile"), err
	}
	uz := unzip.New()
	_, err = uz.Extract(zipFile, profileDir)
	if err != nil {
		return filepath.Join(profileDir, "i2p.firefox.base.profile"), err
	}
	return filepath.Join(profileDir, "i2p.firefox.base.profile"), nil
}

// UnpackUsability unpacks a "Usability" mode profile into the "profileDir" and returns the
// path to the profile and possibly, an error if something goes wrong. If everything
// works, the error will be nil
func UnpackUsability(profileDir string) (string, error) {
	if existsAlready(profileDir) {
		log.Println(profileDir, "exists already")
		return filepath.Join(profileDir, "i2p.firefox.usability.profile"), nil
	}
	os.MkdirAll(filepath.Dir(profileDir), 0755)
	zipFile := filepath.Join(filepath.Dir(profileDir), "i2p.firefox.usability.profile.zip")
	err := ioutil.WriteFile(zipFile, UsabilityProfile, 0644)
	if err != nil {
		return filepath.Join(profileDir, "i2p.firefox.usability.profile"), err
	}
	uz := unzip.New()
	_, err = uz.Extract(zipFile, profileDir)
	if err != nil {
		return filepath.Join(profileDir, "i2p.firefox.usability.profile"), err
	}
	return filepath.Join(profileDir, "i2p.firefox.usability.profile"), nil
}
