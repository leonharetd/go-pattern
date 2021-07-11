package options

import (
	"testing"
)

func TestOptions(t *testing.T) {
	NewDefaultOptions(WithUID("80"), WithGID("ddd"), WithGender(false))
	NewDefaultOptions(WithGID("ccc"))
	NewDefaultOptions(WithGID("ddd"), WithGender(false))
	NewDefaultOptions(WithUID("801"))
	NewDefaultOptions(WithCompany("company"))
}
