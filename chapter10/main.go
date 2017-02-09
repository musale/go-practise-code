// How do you specify the direction of a channel type?
// by using <- i.e.func pinger(c chan<- string) will only be sent to

package main

import (
	"fmt"
	"time"
)

func Sleep(sec int32) {
	go func() {
		for {
			select {
			case <-time.After(time.Duration(sec)):
				fmt.Println("Sleeping=========================================")
			default:
				fmt.Println("Doing Stuff")
			}
		}
	}()
}

func main() {
	Sleep(34)
	var input string
  fmt.Scanln(&input)
}
