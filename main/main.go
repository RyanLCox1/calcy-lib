package main

import (
	"errors"
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
		return fmt.Errorf("%w: two args required (you provided %d)", errWrongNumArgs, len(args))
	}

	a, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("%w: first arg (%s) %w", errMalformedArgument, args[0], err)
	}
	b, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("%w: second arg (%s) %w", errMalformedArgument, args[1], err)
	}
	c := this.calculator.Calculate(a, b)
	_, err = fmt.Fprintf(this.output, "Result: %d\n", c)
	if err != nil {
		return fmt.Errorf("%w write failed, %w", errOutputWriteErr, err)
	}
	return nil
}

var (
	errWrongNumArgs      = errors.New("usage: calc <a> <b>")
	errMalformedArgument = errors.New("invalid argument")
	errOutputWriteErr    = errors.New("output writer err")
)
