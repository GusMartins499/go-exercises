/* Faça um programa em Go utilizando o pacote `strings` para, a partir do
   texto "Go é divertido", imprimir o seguinte resultado:

   GO É DIVERTIDO
   go é divertido
   Palavras: 3
   Contém "divertido": true */

package main

import (
	"fmt"
	"strings"
)

func main() {
	const text = "Go é divertido"

	fmt.Println(strings.ToUpper(text))
	fmt.Println(strings.ToLower(text))
	fmt.Println("Palavras:", len(strings.Fields(text)))
	fmt.Println("Contém 'divertido'?", strings.Contains(text, "divertido"))
}