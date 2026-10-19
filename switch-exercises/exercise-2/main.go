/* Faça um programa em Go utilizando switch sem expressão (switch { ... })
   para classificar as temperaturas 35, 22 e 5 (>= 30: Muito quente,
   >= 18: Agradável, abaixo disso: Frio), imprimindo o seguinte resultado:

   35°C: Muito quente
   22°C: Agradável
   5°C: Frio */

package main

import "fmt"

func main() {
	classifyTemperature := func(temp int) {
		switch {
		case temp >= 30:
			fmt.Println(temp, "°C: Muito quente")
		case temp >= 18:
			fmt.Println(temp, "°C: Agradável")
		default:
			fmt.Println(temp, "°C: Frio")
		}
	}

	classifyTemperature(35)
	classifyTemperature(22)
	classifyTemperature(5)
}
