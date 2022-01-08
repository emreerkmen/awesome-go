package mapExample

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func MapExample() {
	fmt.Println("Map example is started.")

	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	mapLiteralsExample()
	mutatingMapExample()
}

func mapLiteralsExample() {
	fmt.Println("Map literal example is started.")

	var mapLiterals = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}

	fmt.Println(mapLiterals)

	//If the top-level type is just a type name, you can omit it from the elements of the literal.
	var easyMapLiterals = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}

	fmt.Println(easyMapLiterals)
}

func mutatingMapExample() {
	fmt.Println("Mutating map example is started.")

	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"] // nice skill
	fmt.Println("The value:", v, "Present?", ok)
}
