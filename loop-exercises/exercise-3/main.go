/* Faça um programa em Go utilizando for para imprimir os números de 1 a 15,
   trocando os múltiplos de 3 por "Fizz", os de 5 por "Buzz" e os de ambos
   por "FizzBuzz":

   1
   2
   Fizz
   4
   Buzz
   Fizz
   7
   8
   Fizz
   Buzz
   11
   Fizz
   13
   14
   FizzBuzz */

package main

import "fmt"

func main() {
	for i := 1; i <= 15; i++ {
		s := ""
		if i%3 == 0 {
			s += "Fizz"
		}
		if i%5 == 0 {
			s += "Buzz"
		}
		if s == "" {
			s = fmt.Sprint(i)
		}
		fmt.Println(s)
	}
}
