package calculator

func Add(a, b int) int {
	var addition = a + b
	return addition
}
func Sub(a, b int) int {
	subtract := a - b
	return subtract
}
func Div(a, b int) (int, int) {

	c, d := a/b, a%b
	return c, d
}
func Mul(a, b int) int {
	var multiplication = a * b
	return multiplication
}
