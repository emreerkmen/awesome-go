package methodExample

import (
	"fmt"
	"math"
)

func MethodExample() {
	fmt.Println("Method example is started.")
	v := VertexForMethod{X: 1, Y: 2}
	result := v.Abs()
	fmt.Printf("Method version: %f\n", result)
	result2 := absJustFunction(v)
	fmt.Printf("Func version %f\n", result2)
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.absNoStructTypes())
	pointerReceiverExample()
}

type VertexForMethod struct {
	X, Y float64
}

//A method is a function with a special receiver argument.
//The receiver appears in its own argument list between the func keyword and the method name.

func (v VertexForMethod) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//Function version of above method.
func absJustFunction(v VertexForMethod) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//You can declare a method on non-struct types, too.

type MyFloat float64

func (f MyFloat) absNoStructTypes() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// You can declare pointer receiver as well as value receiver
// Since methods often need to modify their receiver, pointer receivers are more common than value receivers.

func (v *VertexForMethod) scalePointerReceiver(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
	fmt.Println("V: ", v)
}

func (v VertexForMethod) scaleValueReceiver(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
	fmt.Println("V: ", v)
}

func pointerReceiverExample() {
	fmt.Println("Pointer and value receiver example is started.")

	v := VertexForMethod{X: 3, Y: 4}
	fmt.Println("Original V: ", v)
	fmt.Println("Value receiver result:")
	v.scaleValueReceiver(5)
	fmt.Println("Original V: ", v)
	fmt.Println("Pointer receiver result:")
	//For the statement v.scalePointerReceiver(5), even though v is a value and not a pointer,
	//the method with the pointer receiver is called automatically.
	//That is, as a convenience, Go interprets the statement v.scalePointerReceiver(5) as (&v).scalePointerReceiver(5)
	//since the scalePointerReceiver method has a pointer receiver.
	//Same things happens in the reverse direction for value receiver
	v.scalePointerReceiver(5)
	fmt.Println("Original V: ", v)

	//Choosing a value or pointer receiver
	//The first is so that the method can modify the value that its receiver points to.
	//The second is to avoid copying the value on each method call. This can be more
	//efficient if the receiver is a large struct, for example.
}
