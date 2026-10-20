/* Faça um programa em Go utilizando uma struct Contador com o método
   IncrementarValor() (receptor por valor) e IncrementarPonteiro() (receptor
   ponteiro), mostrando a diferença entre os dois, para imprimir o seguinte
   resultado:

   Após IncrementarValor: 0
   Após IncrementarPonteiro: 1 */

package main

import "fmt"

type Contador struct {
	valor int
}

func (c *Contador) IncrementarPonteiro() {
	c.valor++
}

func (c Contador) IncrementarValor() {
	c.valor++
}

func main() {
	count := Contador{
		valor: 0,
	}

	count.IncrementarValor()
	fmt.Println("Após IncrementarValor", count.valor)
	count.IncrementarPonteiro()
	fmt.Println("Após IncrementarPonteiro", count.valor)
}
