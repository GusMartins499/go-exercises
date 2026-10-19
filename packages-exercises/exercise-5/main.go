/* Faça um programa em Go criando um pacote próprio chamado `calculadora`
   (em uma pasta separada) com as funções exportadas Somar e Subtrair, e
   use esse pacote no main para imprimir o seguinte resultado:

   10 + 5 = 15
   10 - 5 = 5 */

package main

import (
	"fmt"
	"package-exercises/exercise-5/calculadora"
)

func main() {
	fmt.Printf("10 + 5 = %d\n", calculadora.Somar(10, 5))
	fmt.Printf("10 - 5 = %d\n", calculadora.Subtrair(10, 5))
}
