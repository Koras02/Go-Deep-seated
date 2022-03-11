// 2강
// Printf()

/*
   표현   설명    비교
   %d     정수    decimal
   %f     실수    float
   %s     문자열  string
   %T     자료형  Type
   %v     {값 목록}  values
   %+v    {키: 값 목록} values+
   %#v    자료형{키: 값 목록}  values++


*/

package main

import "fmt"

// type Vertex struct {
// 	X int
// 	Y int
// }

// func main() {
// 	vertex := Vertex{1, 2}

// 	fmt.Println(vertex)        // {1 2}
// 	fmt.Printf("%v\n", vertex) // {1 2}

// 	fmt.Printf("%+v\n", vertex) // {X : 1, Y: 2}
// 	fmt.Printf("%#v\n", vertex) // main.Vertex{X:1, Y:2}
// 	fmt.Printf("%T\n", vertex)  // main.Vertex
// }

func main() {
	friuts := []string{"apple", "banana", "mango"}

	fmt.Println(friuts)         // [apple banana mango]
	fmt.Printf("%v\n", friuts)  // [apple banana mango]
	fmt.Printf("%+v\b", friuts) // [apple banana mango]

	fmt.Printf("%q\n", friuts)  // ["apple", "banana", "mango"]
	fmt.Printf("%#v\n", friuts) // []string{"apple", "banana", "mango"}
	fmt.Printf("%T\n", friuts)  // []string
}
