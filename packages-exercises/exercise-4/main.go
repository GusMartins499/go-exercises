/* Faça um programa em Go utilizando o pacote `sort` para ordenar
   []int{5, 2, 8, 1, 9} e []string{"pera", "banana", "uva", "abacaxi"},
   imprimindo o seguinte resultado:

   [1 2 5 8 9]
   [abacaxi banana pera uva] */

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{5, 2, 8, 1, 9}
	texts := []string{"pera", "banana", "uva", "abacaxi"}
	
	sort.Ints(nums)
	sort.Strings(texts)

	//slices.Sort(nums)
	//slices.Sort(texts)

	fmt.Println(nums)
	fmt.Println(texts)
}