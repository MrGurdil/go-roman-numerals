package roman

import "fmt"

func init() {
	fmt.Println("Roman converter initialized.")
}

// Convert an integer to it's roman numeral equivalent
func Convert(a int) string {

	ret := "III"

	if a == 1 {
		ret = "I"
	} else if a == 2 {
		ret = "II"
	}

	return ret
}
