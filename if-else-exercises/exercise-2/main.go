/* Faça um programa em Go utilizando if / else if / else para classificar
   as notas 9.5, 7.0, 5.5 e 3.0 (>= 9: A, >= 7: B, >= 5: C, abaixo disso:
   Reprovado), imprimindo o seguinte resultado:

   Nota 9.5: A
   Nota 7.0: B
   Nota 5.5: C
   Nota 3.0: Reprovado */

package main

import "fmt"

func classifyGrade(grade float64) string {
	if grade >= 9 {
		return "A"
	} else if grade >= 7 {
		return "B"
	} else if grade >= 5 {
		return "C"
	} else {
		return "Reprovado"
	}
}

func main () {
	classes := []float64{9.5, 7.0, 5.5, 3.0}
	for _, grade := range classes {
		result := classifyGrade(grade)
		fmt.Println("Nota", grade, ":", result)
	}
}