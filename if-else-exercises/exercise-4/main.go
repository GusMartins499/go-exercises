/* Faça um programa em Go utilizando if/else com && e || para verificar se
   os anos 2024, 1900, 2000 e 2023 são bissextos (divisível por 4 e não por
   100, ou divisível por 400), imprimindo o seguinte resultado:

   2024 é bissexto
   1900 não é bissexto
   2000 é bissexto
   2023 não é bissexto */

package main

import "fmt"

func verify(year uint) {
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		fmt.Println(year, "é bissexto")
	} else {
		fmt.Println(year, "não é bissexto")
	}
}

func main() {
	verify(2024)
	verify(1900)
	verify(2000)
	verify(2023)
}
