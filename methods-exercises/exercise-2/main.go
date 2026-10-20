/* Faça um programa em Go utilizando uma struct ContaBancaria com os métodos
   Depositar(valor) e Sacar(valor) error, ambos com receptor ponteiro. Sacar
   retorna um erro se não houver saldo. Comece com saldo 100, deposite 50,
   saque 30 e tente sacar 500, imprimindo o seguinte resultado:

   Saldo: 150.00
   Saldo: 120.00
   Erro: saldo insuficiente
   Saldo final: 120.00 */

package main

import (
	"errors"
	"fmt"
)

type ContaBancaria struct {
	amount float64
}

func (c *ContaBancaria) Sacar(value float64) error {
	if c.amount < value {
		return errors.New("saldo insuficiente")
	}
	c.amount -= value
	return nil
}

func (c *ContaBancaria) Depositar(value float64) {
	c.amount += value
}

func main() {
	account := ContaBancaria{
		amount: 100.0,
	}

	account.Depositar(50)
	fmt.Printf("Saldo: %.2f\n", account.amount)

	if err := account.Sacar(30); err != nil {
		fmt.Println("Erro:", err)
	}
	fmt.Printf("Saldo: %.2f\n", account.amount)

	if err := account.Sacar(500); err != nil {
		fmt.Println("Erro:", err)
	}
	fmt.Printf("Saldo final: %.2f\n", account.amount)
}