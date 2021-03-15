package main

import (
	"fmt"

	"github.com/wends155/greetings"
)

type user struct {
	name string
}

func ping(pings chan<- user, msg user) {
	pings <- msg
}

func pong(pings <-chan user, pongs chan<- user) {
	msg := <-pings
	pongs <- msg
}

func main() {

	message := greetings.Hello("Wendell")
	fmt.Println(message)

	pings := make(chan user, 1)
	pongs := make(chan user, 1)

	go ping(pings, user{name: "hello"})
	go pong(pings, pongs)

	fmt.Println(<-pongs)

}
