package main

import (
	"fmt"

	"github.com/ProjectOrangeJuice/golib-logging/logging"
)

func main() {
	fmt.Println("test")
	logging.Log("test", "hello", nil)
}
