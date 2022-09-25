package goi2pbrowser

import (
	"testing"
)

func TestUnEmbded(t *testing.T) {
	t.Log("testing base")
	UnpackBase("testing/Base")
	t.Log("testing usability")
	UnpackUsability("testing/Usability")

}
