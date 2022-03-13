package main

import "fmt"

func main() {
	// string
	fmt.Println("hello")           // hello
	fmt.Println("hello" + "world") // helloworld
	fmt.Println("hello", "world")  // hello world

	// Number
	fmt.Println("1+1=", 1+1)         // 1+1 = 2
	fmt.Println("8.0/3.0=", 8.0/3.0) // 8.0/3.0 = 2.6666666666666665
	fmt.Println("9.0/3.0=", 9.0/3.0) // 9.0/3.0 = 3

	// boolean
	fmt.Println(true)  // true
	fmt.Println(false) // false
	fmt.Println(nil)   // <nil> == null
}
