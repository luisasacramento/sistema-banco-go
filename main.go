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
        Profissao: "Desenvolvedor"},
        NumeroAgencia:123, NumeroConta: 123456}

	contaDoBruno.Dspositar(100)

    fmt.Println(contaDoBruno.ObterSaldo())
}
	
	
