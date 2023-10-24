package main

import (
	"fmt"
	"time"
)

func main() {
	go myProcess("A")
	go myProcess("B")

	time.Sleep(2 * time.Second)

	myfirstChannel := make(chan string)

	go myProcessWithChannel("C", myfirstChannel)

	result := <-myfirstChannel
	fmt.Println(result)
	close(myfirstChannel)

	channelA := make(chan string)
	channelB := make(chan string)

	go myProcessWithChannel("D", channelA)
	go myProcessWithChannel("E", channelB)

	resultA := <-channelA
	resultB := <-channelB

	fmt.Println(resultA)
	fmt.Println(resultB)

	close(channelA)
	close(channelB)

}

func myProcess(p string) {
	i := 0

	for i < 15 {
		time.Sleep(150 * time.Millisecond)
		i++
		fmt.Printf("%s - %d\n", p, i)
	}

}

func myProcessWithChannel(p string, c chan string) {
	i := 0

	for i < 15 {
		time.Sleep(150 * time.Millisecond)
		i++
		fmt.Printf("process: %s - num: %d\n", p, i)
	}

	c <- "Channel finish - OK"
}
