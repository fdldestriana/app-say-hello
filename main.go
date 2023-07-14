package main

import (
	"fmt"

	go_say_hello "github.com/fdldestriana/go-say-hello/v2"
)

func main() {
	var helloMessage string = go_say_hello.SayHello("Eko")
	fmt.Println(helloMessage)
}
