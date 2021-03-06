package romantests

import (
	"converter/roman"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test naming convention : Test{Func name}_{Given}_{Returns}(t *testing.T)

// TestConvert_1_I Given a parameter of 1 roman.Convert() returns "I"
func TestConvert_1_I(t *testing.T) {

	var convertedNumber = roman.Convert(1)
	assert.Equal(t, "I", convertedNumber, "1 = I")
}

// TestConvert_2_II Given a parameter of 2, roman.Convert() returns "II"
func TestConvert_2_II(t *testing.T) {

	var convertedNumber = roman.Convert(2)
	assert.Equal(t, "II", convertedNumber, "2 = II")
}

func TestConvert_3_III(t *testing.T) {

	var convertedNumber = roman.Convert(3)
	assert.Equal(t, "III", convertedNumber, "3 = III")
}

func TestConvert_4_IV(t *testing.T) {

	var convertedNumber = roman.Convert(4)
	assert.Equal(t, "IV", convertedNumber, "4 = IV")
}

func TestConvert_5_V(t *testing.T) {

	var convertedNumber = roman.Convert(5)
	assert.Equal(t, "V", convertedNumber, "5 = V")
}
