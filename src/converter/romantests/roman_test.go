package romantests

import (
	"converter/roman"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test naming convention : Test{Func name}_{Given}_{Returns}(t *testing.T)

// TestConvert_1_I Given a parameter of 1 roman.Convert() returns "I"
func TestConvert_1_I(t *testing.T) {

	var convertedNumber1 = roman.Convert(1)
	assert.Equal(t, "I", convertedNumber1, "1 = I")
}

// TestConvert_2_II Given a parameter of 2, roman.Convert() returns "II"
func TestConvert_2_II(t *testing.T) {

	var convertedNumber2 = roman.Convert(2)
	assert.Equal(t, "II", convertedNumber2, "2 = II")
}
