package moneyfraction

import "fmt"

func MoneyFraction(money int) {
	var fraction = []int{100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100}
	for i := 0; i < len(fraction); i++ {
		total := money / fraction[i]
		money -= (fraction[i] * total)
		if total == 0 {
			continue
		}
		fmt.Printf("%v : %d\n", fraction[i], total)
		if money < fraction[len(fraction)-1] && money != 0 {
			defer fmt.Printf("%v : %d\n", fraction[len(fraction)-1], 1)
		}
	}

}

/*
	How to run? Run the main function as follows:

	func main() {
	var n int
	fmt.Print("input money:")
	fmt.Scan(&n)
	MoneyFraction(n)
}

*/
