/* Faça um programa em Go com a função Classificar(nota float64) string
   (a mesma regra do exercício 2 de IF / ELSE) e um teste em tabela
   (table-driven): um slice de structs com os campos nome, nota e esperado,
   percorrido com for e t.Run para criar um subteste por caso. Use os
   casos "nota A" (9.5), "nota B" (7.0), "nota C" (5.5) e "reprovado" (3.0).
   Ao rodar `go test -v -run TestClassificar`, deve aparecer o seguinte
   resultado:

   === RUN   TestClassificar
   === RUN   TestClassificar/nota_A
   === RUN   TestClassificar/nota_B
   === RUN   TestClassificar/nota_C
   === RUN   TestClassificar/reprovado
   --- PASS: TestClassificar (0.00s)
       --- PASS: TestClassificar/nota_A (0.00s)
       --- PASS: TestClassificar/nota_B (0.00s)
       --- PASS: TestClassificar/nota_C (0.00s)
       --- PASS: TestClassificar/reprovado (0.00s)
   PASS
   ok  	tests	0.002s

   Repare que o Go troca os espaços do nome do subteste por "_". */

package main

import "testing"

type Notas struct {
	nome     string
	nota     float32
	esperado string
}

func TestClassificar(t *testing.T) {
	notas := []Notas{
		{nome: "Nota >= 7 deve ser B", nota: 7, esperado: "B"},
		{nome: "Nota >= 9 deve ser A", nota: 9, esperado: "A"},
		{nome: "Nota >= 5 deve ser C", nota: 5, esperado: "C"},
		{nome: "Nota < 5 deve ser Reprovado", nota: 4, esperado: "Reprovado"},
		{nome: "Nota >= 7 deve ser B", nota: 7.5, esperado: "B"},
		{nome: "Nota >= 9 deve ser A", nota: 10, esperado: "A"},
		{nome: "Nota < 5 deve ser Reprovado", nota: 4.5, esperado: "Reprovado"},
		{nome: "Nota >= 9 deve ser A", nota: 9.5, esperado: "A"},
	}

	for _, nota := range notas {
		t.Run("Classificar nota", func(t *testing.T) {
			if got := ClassificarNota(float64(nota.nota)); got != nota.esperado {
				t.Errorf(
					"Erro na classificação da nota: %v, esperado: %s, recebido: %s. Regra: %s",
					nota.nota, nota.esperado, got, nota.nome,
				)
			}
		})
	}
}
