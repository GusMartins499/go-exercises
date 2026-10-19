/* Faça um programa em Go utilizando o mesmo map do exercício anterior e
   a forma "valor, ok := mapa[chave]" para verificar se "Bruno" e "Zeca"
   existem, imprimindo o seguinte resultado:

   Bruno encontrado, idade 30
   Zeca não encontrado */

package main

import "fmt"

func main() {
	idades := map[string]int{
		"Ana":   25,
		"Bruno": 30,
		"Carla": 22,
	}

	if idade, ok := idades["Bruno"]; ok {
		fmt.Println("Bruno encontrado, idade", idade)
	} else {
		fmt.Println("Bruno não encontrado")
	}

	if idade, ok := idades["Zeca"]; ok {
		fmt.Println("Zeca encontrado, idade", idade)
	} else {
		fmt.Println("Zeca não encontrado")
	}
}
