# goi2pbrowser
--
    import "github.com/eyedeekay/go-i2pbrowser"

Package goi2pbrowser is a package which can be used to manage an I2P browsing
profile using a pre-configured, common profile which is used by the I2P
Easy-Install bundle and the i2p.plugins.firefox profile manager. It is a Go
clone of i2p.plugins.firefox for use in native applications.

## Usage

```go
var BaseProfile []byte
```
BaseProfile is a zip of a Firefox profile with NoScript, HTTPS Everywhere, and
I2PIPB

```go
var UsabilityProfile []byte
```
UsabilityProfile is a zip of a Firefox profile with Jshelter, HTTPS Everywhere,
uBlock Origin, LocalCDN, OICT, and I2PIPB

#### func  BrowseStrict

```go
func BrowseStrict(profileDir, url string)
```
BrowseStrict launches a Firefox browser configured to use I2P and waits for it
to exit. The profile is in "Strict" mode

#### func  BrowseUsability

```go
func BrowseUsability(profileDir, url string)
```
BrowseUsability launches a Firefox browser configured to use I2P and waits for
it to exit. The profile is in "Usability" mode

#### func  UnpackBase

```go
func UnpackBase(profileDir string) (string, error)
```
UnpackBase unpacks a "Strict" mode profile into the "profileDir" and returns the
path to the profile and possibly, an error if something goes wrong. If
everything works, the error will be nil

#### func  UnpackUsability

```go
func UnpackUsability(profileDir string) (string, error)
```
UnpackUsability unpacks a "Usability" mode profile into the "profileDir" and
returns the path to the profile and possibly, an error if something goes wrong.
If everything works, the error will be nil
