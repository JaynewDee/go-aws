package main

import "fmt"

func main() {
	fmt.Println("Hello from Lambda module!")
}

func Init() {
	fmt.Println("Hello from go-aws/lambda!")
}
