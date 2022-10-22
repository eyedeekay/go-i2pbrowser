# goi2pbrowser

Package goi2pbrowser is a package which can be used to manage an I2P browsing
profile using a pre-configured, common profile which is used by the I2P Easy-Install
bundle and the i2p.plugins.firefox profile manager. It is a Go clone of
i2p.plugins.firefox for use in native applications.

## Variables

BaseProfile is a zip of a Firefox profile with NoScript, HTTPS Everywhere, and I2PIPB

```golang
var BaseProfile []byte
```

UsabilityProfile is a zip of a Firefox profile with Jshelter, HTTPS Everywhere, uBlock Origin, LocalCDN, OICT, and I2PIPB

```golang
var UsabilityProfile []byte
```

## Functions

### func [BrowseApp](/run.go#L53)

`func BrowseApp(profileDir string, url ...string)`

BrowseApp launches a Firefox browser configured to use I2P and waits for it to exit.
The profile is in "Usability" mode

### func [BrowseStrict](/run.go#L15)

`func BrowseStrict(profileDir string, url ...string)`

BrowseStrict launches a Firefox browser configured to use I2P and waits for it to exit.
The profile is in "Strict" mode

### func [BrowseUsability](/run.go#L34)

`func BrowseUsability(profileDir string, url ...string)`

BrowseUsability launches a Firefox browser configured to use I2P and waits for it to exit.
The profile is in "Usability" mode

### func [UnpackBase](/unembed.go#L22)

`func UnpackBase(profileDir string) (string, error)`

UnpackBase unpacks a "Strict" mode profile into the "profileDir" and returns the
path to the profile and possibly, an error if something goes wrong. If everything
works, the error will be nil

### func [UnpackUsability](/unembed.go#L41)

`func UnpackUsability(profileDir string) (string, error)`

UnpackUsability unpacks a "Usability" mode profile into the "profileDir" and returns the
path to the profile and possibly, an error if something goes wrong. If everything
works, the error will be nil

## Sub Packages

* [i2pbrowser](./i2pbrowser)

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
