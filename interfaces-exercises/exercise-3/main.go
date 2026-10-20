/* Faça um programa em Go utilizando a interface fmt.Stringer (método
   String() string) em uma struct Produto (Nome, Preco), de forma que
   fmt.Println(produto) imprima o seguinte resultado:

   Notebook - R$ 3500.00 */

package main

import "fmt"

type Produto struct {
	Nome  string
	Preco float64
}

func (p Produto) String() string {
	return fmt.Sprintf("%s R$ %.2f", p.Nome, p.Preco)
}

func main() {
	p := Produto{
		Nome:  "Macbook",
		Preco: 13599.90,
	}

	fmt.Println(p)
}
