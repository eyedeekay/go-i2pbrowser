package goi2pbrowser

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

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
	log.Println(profileDir, "exists already")
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
	log.Println(profileDir, "exists already")
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

// UnpackApp unpacks a "App" mode profile into the "profileDir" and returns the
// path to the profile and possibly, an error if something goes wrong. If everything
// works, the error will be nil
func UnpackApp(profileDir string) (string, error) {
	log.Println(profileDir, "exists already")
	os.MkdirAll(filepath.Dir(profileDir), 0755)
	zipFile := filepath.Join(filepath.Dir(profileDir), "i2p.firefox.usability.profile.zip")
	err := ioutil.WriteFile(zipFile, UsabilityProfile, 0644)
	if err != nil {
		return filepath.Join(profileDir, "i2p.firefox.app.profile"), err
	}
	uz := unzip.New()
	_, err = uz.Extract(zipFile, profileDir)
	unzippedFiles := strings.Replace(filepath.Join(profileDir, "i2p.firefox.usability.profile.zip"), ".zip", "", 1)
	if !existsAlready(filepath.Join(profileDir, "i2p.firefox.app.profile")) {
		if err = os.Rename(unzippedFiles, filepath.Join(profileDir, "i2p.firefox.app.profile")); err != nil {
			return filepath.Join(profileDir, "i2p.firefox.app.profile"), err
		}
	}
	if err := os.MkdirAll(filepath.Join(profileDir, "i2p.firefox.app.profile", "chrome"), 0755); err != nil {
		return filepath.Join(profileDir, "i2p.firefox.app.profile"), err
	}
	if err := ForceUserChromeCSS(filepath.Join(profileDir, "i2p.firefox.app.profile", "chrome", "userChrome.css")); err != nil {
		return filepath.Join(profileDir, "i2p.firefox.app.profile"), err
	}
	if err := AppifyUserJS(filepath.Join(profileDir, "i2p.firefox.app.profile", "user-overrides.js")); err != nil {
		return filepath.Join(profileDir, "i2p.firefox.app.profile"), err
	}
	if err := AppifyUserJS(filepath.Join(profileDir, "i2p.firefox.app.profile", "user.js")); err != nil {
		return filepath.Join(profileDir, "i2p.firefox.app.profile"), err
	}
	if err := AppifyUserJS(filepath.Join(profileDir, "i2p.firefox.app.profile", "prefs.js")); err != nil {
		return filepath.Join(profileDir, "i2p.firefox.app.profile"), err
	}
	return filepath.Join(profileDir, "i2p.firefox.app.profile"), nil
}

func AppifyUserJS(profile string) error {
	content, err := ioutil.ReadFile(profile)
	if err != nil {
		return err
	}
	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		if strings.Contains(line, "toolkit.legacyUserProfileCustomizations.stylesheets\"") {
			if strings.Contains(line, "true") {
				return nil
			} else {
				line = strings.Replace(line, "false", "true", 1)
			}
		}
		lines[i] = line
	}
	out := strings.Join(lines, "\n")
	if err := ioutil.WriteFile(profile, []byte(out), 0644); err != nil {
		return err
	}
	return nil
}

func ForceUserChromeCSS(profile string) error {
	var userChrome = `@namespace url("http://www.mozilla.org/keymaster/gatekeeper/there.is.only.xul");

/* only needed once */

@namespace html url("http://www.w3.org/1999/xhtml");
#PersonalToolbar,
#PanelUI-Button,
#PanelUI-menu-button,
#star-button,
#forward-button,
#home-button,
#bookmarks-toolbar-button,
#library-button,
#sidebar-button,
#pocket-button,
#fxa-toolbar-menu-button,
#reader-mode-button,
#identity-icon {
    visibility: collapse;
}

#urlbar-background {
    background-color: black !important;
}


/* Remove back button circle */

#back-button:not(:hover),
#back-button:not(:hover)>.toolbarbutton-icon {
    background: transparent !important;
    border: none !important;
    box-shadow: none !important;
}

#back-button:hover,
#back-button:hover>.toolbarbutton-icon {
    border: none !important;
    border-radius: 2px !important;
}

#urlbar-container {
    visibility: collapse !important
}

#TabsToolbar-customization-target {
    min-width: 50vw;
    max-width: 50vw;
    width: 50vw;
}
`
	if err := ioutil.WriteFile(profile, []byte(userChrome), 0644); err != nil {
		return err
	}
	return nil
}
