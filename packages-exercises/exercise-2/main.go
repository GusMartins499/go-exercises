/* Faça um programa em Go utilizando o pacote `math` para imprimir o
   seguinte resultado:

   Raiz de 144: 12
   2 elevado a 10: 1024
   Maior entre 7 e 12: 12
   Pi: 3.14 */

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Raiz de 144:", math.Sqrt(144))
   fmt.Println("2 elevado a 10:", math.Pow(2, 10))
   fmt.Println("Maior entre 7 e 12:", math.Max(7, 12))
   fmt.Printf("Pi: %.2f\n", math.Pi)
}