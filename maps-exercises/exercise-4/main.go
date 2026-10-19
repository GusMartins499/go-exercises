/* Faça um programa em Go utilizando um map[string][]string com as
   categorias "frutas" (maçã, uva, banana) e "legumes" (cenoura, batata).
   Imprima as categorias em ordem alfabética, gerando o seguinte resultado:

   frutas: [banana maçã uva]
   legumes: [batata cenoura] */

package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	categorias := map[string][]string{
		"legumes": {"cenoura", "batata"},
		"frutas":  {"maçã", "uva", "banana"},
	}

	for _, categoria := range slices.Sorted(maps.Keys(categorias)) {
		itens := categorias[categoria]
		slices.Sort(itens)
		fmt.Printf("%s: %v\n", categoria, itens)
	}
}
