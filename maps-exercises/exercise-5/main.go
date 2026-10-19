/* Faça um programa em Go utilizando um map[string]float64 como carrinho de
   compras com arroz: 25.90, feijão: 8.50 e café: 18.75. Some os valores e
   imprima o seguinte resultado:

   Itens no carrinho: 3
   Total: R$ 53.15 */

package main

import "fmt"

func main() {
	carrinho := map[string]float64{
		"arroz":  25.90,
		"feijão": 8.50,
		"café":   18.75,
	}
	var total float64
	for _, valor := range carrinho {
		total += valor
	}
	fmt.Printf("Itens no carrinho: %d\n", len(carrinho))
	fmt.Printf("Total: R$ %.2f\n", total)
}
