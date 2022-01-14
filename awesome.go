//Every program start with package named with main
package main

import (
	"awesome-go/arrayExample"
	"awesome-go/errorExample"
	"awesome-go/functionExample"
	"awesome-go/interfaceExample"
	"awesome-go/mapExample"
	"awesome-go/methodExample"
	"awesome-go/pointerExample"
	"awesome-go/sliceExample"
	"awesome-go/structExample"
	"fmt"
	"math"
	"runtime"
	"time"
)

var c, python, java bool

const Pi = 3.14 //Const's variables can not be declared using with ":="
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
	forExample()
	fmt.Println(ifExamplePow(3, 2, 10))
	switchExample()
	deferExample()
	deferExample2()
	pointerExample.Pointers()
	structExample.StructExample()
	arrayExample.ArrayExample()
	sliceExample.SliceExample()
	mapExample.MapExample()
	functionExample.FunctionExample()
	methodExample.MethodExample()
	interfaceExample.InterfaceExample()
	errorExample.ErrorExample()
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

func forExample() {
	sum := 0
	//There is no need parentheses like in other programming languages
	for i := 0; i < 10; i++ {
		sum += i
		//There is no need parentheses for if condition
		if sum > 10 {
			fmt.Printf("Sum is higher than 10. Sum: %v\n", sum)
		}
	}
	fmt.Println(sum)

	//init and post statement are optional. So it is basically "while" statement
	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	//infinite loop
	/*
		for {
		}
	*/
}

func ifExamplePow(x, n, lim float64) float64 {
	//Like for statement you can put nit statement for if in go
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func switchExample() {
	//n effect, the break statement that is needed at the end of each case
	//in other languages is provided automatically in Go.
	//Because of that there is no need break everywhere
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	//Switch cases evaluate cases from top to bottom, stopping when a case succeeds.
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	//If there is no condition it is means that "switch true"
	//This construct can be a clean way to write long if-then-else chains.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func deferExample() {
	defer fmt.Println("after surrounding func return")

	fmt.Println("You will see")
	fmt.Println("defer result ")
}

func deferExample2() {
	fmt.Println("counting")

	//When a function returns, its deferred calls are executed in last-in-first-out order like a stack
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
