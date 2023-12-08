package main

import "fmt"

type ContaCorrente struct {
	titular string
	numeroAgencia int
	numeroConta int 
	saldo float64

}

func (c *ContaCorrente) Sacar (valorDoSaque float64) string{
	podeSacar := valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	}else{
		return "Saldo insuficiente"
	}
}

func main() {

	contaDoGuilherme := ContaCorrente{titular: "Guilherme", numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}// se fizermos uma comparação de tipo so mudando o nome da variavl = true a não ser que mude algum valor

	
	contaDaLuisa := ContaCorrente{"Bruna", 222, 111222, 200} // o mesmo para esse tipo de declaração 
	
	
	fmt.Println(contaDoGuilherme)
	fmt.Println(contaDaLuisa)

	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"
	contaDaCris.saldo = 500

	// Aqui se fizer a comparação (contaDaCris == contaDaCris) da false pois ele vai comparar os endereçosonde esta alocada a variavel

	//porém se fizer (*contaDaCris == *contaDaCris) vai dar true pois está comparando o conteúdo

	fmt.Println(*contaDaCris)//conteudo de dentro sem o & (de endereço de memoria)
	fmt.Println(contaDaCris)//com o & comercial

}