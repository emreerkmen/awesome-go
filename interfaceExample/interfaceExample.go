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

	interfaceNilReceiver()
	emptyInterfaceExample()
	typeAssertionsExample()

	var i interface{}
	i = "hello"
	typeSwitchesExample(i)
	i = 5
	typeSwitchesExample(i)
	i = true
	typeSwitchesExample(i)

	stringerExample()

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
	if v == nil {
		fmt.Println("<nil>")
		return 0
	}
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//Under the hood, interface values can be thought of as a tuple of a value and a concrete type:
//
//(value, type)
func describe(a Abser) {
	fmt.Printf("(%v, %T)\n", a, a)
}

func interfaceNilReceiver() {

	fmt.Println("Nil receiver example is started.")

	var abser Abser
	//A nil interface value holds neither value nor concrete type
	describe(abser)
	//Calling a method on a nil interface is a run-time error because there is no type inside
	//the interface tuple to indicate which concrete method to call.
	//abser.Abs()

	var vertex *Vertex
	//Note that an interface value that holds a nil concrete value is itself non-nil.
	abser = vertex
	describe(abser)
	abser.Abs()

	abser = &Vertex{1, 2}
	describe(abser)
	abser.Abs()
}

func emptyInterfaceExample() {
	fmt.Println("Empty interface example is started.")

	//An empty interface may hold values of any type. (Every type implements at least zero methods.)
	//Empty interfaces are used by code that handles values of unknown type.
	//For example, fmt.Print takes any number of arguments of type interface{}.
	var emptyInterface interface{}
	describeEmptyInterface(emptyInterface)
}

func describeEmptyInterface(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func typeAssertionsExample() {
	fmt.Println("Type assertion example is started.")

	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64) // no panic
	fmt.Println(f, ok)

	//f = i.(float64) // panic
	fmt.Println(f)
}

func typeSwitchesExample(i interface{}) {

	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func stringerExample() {
	fmt.Println("Stringer interface example is started.")

	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
