/* Faça um programa em Go utilizando uma função variádica
   somar(numeros ...int) int para imprimir o seguinte resultado:

   somar(1, 2, 3, 4, 5) = 15
   somar() = 0 */

package main

import "fmt"

func main() {
	somar := func(nums ...int) int {
		total := 0
		for _, num := range nums {
			total += num
		}
		return total
	}
	fmt.Println(somar(1, 2, 3, 4, 5))
}
