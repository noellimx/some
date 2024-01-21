package main

import "some/bahttp"

func main() {
	r2 := bahttp.NewEngine()
	_ = r2.Run(":8081")
}
