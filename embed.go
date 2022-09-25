//go:build !generate
// +build !generate

package goi2pbrowser

import _ "embed"

//go:embed i2p.firefox.base.profile.zip
var baseProfile []byte

//go:embed i2p.firefox.usability.profile.zip
var usabilityProfile []byte

//go:generate go run --tags=generate generate.go
