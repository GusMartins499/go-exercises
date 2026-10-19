/* Faça um programa em Go utilizando o pacote `strconv` para converter o
   texto "42" em número, somar 8 e converter o resultado de volta para
   texto, imprimindo o seguinte resultado:

	 Número: 50
   Texto "50" tem 2 caracteres
*/

package main

import (
	"fmt"
	"strconv"
)


func main() {
	const textNumber = "42"
	var number, _ = strconv.Atoi(textNumber)
	fmt.Printf("Número: %d\n", number + 8)
	fmt.Printf("Texto '50' tem %d caracteres\n", len(strconv.Itoa(number + 8)))
}