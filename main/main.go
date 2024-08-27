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
	handler := NewHandler(calc.Addition{}, os.Stdout)
	err := handler.Handle(os.Args[1:])
	if err != nil {
		panic(err)
	}
}

type Handler struct {
	calculator Calculator
	output     io.Writer
}

func NewHandler(calculator Calculator, output io.Writer) *Handler {
	return &Handler{calculator: calculator, output: output}
}

func (this *Handler) Handle(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("expected 2 inputs, got %d", len(args))
	}

	a, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}
	b, err := strconv.Atoi(args[1])
	if err != nil {
		return err
	}
	this.calculator.Calculate(a, b)
	_, err = fmt.Fprintf(this.output, "Result: %d\n", this.calculator.Calculate(a, b))
	if err != nil {
		return err
	}
	return nil
}
