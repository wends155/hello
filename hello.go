package main

import (
	"fmt"

	"github.com/wends155/greetings"
)

func main() {
	message := greetings.Hello("Wendell")
	fmt.Println(message)
}
