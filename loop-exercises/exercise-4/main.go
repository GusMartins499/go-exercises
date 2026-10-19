/* Faça um programa em Go utilizando for com range sobre o slice
   []string{"maçã", "banana", "uva"}, imprimindo o seguinte resultado:

   0: maçã
   1: banana
   2: uva */

package main

import "fmt"

func main() {
	fruits := []string{"maçã", "banana", "uva"}
	for index, f := range fruits {
		fmt.Printf("%d: %s\n", index, f)
	}
}
