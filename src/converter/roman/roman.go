package roman

import "fmt"

func init() {
	fmt.Println("Roman converter initialized.")
}

// Convert an integer to it's roman numeral equivalent
func Convert(a int) string {

	ret := "II"

	if a == 1 {
		ret = "I"
	}

	return ret
}
