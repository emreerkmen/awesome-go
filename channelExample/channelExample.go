package channelExample

import (
	"fmt"
	"sync"
	"time"
)

func ChannelExample() {
	fmt.Println("Channel example is started.")

	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// channels are like message buses but in programing language level. You can use this channels between go routines.
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	//By default, sends and receives block until the other side is ready.
	//This allows goroutines to synchronize without explicit locks or condition variables.
	x, y := <-c, <-c // receive from c  // It will wait (or block execution) until previous sends are finished.
	fmt.Println("Receive is finished.")
	fmt.Println(x, y, x+y)

	channelsDeadlockExample()
	bufferedChannelExample()
	closeChannelExample()
	selectExample()
	defaultSelectionExample()
	mutualExclusionExample()
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		fmt.Println(v)
		sum += v
	}
	fmt.Println("Send value")
	c <- sum // send sum to c
}

func channelsDeadlockExample() {
	fmt.Println("Channel deadlock example is started.")

	// It will send error because in 48. row the execution stop but there is no any go routine
	//that continue to run to receive from channel any value.
	/*ch1 := make(chan int)
	ch1 <- 1 // block
	ch1 <- 2
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)*/

	// It will send error because in 55. row the execution stop but there is no any go routine
	//that continue to run to receive from channel any value.
	/*ch1 := make(chan int)
	ch1 <- 1
	fmt.Println(<-ch1)
	ch1 <- 2
	fmt.Println(<-ch1)*/

	// It will run without error
	ch2 := make(chan int)
	go func() {
		ch2 <- 1
	}()
	fmt.Println(<-ch2)
	go func() {
		ch2 <- 2
	}()
	fmt.Println(<-ch2)
}

func bufferedChannelExample() {
	fmt.Println("Buffered channel example is started.")

	// Look at example that at 61. row
	// With second argument we can give a buffer to channel to not block execution.
	// Sends to a buffered channel block only when the buffer is full.
	// Receives block when the buffer is empty.
	ch1 := make(chan int, 2)
	ch1 <- 1 // not block
	ch1 <- 2 // not block
	// ch1 <- 3 // block and send deadlock error
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)
}

func closeChannelExample() {
	fmt.Println("Close channel example is started.")

	// A sender can close a channel to indicate that no more values will be sent.
	// Receivers can test whether a channel has been closed by assigning a second
	// parameter to the receive expression

	//ok is false if there are no more values to receive and the channel is closed.
	//v, ok := <-ch

	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	//The loop for i := range c receives values from the channel repeatedly until it is closed.
	for i := range c {
		fmt.Println(i)
	}

	//Note: Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.
	//
	//Another note: Channels aren't like files; you don't usually need to close them.
	//Closing is only necessary when the receiver must be told there are no more values coming,
	//such as to terminate a range loop.
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		fmt.Println("Send to channel x: ", x)
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func selectExample() {
	fmt.Println("Select example is started.")

	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacciSelect(c, quit)
}

func fibonacciSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		// A select blocks until one of its cases can run, then it executes that case.
		// It chooses one at random if multiple are ready.
		select {
		case c <- x: // it is run because a value is received from c
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func defaultSelectionExample() {
	fmt.Println("Default selection example is started.")

	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			// It will run infinitely if there is no other case
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
func mutualExclusionExample() {
	fmt.Println("Mutual Exclusion example is started.")

	c := SafeCounter{v: make(map[string]string)}
	characters1 := [5]string{"a", "b", "c", "d", "e"}
	characters2 := [5]string{"f", "g", "h", "i", "j"}
	for i, v := range characters1 {
		fmt.Println("Index: ", i)
		go c.Inc("somekey", v)
		go c.Inc("somekey", characters2[i])
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]string
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string, value string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key] = c.v[key] + value
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) string {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}
