package functionExample

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func FunctionExample() {
	fmt.Println("Functions as values example is started.")

	//Functions are values too. They can be passed around just like other values.
	//Function values may be used as function arguments and return values.
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	functionClosuresExample()
}

func functionClosuresExample() {
	fmt.Println("Function closures example is started.")

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("Iteration i:%d |", i)
		fmt.Println("Result:",
			pos(i),
			neg(-2*i),
		)
	}
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		fmt.Printf("inner x paramter:%d |", x)
		sum += x
		return sum
	}
}
