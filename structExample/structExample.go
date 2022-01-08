package structExample

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func StructExample() {
	fmt.Println("Struct example is started.")
	v := Vertex{1, 2}
	fmt.Println(v)
	v.Y = 4
	v.X = 2
	fmt.Println(v)
	pointerV := &v
	pointerV.X = 4 // When accessing struct element in pointer, you don't have to use * operand
	fmt.Println(*pointerV)
	fmt.Println(v1, p, v2, v3)
}
