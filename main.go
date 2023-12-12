package main

import (
	"fmt"
	"banco/contas"
	"banco/clientes"

)


func main() {
	contaDoBruno := contas.ContaCorrente{Titular: clientes.Titular{
		Nome: "Bruno",
		CPF: "123.111.123.12",
		Profissao: "Desenvovedor"},
		NumeroAgencia: 123, NumeroConta: 12345, Saldo: 100} 

	fmt.Println(contaDoBruno)
	}
	
