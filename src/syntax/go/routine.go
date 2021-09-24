package goroutine

import (
	"fmt"
	"time"
)

func Go_routine_example() {
	channel := make(chan int)

	go wait_a_long_time(10, channel)
	fmt.Println("Here we back in the function that called the wait")
	<-channel
	fmt.Printf("We know have %v sent back from waiting", channel)
	fmt.Println("This function is now done")

}

func wait_a_long_time(time_to_sleep int, channel chan int) {
	fmt.Println("After me we are waiting for 10 seconds then returning")
	time.Sleep(time.Duration(time_to_sleep))
	fmt.Println("The long wait is now over")
	channel <- time_to_sleep
}
