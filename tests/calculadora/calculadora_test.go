package calculadora

import (
	"errors"
	"testing"
)

func TestDividir(t *testing.T) {
	t.Parallel()
	t.Run("Divisão normal", func(t *testing.T) {
		resultado, err := Dividir(10, 2)
		if err != nil {
			t.Fatalf("Dividir(10, 2) retornou erro inesperado: %v", err)
		}
		if resultado != 5 {
			t.Errorf("Dividir(10, 2) = %d; esperado 5", resultado)
		}
	})
	t.Run("divisão por zero", func(t *testing.T) {
		t.Parallel()
		_, err := Dividir(10, 0)
		if !errors.Is(err, ErrDivisaoPorZero) {
			t.Errorf("Dividir(10, 0) erro = %v; esperado %v", err, ErrDivisaoPorZero)
		}
	})
}
