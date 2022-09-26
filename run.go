// Package goi2pbrowser is a package which can be used to manage an I2P browsing
// profile using a pre-configured, common profile which is used by the I2P Easy-Install
// bundle and the i2p.plugins.firefox profile manager. It is a Go clone of
// i2p.plugins.firefox for use in native applications.
package goi2pbrowser

import (
	"log"

	fcw "github.com/eyedeekay/go-fpw"
)

// BrowseStrict launches a Firefox browser configured to use I2P and waits for it to exit.
// The profile is in "Strict" mode
func BrowseStrict(profileDir, url string) {
	var profilePath string
	var err error
	profilePath, err = UnpackBase(profileDir)
	if err != nil {
		log.Println(err)
		return
	}
	FIREFOX, ERROR := fcw.BasicFirefox(profilePath, false, url)
	if ERROR != nil {
		log.Println(ERROR)
		return
	}
	defer FIREFOX.Close()
	<-FIREFOX.Done()
}

// BrowseUsability launches a Firefox browser configured to use I2P and waits for it to exit.
// The profile is in "Usability" mode
func BrowseUsability(profileDir, url string) {
	var profilePath string
	var err error
	profilePath, err = UnpackUsability(profileDir)
	if err != nil {
		log.Println(err)
		return
	}
	FIREFOX, ERROR := fcw.BasicFirefox(profilePath, false, url)
	if ERROR != nil {
		log.Println(ERROR)
		return
	}
	defer FIREFOX.Close()
	<-FIREFOX.Done()
}
