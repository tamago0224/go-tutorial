package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	var a []int = primes[0:5]
	fmt.Println(s)
	fmt.Println(a)
}
