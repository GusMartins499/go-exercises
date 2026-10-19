/* Faça um programa em Go utilizando switch com fallthrough para listar as
   permissões acumuladas de um usuário de nível 2 (nível 3: relatórios,
   nível 2: edição, nível 1: leitura), imprimindo o seguinte resultado:

   Acesso a edição
   Acesso a leitura */

package main

func main() {
	verifyRole := func(role int) {
		switch role {
		case 3:
			println("Acesso a relatórios")
			fallthrough
		case 2:
			println("Acesso a edição")
			fallthrough
		case 1:
			println("Acesso a leitura")
		default:
			println("Acesso negado")
		}
	}

	verifyRole(1)
	verifyRole(2)
	verifyRole(3)
}
