//Every program start with package named with main
package main

import (
	"fmt"
	"math"
)

var c, python, java bool

const Pi = 3.14 //Can not be declared using with ":="
//An untyped numeric constant takes the type needed by its context.

func main() {
	var i int
	fmt.Println(i)
	fmt.Println("Hello world.")
	fmt.Println(add(2, 3))
	fmt.Println(addWithEasyArg(2, 3))
	a, b := 2, 3
	fmt.Printf("Before swap a:%v b:%v\n", a, b)
	a, b = swap(a, b)
	fmt.Printf("After swap a:%v b:%v\n", a, b)
	fmt.Println(swapTimes2NakedReturn(a, b))
	fmt.Println(python)
	fmt.Println(c, python, java)

	var csharp, pythons, java = true, false, "no!"
	fmt.Println(csharp, pythons, java)

	//Unlike in C, in Go assignment between items of different type requires an explicit conversion.
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	//Or
	//i := 42
	//f := float64(i)
	//z := uint(f)
	fmt.Println(x, y, z)
}

//Basic func
func add(x int, y int) int {
	return x + y
}

//If you use more than one same type arg. you don't have to declare type but last one
func addWithEasyArg(x, y int) int {
	return x + y
}

// Func can return more than one result
func swap(a, b int) (int, int) {
	return b, a
}

//In go you can name return values
//Naked returns is much better for short func
func swapTimes2NakedReturn(a, b int) (x, y int) {
	x = b * 2
	y = a * 2
	return
}
