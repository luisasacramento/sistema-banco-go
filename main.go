package main

import "fmt"

type ContaCorrente struct {
	titular string
	numeroAgencia int
	numeroConta int 
	saldo float64

}

func (c *ContaCorrente) Sacar (valorDoSaque float64) string{
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	}else{
		return "Saldo insuficiente"
	}
}

func (c*ContaCorrente) Dspositar(valorDoDeposito float64) (string, float64){
	if valorDoDeposito>0 {
		c.saldo+= valorDoDeposito
		return "Deposito realizado com sucesso", c.saldo
	}else{
		return "O valor do deposito menor do que 0", c.saldo
	}

}

func (c*ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino ContaCorrente) bool{
	if valorDaTransferencia < c.saldo {
		c.saldo -= valorDaTransferencia
		contaDestino.Dspositar(valorDaTransferencia)
		return true
		
	}else{
		return false
	}

}

func main() {

	contaDaLuisa := ContaCorrente{titular: "Luisa", saldo: 300}
	contaDoGuilherme := ContaCorrente{titular: "Guilherme", saldo: 100}

	status := contaDaLuisa.Transferir(200, contaDoGuilherme)

	fmt.Println(status)
	fmt.Println(contaDaLuisa)
	fmt.Println(contaDoGuilherme)

	

}