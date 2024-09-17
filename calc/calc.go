package calc

type Addition struct {
}
type Subtraction struct{}
type Division struct{}
type Multiplication struct{}

func (this Addition) Calculate(a, b int) int {
	return a + b
}

func (this Subtraction) Calculate(a, b int) int {
	return a - b
}

func (this Division) Calculate(a, b int) int {
	return a / b
}

func (this Multiplication) Calculate(a, b int) int {
	return a * b
}
