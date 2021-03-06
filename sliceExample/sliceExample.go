package sliceExample

import (
	"fmt"
	"strings"
)

func SliceExample() {
	fmt.Println("Slice example is started.")
	//A slice is a dynamically-sized, flexible view into the elements of an array
	//In practice, slices are much more common than arrays.
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4] // include 1. element but exclude 4. element
	fmt.Println(s)

	//Slices are like references to arrays
	//A slice does not store any data, it just describes a section of an underlying array.
	//Changing the elements of a slice modifies the corresponding elements of its underlying array.
	//Other slices that share the same underlying array will see those changes.

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	sliceLiteralsExample()
	sliceLengthAndCapacityExample()
	nilSliceExample()
	creatingSliceWithMake()
	sliceOfSlice()
	appendingSliceExample()
	rangeForSlice()
}

func sliceLiteralsExample() {
	fmt.Println("Slice literals example is started.")

	//This creates the "q := [6]int{2, 3, 5, 7, 11, 13}" array implicitly, then builds a slice that references
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	defaultSlice := []int{2, 3, 5, 7, 11, 13}

	defaultSlice = defaultSlice[1:4] // low and high not used as default
	fmt.Println(defaultSlice)

	defaultSlice = defaultSlice[:2] // low is used to as default, simply do not put specific number for low. default of low is 0
	fmt.Println(defaultSlice)

	defaultSlice = defaultSlice[1:] // high is used to as default, simply do not put specific number for high. default of high is length
	fmt.Println(defaultSlice)
}

func sliceLengthAndCapacityExample() {
	fmt.Println("Slice length and capacity example is started.")
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:5]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func nilSliceExample() {
	fmt.Println("Nil slice example is started.")
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

func creatingSliceWithMake() {
	fmt.Println("Make example is started.")

	//The make function allocates a zeroed array and returns a slice that refers to that array
	a := make([]int, 5)
	printSliceMake("a", a)

	//To specify a capacity, pass a third argument to make
	b := make([]int, 0, 5)
	printSliceMake("b", b)

	c := b[:2]
	printSliceMake("c", c)

	d := c[2:5]
	printSliceMake("d", d)
}

func printSliceMake(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func sliceOfSlice() {
	fmt.Println("Slice of slice example is started.")

	//Slices can contain any type, including other slices.
	//Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	//The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func appendingSliceExample() {
	fmt.Println("Appending slice example is started.")

	var s []int
	printSlice(s)

	//If the backing array of s is too small to fit all the given values a bigger array will be allocated.
	//The returned slice will point to the newly allocated array.
	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func rangeForSlice() {
	fmt.Println("Range example is started.")

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	//to skip value;   index, _   := range pow
	//to skip index;     _, value := range pow
	//to omit just index:   index := range pow
	for index, value := range pow {
		fmt.Printf("2**%d = %d\n", index, value)
	}
}
