/* Faça um programa em Go utilizando um map para contar quantas vezes
   cada palavra aparece na frase "o rato roeu a roupa do rei de roma o rato".
   Imprima as palavras em ordem alfabética (dica: pacote `sort`), gerando o
   seguinte resultado:

   a: 1
   de: 1
   do: 1
   o: 2
   rato: 2
   rei: 1
   roeu: 1
   roma: 1
   roupa: 1 */

package main

import (
	"fmt"
	"maps"
	"slices"
	"strings"
)

func main() {
	frase := "o rato roeu a roupa do rei de roma o rato"
	fraseSlice := strings.Fields(frase)
	dic := make(map[string]int)
	
	for _, palavra := range fraseSlice {
		dic[palavra]++
	}

	for _, palavra := range slices.Sorted(maps.Keys(dic)) {
		fmt.Printf("%s: %d\n", palavra, dic[palavra])
	}
}