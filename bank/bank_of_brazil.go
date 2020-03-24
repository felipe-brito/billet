package bank

import (
	"github.com/felipe-brito/billet/entity"
	bank "github.com/felipe-brito/billet/interfaces"
	"github.com/felipe-brito/billet/model"
)

var (
	configurationBankOfBrazil = bank.BankConfiguration{
		Id:             "001",
		Accept:         false,
		Currency:       9,
		CurrencyName:   "R$",
		AgencyMaxSize:  4,
		AccountMaxSize: 8,
	}
)

type BenchOfBrazil struct {
}

func (b *BenchOfBrazil) WithCompany(company entity.Company) bank.Bank {
	return nil
}

func (b *BenchOfBrazil) WithDocument(document entity.Document) bank.Bank {
	return nil
}

func (b *BenchOfBrazil) WithInstructions(instructions entity.Instructions) bank.Bank {
	return nil
}

func (b *BenchOfBrazil) WithPayer(payer entity.Payer) bank.Bank {
	return nil
}

func (b *BenchOfBrazil) Build() (model.Billet, error) {
	return model.Billet{}, nil
}
