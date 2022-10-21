package currency

import "fmt"

type f64 float64

type CURR string

const (
	USD CURR = "USD"
	JPY CURR = "JPY"
	GBP CURR = "GBP"
	CNY CURR = "CNY"
	SGD CURR = "SGD"
	MYR CURR = "MYR"
	EUR CURR = "EUR"
)

var cmapFromEUR map[CURR]f64 = map[CURR]f64{
	USD: 1.1318,
	JPY: 121.05,
	GBP: 0.90630,
	CNY: 7.9944,
	SGD: 1.5743,
	MYR: 4.8390,
	EUR: 1,
}

func currConvertFromEUR(c CURR, v f64) f64 {
	return cmapFromEUR[c] * v
}

func currConvertToEUR(c CURR, v f64) f64 {
	return cmapFromEUR[c] / v
}

func currConvertFromTo(c CURR, d CURR, v f64) f64 {

	return v / cmapFromEUR[c] * cmapFromEUR[d]
}

func Activity() {

	var (
		fromC CURR
		toC   CURR

		amt f64
	)

	fmt.Println("From Currency: ")

	fmt.Scanln(&fromC)

	fmt.Println("Amount: ")

	fmt.Scanln(&amt)

	fmt.Println("To Currency: ")

	fmt.Scanln(&toC)

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("[Recovered] Oops something went wrong with calculating...")
		}
	}()

	euros := currConvertToEUR(fromC, amt)

	toAmt := currConvertFromEUR(toC, euros)

	toAmt2 := currConvertFromTo(fromC, toC, amt)
	fmt.Printf("%s %0.8f -> %s %0.8f %0.8f", fromC, amt, toC, toAmt, toAmt2)
}
