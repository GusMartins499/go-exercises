/* Faça um programa em Go utilizando um map[string]int com as idades
   Ana: 25, Bruno: 30 e Carla: 22. Imprima a idade do Bruno, adicione
   Diego: 28, remova a Carla e imprima o tamanho do map, gerando o
   seguinte resultado:

   Idade do Bruno: 30
   Total de pessoas: 3 */

package main

import "fmt"

func main() {
	idades := map[string]int{
		"Ana":   25,
		"Bruno": 30,
		"Carla": 22,
	}

	fmt.Println("Idade do Bruno:", idades["Bruno"])

	idades["Diego"] = 28
	delete(idades, "Carla")

	fmt.Println("Total de pessoas:", len(idades))
}
