/* Faça um programa em Go utilizando um tipo que não é struct,
   type Celsius float64, com o método ParaFahrenheit() (F = C * 9 / 5 + 32),
   para converter 100 e 0, imprimindo o seguinte resultado:

   100.0°C = 212.0°F
   0.0°C = 32.0°F */

package main

import "fmt"

type Celsius float64

func (c Celsius) ParaFahrenheit() float64 {
	return float64(c)*9/5 + 32
}

func main() {
	temp := Celsius(100.0)
	fmt.Println(temp.ParaFahrenheit())
}
