/* Faça um programa em Go utilizando for e continue para somar apenas os
   números pares de 1 a 100, imprimindo o seguinte resultado:

   Soma dos pares: 2550 */

package main

import "fmt"

func main() {
	total := 0
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			total += i
		}
	}

	fmt.Printf("Soma dos pares: %d\n", total)
}
