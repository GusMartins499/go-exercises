package main

import "testing"

func TestInverterString(t *testing.T) {
	s := "Original"
	esperado := "lanigirO"
	invertida := InverterString(s)

	if invertida != esperado {
		t.Error("Teste de inverter string falhou")
	}
}
