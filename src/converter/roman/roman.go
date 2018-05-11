package roman

import "fmt"

func init() {
	fmt.Println("Roman converter initialized.")
}

// Convert an integer to it's roman numeral equivalent
func Convert(a int) string {

	var ret string

	if a == 4 {
		ret = "IV"
	} else if a < 4 {
		for a > 0 {
			ret += "I"
			a--
		}
	}

	return ret
}
