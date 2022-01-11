package interfaceExample

import (
	"awesome-go/methodExample"
	"fmt"
	"math"
)

func InterfaceExample() {
	fmt.Println("Interface example is started.")

	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}
	//Implicit interfaces decouple the definition of an interface from its implementation,
	//which could then appear in any package without prearrangement.
	vm := methodExample.VertexForMethod{X: 3, Y: 4}

	a = f // a MyFloat implements Abser
	fmt.Println(a.Abs())
	a = &v // a *Vertex implements Abser
	fmt.Println(a.Abs())
	a = vm
	fmt.Println(vm.Abs())

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	//a = v

	describe(f)
	describe(&v)
	describe(vm)

}

type Abser interface {
	Abs() float64
}

type MyFloat float64

// Abs Interfaces are implemented implicitly
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

// Abs Interfaces are implemented implicitly
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//Under the hood, interface values can be thought of as a tuple of a value and a concrete type:
//
//(value, type)
func describe(a Abser) {
	fmt.Printf("(%v, %T)\n", a, a)
}
