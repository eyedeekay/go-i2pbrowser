package goi2pbrowser

import _ "embed"

// BaseProfile is a zip of a Firefox profile with NoScript, HTTPS Everywhere, and I2PIPB
//
//go:embed i2p.firefox.base.profile.zip
var BaseProfile []byte

// UsabilityProfile is a zip of a Firefox profile with Jshelter, HTTPS Everywhere, uBlock Origin, LocalCDN, OICT, and I2PIPB
//
//go:embed i2p.firefox.usability.profile.zip
var UsabilityProfile []byte

//go:generate go run --tags=generate generate.go
