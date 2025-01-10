package main

import (
	"fmt"

	"github.com/morjuax/go-desde-0/variables"
)

func main() {
	state, text := variables.ConvertToText(1588)

	fmt.Println(state, text)
}