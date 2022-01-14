package errorExample

import (
	"fmt"
	"strconv"
	"time"
)

func ErrorExample() {
	fmt.Println("Error example is started.")

	//Go programs express error state with error values.
	//The error type is a built-in interface similar to fmt.Stringer:
	//
	//type error interface {
	//    Error() string
	//}

	//Functions often return an error value, and calling code should handle errors
	//by testing whether the error equals nil.
	i, err := strconv.Atoi("42")
	if err != nil {
		fmt.Printf("couldn't convert number: %v\n", err)
		return
	}
	fmt.Println("Converted integer:", i)

	errorInterfaceExample()
	errorReturnExample()
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func errorInterfaceExample() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

type ErrNegativeSqrt float64

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return x, nil
}

func (en ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("at %f", en)
}

func errorReturnExample() {
	fmt.Println("Error return example is started.")

	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
