package main

import (
	"fmt"

	go_say_hello "github.com/fdldestriana/go-say-hello"
)

func main() {
	var helloMessage string = go_say_hello.SayHello()
	fmt.Println(helloMessage)
}
