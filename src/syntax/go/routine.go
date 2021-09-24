package goroutine

import (
	"fmt"
	"time"
)

func Go_routine_example() {
	channel := make(chan int)

	go wait_a_long_time(6000000000, channel)
	time.Sleep(1000000)
	fmt.Println("Here we back in the function that called the wait")
	time.Sleep(300000000)
	fmt.Println("We are now free to work on other stuff as we wait")
	time.Sleep(3000000000)
	fmt.Println("This is useful since all these prints can run")
	fmt.Println("While something that takes a long time runs")
	time.Sleep(500000000)
	fmt.Println("in the background")
	<-channel
	fmt.Printf("We know have %v sent back from function\n", channel)
	fmt.Println("spent 60 seconds waiting")
	fmt.Println("This function is now done")

}

func wait_a_long_time(time_to_sleep int, channel chan int) {
	fmt.Println("After me we are waiting for 60 seconds then returning")
	time.Sleep(time.Duration(time_to_sleep))
	fmt.Println("The long wait is now over")
	channel <- time_to_sleep
}
