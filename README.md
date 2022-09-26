# goi2pbrowser
--
    import "github.com/eyedeekay/go-i2pbrowser"


## Usage

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
