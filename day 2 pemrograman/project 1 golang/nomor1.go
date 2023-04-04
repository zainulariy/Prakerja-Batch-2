package main

import "fmt"

func main() {
	var number int
	fmt.Print("Masukkan sebuah bilangan: ")
	fmt.Scan(&number)

	isPrime := true

	for i := 2; i <= number/2; i++ {
		if number%i == 0 {
			isPrime = false
			break
		}
	}

	if isPrime {
		fmt.Printf("%d adalah bilangan prima\n", number)
	} else {
		fmt.Printf("%d bukan bilangan prima\n", number)
	}
}
