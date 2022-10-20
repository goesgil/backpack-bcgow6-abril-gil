package calculadora

import "fmt"

var (
	errorDivisor = fmt.Errorf("error, divisor must be greater than zero")
)

func Dividir(a, b int) (int, error) {
	if b < 1 {
		return 0, errorDivisor
	}
	return a / b, nil
}
