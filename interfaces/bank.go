package interfaces

import (
	"github.com/felipe-brito/billet/entity"
	"github.com/felipe-brito/billet/generator"
	"github.com/felipe-brito/billet/model"
)

// Define
type Bench string

const (
	Santander          Bench = "033"
	Bradesco           Bench = "237"
	BankOfBrazil       Bench = "001"
	Itau               Bench = "341"
	FederalEconomicBox Bench = "104"
)

// Define
type BankConfiguration struct {
	Id             string
	Accept         bool
	Currency       int
	CurrencyName   string
	AgencyMaxSize  int
	AccountMaxSize int
}

// Define
type Bank interface {
	// auxiliary methods
	GetDigitGenerator() generator.DigitGenerator
	GetFormattedAgency() string
	GetFormattedWallet() string
	GetFormattedAccount() string
	GetOurFormattedNumber() string
	GetFormattedAgreement() string
	GetOurNumberWithCheckDigit() string

	// methods to build the barcode and the digitable line
	WithBeneficiary(beneficiary entity.Beneficiary) Bank
	WithCompany(company entity.Company) Bank
	WithDocument(document entity.Document) Bank
	WithInstructions(instructions entity.Instructions) Bank
	WithPayer(payer entity.Payer) Bank
	WithDate(date entity.Dates) Bank
	Build() (model.Billet, error)
}
