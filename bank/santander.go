package bank

import (
	"github.com/felipe-brito/billet/entity"
	bank "github.com/felipe-brito/billet/interfaces"
	"github.com/felipe-brito/billet/model"
)

var (
	configurationSantander = bank.BankConfiguration{
		Id:             "033",
		Accept:         false,
		Currency:       9,
		CurrencyName:   "R$",
		AgencyMaxSize:  4,
		AccountMaxSize: 7,
	}
)

type Santander struct {
}

func (s *Santander) WithBeneficiary(beneficiary entity.Beneficiary) bank.Bank {
	return nil
}

func (s *Santander) WithCompany(company entity.Company) bank.Bank {
	return nil
}

func (s *Santander) WithDocument(document entity.Document) bank.Bank {
	return nil
}

func (s *Santander) WithInstructions(instructions entity.Instructions) bank.Bank {
	return nil
}

func (s *Santander) WithPayer(payer entity.Payer) bank.Bank {
	return nil
}

func (s *Santander) Build() (model.Billet, error) {
	return model.Billet{}, nil
}
