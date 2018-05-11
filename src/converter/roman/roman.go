package roman

import "fmt"

func init() {
	fmt.Println("Roman converter initialized.")
}

// Convert an integer to it's roman numeral equivalent
func Convert(a int) string {

	ret := "V"

	if a == 4 {
		ret = "I" + ret
	} else if a < 4 {
		ret = ""
		for a > 0 {
			ret += "I"
			a--
		}
	}

	return ret
}
