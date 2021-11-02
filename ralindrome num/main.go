package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))

}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var reverse int = 0
	var num = x
	for num > 0 {
		reverse = reverse*10 + num%10
		num = num / 10
	}
	return reverse == x
}
