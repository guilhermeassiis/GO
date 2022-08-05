package main

import "fmt"

// Objeto e atributos
type Endereco struct {
	logradouro string // Atributos privados
	num        int
	bairro     string
}
type Pessoa struct {
	name     string
	cpf      string
	endereco Endereco
}
type Conta struct {
	titular         Pessoa
	numeroDaConta   string
	numeroDaAgencia string
	saldo           float64
}

//Metodos
func (c *Conta) Sacar(valor float64) (float64, string) {
	valido := valor > 0 && valor < c.saldo

	if valido {
		c.saldo -= valor
		return c.saldo, "Saque efetuado com sucesso!"
	} else {
		return c.saldo, "Saldo insuficiente ou valor invalido!"
	}
}

func (c *Conta) Depositar(valor float64) (float64, string) {
	valido := valor > 0
	if valido {
		c.saldo += valor
		return c.saldo, "Deposito realizado com sucesso"
	} else {
		return c.saldo, "Deposito não realizado, checar o valor inserido"
	}

}

func (c *Conta) Transferir(valor float64, conta *Conta) string {
	valido := valor > 0 && valor < c.saldo
	if valido {
		c.saldo -= valor
		conta.saldo += valor
		return "Transferencia efetuada!"
	} else {
		return "Saldo insuficiente para a transferencia"
	}
}

func (p *Pessoa) ToString() {
	fmt.Println("Nome: ", p.name, " CPF:", p.cpf, " Endereço: ", p.endereco.logradouro, " N°", p.endereco.num, " Em ", p.endereco.bairro)
}

func main() {

	endereco := Endereco{"Rua Bolivia", 1, "Anjo da Guarda"}
	guilhermeObject := Pessoa{"Guilherme", "131.456.789-00", endereco}
	contaGuilherme := Conta{guilhermeObject, "12345", "456", 225.50}

	guilhermeObject.ToString()

	fmt.Println(contaGuilherme)
}
