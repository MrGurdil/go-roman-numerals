package romantests

import (
	"converter/roman"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestRomanConverter is the main test function for the converter/roman package
func TestRomanConverter(t *testing.T) {

	// Test that Convert(1) returns I
	var convertedNumber1 = roman.Convert(1)
	assert.Equal(t, "I", convertedNumber1, "1 = I")

	// Test that Convert(2) returns II
	var convertedNumber2 = roman.Convert(2)
	assert.Equal(t, "II", convertedNumber2, "2 = II")
}
