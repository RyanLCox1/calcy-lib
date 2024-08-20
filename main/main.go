package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/smartystreets/calcy-lib/calc"
)

type Calculator interface {
	Calculate(a, b int) int
}

func main() {
	var (
		inputs              = os.Args[1:]
		addition Calculator = calc.Addition{}
		output   io.Writer  = os.Stdout
	)

	if len(inputs) != 2 {
		panic(fmt.Errorf("expected 2 inputs, got %d", len(inputs)))
	}

	a, err := strconv.Atoi(inputs[0])
	if err != nil {
		panic(err)
	}
	b, err := strconv.Atoi(inputs[1])
	if err != nil {
		panic(err)
	}
	addition.Calculate(a, b)
	_, err = fmt.Fprintf(output, "Result: %d\n", addition.Calculate(a, b))
	if err != nil {
		panic(err)
	}
}
