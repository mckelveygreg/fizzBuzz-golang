// fizzbuzz test from https://www.golang-book.com/books/intro/5
package main

import (
	"fmt"
)

func main() {
	var response string
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			response = "Fizz"
		}
		if i%5 == 0 {
			response += "Buzz"
		}
		if response == "" {
			//response = i
		}
		if response != "" {
			fmt.Println(response)
		} else {
			fmt.Println(i)
		}
		response = ""
	}
}
