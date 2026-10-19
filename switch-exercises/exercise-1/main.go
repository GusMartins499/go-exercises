/* Faça um programa em Go utilizando switch para converter um número no dia
   da semana (1 = Domingo ... 7 = Sábado). Teste os números 1, 3 e 9,
   imprimindo o seguinte resultado:

   1: Domingo
   3: Terça-feira
   9: Dia inválido */

package main

import "fmt"

func verifyDay(day int) {
	switch {
	case day == 1:
		fmt.Println("1: Domingo")
	case day == 2:
		fmt.Println("2: Segunda-feira")
	case day == 3:
		fmt.Println("3: Terça-feira")
	case day == 4:
		fmt.Println("4: Quarta-feira")
	case day == 5:
		fmt.Println("5: Quinta-feira")
	case day == 6:
		fmt.Println("6: Sexta-feira")
	case day == 7:
		fmt.Println("7: Sábado")
	default:
		fmt.Println("Dia inválido")
	}
}

func main() {
	verifyDay(1)
	verifyDay(3)
	verifyDay(9)
}
