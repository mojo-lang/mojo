package fake

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPhoneNumberFormat(t *testing.T) {
	for i := 0; i < 100; i++ {
		assert.Regexp(t, `^1[0-9]{10}$`, PhoneNumber())
	}
}
