package main

import (
	"fmt"
	"time"
)

func Philosopher(state chan bool, number int) {
	num := number

	var numOfEatings int = 0
	var eating bool = false

	fmt.Println("Philosopher #", num, "has arrived")

	if eating {
		fmt.Println("Philosopher #", num, "is eating")
		numOfEatings++
	} else {
		fmt.Println("Philosopher #", num, "is thinking")
	}

	if numOfEatings == 3 {
		fmt.Println("Philosopher #", num, "is done eating!")
	}
}

func Fork(isUsed chan bool) {
	inUse := <-isUsed

	if inUse {
		fmt.Println("In use")
	} else {
		fmt.Println("Not on use")
	}
}

func main() {

	state := make(chan bool)

	for i := 1; i < 6; i++ {
		go Philosopher(state, i)
		time.Sleep(time.Millisecond * 1000)
	}

	<-state
}
