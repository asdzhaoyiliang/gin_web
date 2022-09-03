package main

import (
	"fmt"
	"github.com/aiwen/aiwen-go-gin/demo"
)

func main() {
	fmt.Println("hello world")
	c := demo.Add(3, 4)
	fmt.Println("c = ", c)
}
