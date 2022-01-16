package goRoutineExample

import (
	"fmt"
	"time"
)

func GoRoutineExample() {
	fmt.Println("Go Routine example is started.")

	//A goroutine is a lightweight thread managed by the Go runtime.
	go say("Go routine hello")
	go say("Go routine hello2")
	say("Normal hello")   // This is normal hello, but it will start to execute because previous commands are go routine.
	say("Normal hello 2") // This is normal hello also, but it will wait to start until previous command start
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
