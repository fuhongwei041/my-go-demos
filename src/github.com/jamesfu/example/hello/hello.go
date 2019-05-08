package main

import (
	"fmt"
	"github.com/jamesfu/example/stringutil"
)

func main() {
	fmt.Printf("hello %v\n", "james.h.fu@g.com")
	fmt.Printf("hello world.%s\n", stringutil.Reverse("cba"))
}
