package main

import (
	"errors"
	"fmt"

	e "github.com/gaplyk/go-swagger-issue/errors"
)

func main() {
	err := errors.New("test")
	fmt.Println(err)

	err2 := e.New("test")

	fmt.Println(err2)
}
