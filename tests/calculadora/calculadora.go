package calculadora

import "errors"

var ErrDivisaoPorZero = errors.New("divisão por zero")

func Dividir(a, b int) (int, error) {
	if b <= 0 {
		return 0, ErrDivisaoPorZero
	}

	return a / b, nil
}
