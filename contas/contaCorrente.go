package contas

import (
	"github.com/matheusppaiva/go-bank/clients"
)

type ContaCorrente struct {
	Titular       clients.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque int) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= int(c.saldo)
	if podeSacar {
		c.saldo -= float64(valorDoSaque)
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}

}

func (c *ContaCorrente) Depositar(ValorDoDeposito float64) (string, float64) {
	podeDepositar := ValorDoDeposito > 0
	if podeDepositar {
		c.saldo += float64(ValorDoDeposito)
		return "Valor depositado", c.saldo
	} else {
		return "error : Deposito menor que 0", c.saldo
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferia float64, contaDestino *ContaCorrente) bool {

	if valorDaTransferia < c.saldo && valorDaTransferia > 0 {
		c.saldo -= valorDaTransferia
		contaDestino.Depositar(valorDaTransferia)
		return true
	} else {
		return false
	}

}

func (c *ContaCorrente) Saldo() float64 {
	return c.saldo
}
