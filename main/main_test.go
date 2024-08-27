package main

import (
	"bytes"
	"errors"
	"testing"

	"github.com/smartystreets/calcy-lib/calc"
)

func TestWrongNumArgs(t *testing.T) {
	output := bytes.Buffer{}
	handler := NewHandler(calc.Addition{}, &output)
	err := handler.Handle([]string{"1"})
	if !errors.Is(err, errWrongNumArgs) {
		t.Errorf("expected %v, got %v", errWrongNumArgs, err)
	}
	if output.Len() > 0 {
		t.Error("expected empty output")
	}
}

func TestInvalidFirstArg(t *testing.T) {
	output := bytes.Buffer{}
	handler := NewHandler(calc.Addition{}, &output)
	err := handler.Handle([]string{"a", "1"})
	if !errors.Is(err, errMalformedArgument) {
		t.Errorf("expected %v, got %v", errOutputWriteErr, err)
	}
	if output.Len() > 0 {
		t.Error("expected empty output")
	}
}

func TestInvalidSecondArg(t *testing.T) {
	output := bytes.Buffer{}
	handler := NewHandler(calc.Addition{}, &output)
	err := handler.Handle([]string{"1", "a"})
	if !errors.Is(err, errMalformedArgument) {
		t.Errorf("expected %v, got %v", errOutputWriteErr, err)
	}
	if output.Len() > 0 {
		t.Error("expected empty output")
	}
}

func TestInvalidWrite(t *testing.T) {
	output := &badOutput{}
	handler := NewHandler(calc.Addition{}, output)
	err := handler.Handle([]string{"1", "1"})
	if !errors.Is(err, errOutputWriteErr) {
		t.Errorf("expected %v, got %v", errOutputWriteErr, err)
	}
}

func TestHappyPath(t *testing.T) {
	output := bytes.Buffer{}
	handler := NewHandler(calc.Addition{}, &output)
	err := handler.Handle([]string{"1", "2"})
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

type badOutput struct {
}

func (b badOutput) Write(p []byte) (n int, err error) {
	return 0, errors.New("Bad writer")
}
