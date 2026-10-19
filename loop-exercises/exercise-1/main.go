/* Faça um programa em Go utilizando for para imprimir a tabuada do 7:

   7 x 1 = 7
   7 x 2 = 14
   7 x 3 = 21
   7 x 4 = 28
   7 x 5 = 35
   7 x 6 = 42
   7 x 7 = 49
   7 x 8 = 56
   7 x 9 = 63
   7 x 10 = 70 */

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("7 x %d = %d\n", i, 7*i)
	}
}
