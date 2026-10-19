/* Faça um programa em Go utilizando if/else e o operador % para verificar
   os números 4 e 7, imprimindo o seguinte resultado:

   4 é par
   7 é ímpar */

package main

import "fmt"

func verifyOdd(num int) {
	if num%2 == 0 {
		fmt.Printf("%d é par\n", num)
	} else {
		fmt.Printf("%d é ímpar\n", num)
	}
}
func main() {
	verifyOdd(4)
	verifyOdd(7)
}
