/* Faça um programa em Go utilizando if/else para simular um login com o
   usuário "admin" e a senha "1234". Teste as combinações
   (admin, 1234), (admin, 0000) e (joao, 1234), imprimindo o seguinte
   resultado:

   admin/1234: Login realizado com sucesso
   admin/0000: Senha incorreta
   joao/1234: Usuário não encontrado */

package main

import "fmt"

func login(user string, password string) {
	if user == "admin" && password == "1234" {
		fmt.Println("Login realizado com sucesso")
	} else if user == "admin" && password != "1234" {
		fmt.Println("Senha incorreta")
	} else {
		fmt.Println("Usuário não encontrado")
	}
}

func main() {
	login("admin", "1234")
	login("admin", "0000")
	login("joao", "1234")
}
