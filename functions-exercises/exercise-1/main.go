/* Faça um programa em Go utilizando uma função com dois retornos
   dividir(a, b int) (int, int), que retorna o quociente e o resto,
   para imprimir o seguinte resultado:

   10 / 3 = 3, resto 1 */

package main

import "fmt"

func main() {
	dividir := func(a, b int) (int, int) {
		return a / b, a % b
	}
	num1 := 10
	num2 := 3
	quociente, resto := dividir(10, 3)
	fmt.Printf("%d / %d = %d, resto %d\n", num1, num2, quociente, resto)
}
