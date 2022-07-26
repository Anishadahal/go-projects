package main

import "fmt"

//1. Absolute value function
//2. Functions are their own datatype

func test(x string) {
	fmt.Println("Hello world", x)
}

func add(x, y int) (int, int) { //means x int, y int and returns int, int(2 values)
	return x + y, x - y

}
func calc(z, y int) (sum, diff, prod int) {
	defer fmt.Println("DEFERRED")
	sum = z + y
	diff = z - y
	prod = z * y
	fmt.Println("CHECK")
	return
}

// func test2(myFunc func(int) int) func(string) int { //takes function and returns function
func test2(myFunc func(int) int) {
	fmt.Println(myFunc(7))
}


//function closure

func returnFunc(x string) func(){
	return func(){
		fmt.Println(x)
	}
}
func main() {
	// test("Welcome")
	// _, diff := add(10, 20)
	// fmt.Println(diff)

	// s, d, prod := calc(10, 3)
	// fmt.Println(s, d, prod)

	// x := test
	// x(", welcome") // similar to test()

	//Function inside function
	// test := func() {
	// 	fmt.Println("Func inside func")
	// }
	// test()

	// test1 := func(x int) int {
	// 	return x + 1
	// }(10) //calling function right where it is defined -->Inline
	// fmt.Println(test1)
	// test3 := func(x int) int {
	// 	return x + 1
	// }
	// test4 := func(x int) int {
	// 	return x * 5
	// }
	// // test2(test3)
	// test2(test4)

	x := returnFunc("hello")
	x() //equiv to : returnFunc("hello")()
}
