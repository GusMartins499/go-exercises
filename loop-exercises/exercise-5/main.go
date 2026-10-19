/* Faça um programa em Go utilizando dois for aninhados para imprimir o
   seguinte resultado:

   *
   **
   ***
   ****
   ***** */

package main

import (
	"fmt"
)

func l(base int) string {
	s := ""
	if base == 1 {
		return "*"
	}
	for range base {
		s = s + "*"
	}
	s = l(base-1) + "\n" + s
	return s
}

func main() {
	base := 5
	fmt.Println(l(base))
}
