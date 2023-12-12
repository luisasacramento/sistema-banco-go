package contas

import "banco/clientes"

type ContaCorrente struct {
	Titular clientes.Titular
	TitularCPF string
	NumeroAgencia int
	NumeroConta int 
	saldo float64

}

func (c *ContaCorrente) Sacar (valorDoSaque float64) string{
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	}else{
		return "saldo insuficiente"
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

func (c*ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool{
	if valorDaTransferencia < c.saldo && valorDaTransferencia > 0{
		c.saldo -= valorDaTransferencia
		contaDestino.Dspositar(valorDaTransferencia)
		return true
		
	}else{
		return false
	}

}

func (c* ContaCorrente) ObterSaldo() float64{
	return c.saldo
}