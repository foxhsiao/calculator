package calculator

var logMessage = "[LOG]"

// Version of the calculator
var Version = 1.0

func internalSum(num int) int {
	return num - 1
}

// Sum two integer numbers
func Sum(number1, number2 int) int {
	return number1 + number2
}
