/* Faça um programa em Go utilizando uma closure: uma função novoContador()
   que retorna outra função que, a cada chamada, retorna o próximo número.
   Chame um contador 3 vezes e depois crie um segundo contador e chame-o
   1 vez, imprimindo o seguinte resultado:

   1
   2
   3
   1 */

package main

import "fmt"

func novoContador() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	contador1 := novoContador()
	fmt.Println(contador1())
	fmt.Println(contador1())
	fmt.Println(contador1())
	contador2 := novoContador()
	fmt.Println(contador2())
}
