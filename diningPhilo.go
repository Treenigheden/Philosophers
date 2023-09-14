package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
By having every philosopher start by picking up either the right or left fork depending
on their index, we ensure that two forks are only ever picked up by a philosopher as their
second fork, meaning that the two forks adjacent to the philosopher are about to become
available. In other words, these forks are either available or in use for eating. It is
therefore impossible that all philosophers are waiting on forks at the same time and thus
we avoid deadlock.
*/
func Philosopher(number int, name string, ch1 chan bool, ch2 chan bool) {
	var eatings int = 0
	for eatings < 3 {
		fmt.Println(name, "is thinking")
		var ran int = rand.Intn(9) * 100
		time.Sleep(time.Millisecond * time.Duration(ran))
		if (number % 2) == 0 {
			var left bool = <-ch1
			if left {
				var right bool = <-ch2
				if right {
					eatings++
					fmt.Println(name, "is eating for the ", eatings, " time")
					ch1 <- false
					ch2 <- false
				} else {
					ch1 <- left
					ch2 <- right
				}
			} else {
				ch1 <- left
			}
		} else {
			var right bool = <-ch2
			if right {
				var left bool = <-ch1
				if left {
					eatings++
					fmt.Println(name, "is eating for the ", eatings, " time")
					ch1 <- false
					ch2 <- false
				} else {
					ch2 <- right
					ch1 <- left
				}
			} else {
				ch2 <- right
			}
		}
	}

}

// forks only writes and reads status from channels (discards read status)
func Fork(ch chan bool) {
	for {
		ch <- true
		<-ch
	}
}

func main() {

	var names [5]string
	names[0] = "Socrates"
	names[1] = "Confucius"
	names[2] = "Nietzche"
	names[3] = "Hegel"
	names[4] = "Kierkegaard"

	//creating channels
	var ch = []chan bool{
		make(chan bool),
		make(chan bool),
		make(chan bool),
		make(chan bool),
		make(chan bool),
	}

	//creating forks
	for i := 0; i < 5; i++ {
		go Fork(ch[i])
	}

	//creating philosophers and assigning the correct channels
	for i := 0; i < 5; i++ {
		go Philosopher(i, names[i], ch[i], ch[(i+1)%5])
	}
	for {
	}
}
