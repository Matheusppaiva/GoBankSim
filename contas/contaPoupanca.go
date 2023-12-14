package contas

import "github.com/matheusppaiva/go-bank/clients"

type ContaPoupanca struct {
	Titular       clients.Titular
	NumeroAgencia int
	NumeroConta   int
	Operacao      int
	saldo         float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= float64(valorDoSaque)
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaPoupanca) Depositar(ValorDoDeposito float64) (string, float64) {
	podeDepositar := ValorDoDeposito > 0
	if podeDepositar {
		c.saldo += float64(ValorDoDeposito)
		return "Valor depositado", c.saldo
	} else {
		return "Error: Depósito menor que 0", c.saldo
	}
}

func (c *ContaPoupanca) Saldo() float64 {
	return c.saldo
}
