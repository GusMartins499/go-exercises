/* Faça um programa em Go utilizando a interface error com um tipo próprio,
   type ErroIdade struct{ Idade int }, e uma função validarIdade(idade int)
   error. Teste as idades 30 e -5, imprimindo o seguinte resultado:

   Idade 30 válida
   Erro: idade inválida: -5 */

package main

import "fmt"

type ErroIdade struct {
	Idade int
}

func (e *ErroIdade) Error() string {
	return fmt.Sprintf("idade inválida: %d", e.Idade)
}

func validarIdade(idade int) error {
	if idade < 0 {
		return &ErroIdade{Idade: idade}
	}
	return nil
}

func main() {
	for _, idade := range []int{30, -5} {
		if err := validarIdade(idade); err != nil {
			fmt.Println("Erro:", err)
		} else {
			fmt.Printf("Idade %d válida\n", idade)
		}
	}
}
