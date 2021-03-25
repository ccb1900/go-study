package main

import (
	"fmt"
	"ppp/util"
)

type Message struct {
	str  string
	wait chan bool
}

func fanIn(inputs ...<-chan Message) <-chan Message {
	c := make(chan Message)
	for i := range inputs {
		input := inputs[i]

		go func() {
			for {
				c <- <-input
			}
		}()
	}

	return c
}

func boring(msg string) <-chan Message {
	c := make(chan Message)
	waitForIt := make(chan bool)

	go func() {
		util.ILoop(func(i int) {
			c <- Message{
				str:  fmt.Sprintf("%d %s", i, msg),
				wait: waitForIt,
			}

			util.Sleep()

			<-waitForIt
		})
	}()

	return c
}

func main() {
	c := fanIn(boring("hello"), boring("world"))
	util.Loop(func(i int, n int) {
		msg1 := <-c
		fmt.Println(msg1.str)
		msg2 := <-c
		fmt.Println(msg2.str)

		msg1.wait <- true
		msg2.wait <- true
	}, 15)

	fmt.Println("i am leaving")
}
