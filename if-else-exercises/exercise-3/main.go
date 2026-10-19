/* Faça um programa em Go utilizando um if com declaração curta
   (if idade := 17; idade >= 18 { ... }) para imprimir o seguinte resultado:

   Acesso negado: faltam 1 ano(s) para completar 18 */

package main

import "fmt"

func main() {
	if idade := 17; idade >= 18 {
		fmt.Println("Acesso liberado")
	} else {
		fmt.Println("Acesso negado: faltam", 18-idade, "ano(s) para completar 18")
	}
}
