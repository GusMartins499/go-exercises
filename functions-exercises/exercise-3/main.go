/* Faça um programa em Go utilizando uma função recursiva fatorial para
   imprimir o seguinte resultado:

   5! = 120
   0! = 1 */

package main

import "fmt"

func factorial(num int) int {
	if num <= 1 {
		return 1
	}
	return num * factorial(num-1)
}

func main() {
	fmt.Println(factorial(5))
}
