/* Faça um programa em Go com a função Somar(a, b int) int e um arquivo
   main_test.go com a função TestSomar(t *testing.T), que usa t.Errorf
   caso Somar(2, 3) não retorne 5. Ao rodar `go test -v`, deve aparecer
   o seguinte resultado:

   === RUN   TestSomar
   --- PASS: TestSomar (0.00s)
   PASS
   ok  	tests	0.002s

   Depois, troque o valor esperado para 6 e rode `go test` de novo, para
   ver como uma falha aparece:

   --- FAIL: TestSomar (0.00s)
       main_test.go:11: Somar(2, 3) = 5; esperado 6
   FAIL
   exit status 1
   FAIL	tests	0.003s

   Volte o valor esperado para 5 antes de seguir. */

package main

import "testing"

func TestSomar(t *testing.T) {
	esperado := 5
	if total := Somar(2, 3); total != esperado {
		t.Errorf("Somar(2, 3) = %d; esperado %d", total, esperado)
	}
}
