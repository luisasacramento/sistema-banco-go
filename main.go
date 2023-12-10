package main

import (
	"fmt"
	"banco/contas"

)


func main() {

	contaDaLuisa := contas.ContaCorrente{Titular: "Luisa", Saldo: 300}
	contaDoGuilherme := contas.ContaCorrente{Titular: "Guilherme", Saldo: 100}

	status := contaDaLuisa.Transferir(200, &contaDoGuilherme)

	fmt.Println(status)
	fmt.Println(contaDaLuisa)
	fmt.Println(contaDoGuilherme)

	

}