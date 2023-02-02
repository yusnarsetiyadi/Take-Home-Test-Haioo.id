package counteditstring

import "fmt"

func CountEditString(string1, string2 string) {
	edit := 0
	for _, v := range string1 {
		for _, w := range string2 {
			if w != v {
				w = v
				edit += 1
			}
		}
	}
	if edit == 1 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}

/*
	How to run? Run the main function as follows:

	func main() {
		CountEditString("telkom", "telecom") //false
		CountEditString("telkom", "tlkom")   //true
	}

*/
