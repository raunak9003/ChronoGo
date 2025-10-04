package main

import (
	"fmt"
	"sync"
)

/*

simple declaring a channel and performing operation

func main() {
	ch := make(chan int)
	go multiplyWithChannel(ch)
	ch <- 10
}

func multiplyWithChannel(ch chan int) {
	fmt.Println(10 * <-ch)
}
*/

/*

//closing a channel

func main() {
	ch := make(chan int)
	close(ch)
	elem, ok := <-ch

	// now on closing if there is any element in channel then the variable elem fetches it and ok is the bool value
	// if ok-> false channel is closed
	fmt.Println(ok, elem)

}
*/

//for loop

func main() {
	c := make(chan string)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			resp, ok := <-c
			if ok == false {
				fmt.Println("channel closed")
				break
			}
			fmt.Println("channel opened", resp)

		}
	}()
	initStrings(c)
	wg.Wait()
}

func initStrings(ch chan string) {
	for i := 0; i < 3; i++ {
		ch <- "hii"
	}
	close(ch)
}
