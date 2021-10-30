package main

import "fmt"

func sayHello() string {
	return "hello"
}

func main() {
	fmt.Printf("%v", sayHello())
}
