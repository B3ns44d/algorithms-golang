package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(fizzBuzz(i))
	}
}

func fizzBuzz(i int) string {
	if i%15 == 0 {
		return "FizzBuzz"
	} else if i%3 == 0 {
		return "Fizz"
	} else if i%5 == 0 {
		return "Buzz"
	} else {
		return fmt.Sprintf("%d", i)
	}
}
