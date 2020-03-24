package factory

import (
	"errors"
	"github.com/felipe-brito/billet/bank"
	"github.com/felipe-brito/billet/interfaces"
)

var banks = [...]BankList{
	{Bank: bank.NewBradesco(), Bench: interfaces.Bradesco},
	{Bank: nil, Bench: interfaces.Santander},
	{Bank: nil, Bench: interfaces.BankOfBrazil},
	{Bank: nil, Bench: interfaces.Itau},
	{Bank: nil, Bench: interfaces.FederalEconomicBox},
}

// Define
type BankList struct {
	Bank  interfaces.Bank
	Bench interfaces.Bench
}

// Define
func GetBank(bench interfaces.Bench) (interfaces.Bank, error) {
	for _, value := range banks {
		if value.Bench == bench {
			return value.Bank, nil
		}
	}
	return nil, errors.New("no bank available")
}
