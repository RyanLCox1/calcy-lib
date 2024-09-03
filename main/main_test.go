package main

import (
	"bytes"
	"errors"
	"reflect"
	"testing"

	"github.com/smartystreets/calcy-lib/calc"
)

// assert equal
func assertEqual(t *testing.T, expected, actual any) {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nExpected: [%v]\nActual: [%v]", expected, actual)
	}
}

func assertError(t *testing.T, expected, actual error) {
	t.Helper()
	if !errors.Is(actual, expected) {
		t.Errorf("\nExpected: [%v]\nActual: [%v]", expected, actual)
	}
}

func TestWrongNumArgs(t *testing.T) {
	output := bytes.Buffer{}
	handler := NewHandler(calc.Addition{}, &output)
	err := handler.Handle([]string{"1"})
	assertError(t, errWrongNumArgs, err)
	assertEqual(t, "", output.String())
}

func TestInvalidFirstArg(t *testing.T) {
	output := bytes.Buffer{}
	handler := NewHandler(calc.Addition{}, &output)
	err := handler.Handle([]string{"a", "1"})
	assertError(t, errMalformedArgument, err)
	assertEqual(t, "", output.String())
}

func TestInvalidSecondArg(t *testing.T) {
	output := bytes.Buffer{}
	handler := NewHandler(calc.Addition{}, &output)
	err := handler.Handle([]string{"1", "a"})
	assertError(t, errMalformedArgument, err)
	assertEqual(t, "", output.String())
}

func TestInvalidWrite(t *testing.T) {
	output := &badOutput{}
	handler := NewHandler(calc.Addition{}, output)
	err := handler.Handle([]string{"1", "1"})
	assertError(t, errOutputWriteErr, err)
}

func TestHappyPath(t *testing.T) {
	output := bytes.Buffer{}
	handler := NewHandler(calc.Addition{}, &output)
	err := handler.Handle([]string{"1", "2"})
	assertEqual(t, nil, err)
	assertEqual(t, "Result: 3\n", output.String())
}

type badOutput struct {
}

func (b badOutput) Write(p []byte) (n int, err error) {
	return 0, errors.New("Bad writer")
}
