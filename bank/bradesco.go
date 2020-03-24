package bank

import (
	"github.com/felipe-brito/billet/builder"
	"github.com/felipe-brito/billet/entity"
	"github.com/felipe-brito/billet/generator"
	"github.com/felipe-brito/billet/interfaces"
	"github.com/felipe-brito/billet/model"
	"github.com/felipe-brito/billet/utils"
)

const (
	ourNumberMaxSize = 11
	agreementMaxSize = 7
)

var (
	configurationBradesco = interfaces.BankConfiguration{
		Id:             "237",
		Accept:         false,
		Currency:       9,
		CurrencyName:   "R$",
		AgencyMaxSize:  4,
		AccountMaxSize: 7,
	}
)

// Define
type Bradesco struct {
	beneficiary  entity.Beneficiary
	document     entity.Document
	company      entity.Company
	instructions entity.Instructions
	payer        entity.Payer
	date         entity.Dates
}

// NewBradesco
func NewBradesco() interfaces.Bank {
	return &Bradesco{}
}

// GetDigitGenerator
func (b *Bradesco) GetDigitGenerator() generator.DigitGenerator {
	return &generator.DigitGeneratorImp{}
}

// GetFormattedAgency
func (b *Bradesco) GetFormattedAgency() string {
	return utils.LeftPadWithZeros(b.beneficiary.Agency, configurationBradesco.AgencyMaxSize)
}

// GetFormattedWallet
func (b *Bradesco) GetFormattedWallet() string {
	return utils.LeftPadWithZeros(b.beneficiary.Wallet, 2)
}

//GetFormattedAccount
func (b *Bradesco) GetFormattedAccount() string {
	return utils.LeftPadWithZeros(b.beneficiary.Account, configurationBradesco.AccountMaxSize)
}

// GetOurFormattedNumber
func (b *Bradesco) GetOurFormattedNumber() string {
	return utils.LeftPadWithZeros(b.document.OurNumber, ourNumberMaxSize)
}

// GetFormattedAgreement
func (b *Bradesco) GetFormattedAgreement() string {
	return utils.LeftPadWithZeros(b.beneficiary.Agreement, agreementMaxSize)
}

// GetOurNumberWithCheckDigit
func (b *Bradesco) GetOurNumberWithCheckDigit() string {
	return "Implement me!"
}

// WithBeneficiary
func (b *Bradesco) WithBeneficiary(beneficiary entity.Beneficiary) interfaces.Bank {
	b.beneficiary = beneficiary
	return b
}

// WithCompany
func (b *Bradesco) WithCompany(company entity.Company) interfaces.Bank {
	b.company = company
	return b
}

// WithDocument
func (b *Bradesco) WithDocument(document entity.Document) interfaces.Bank {
	b.document = document
	return b
}

// WithInstructions
func (b *Bradesco) WithInstructions(instructions entity.Instructions) interfaces.Bank {
	b.instructions = instructions
	return b
}

//	WithDate
//	@date
func (b *Bradesco) WithDate(date entity.Dates) interfaces.Bank {
	b.date = date
	return b
}

// WithPayer
func (b *Bradesco) WithPayer(payer entity.Payer) interfaces.Bank {
	b.payer = payer
	return b
}

// Define
func (b *Bradesco) Build() (model.Billet, error) {

	stringBuilder := utils.NewStringBuilder()

	str := stringBuilder.
		Append(b.GetFormattedAgency()).
		Append(b.GetFormattedWallet()).
		Append(b.GetOurFormattedNumber()).
		Append(b.GetFormattedAccount()).
		Append("0").
		Builder()

	barcodeBuilder := builder.NewBarcodeBuilder()

	barcode, err := barcodeBuilder.
		WithDigitGenerator(b.GetDigitGenerator()).
		WithBankId(configurationBradesco.Id).
		WithSpecialCurrency(configurationBradesco.Currency).
		WithDocument(b.document).
		WithDate(b.date).
		WithAdjunct(str).
		Barcode()

	if err != nil {
		return model.Billet{}, err
	}

	return model.Billet{
		Barcode:   barcode,
		Digitable: barcodeBuilder.Digitable(),
	}, nil
}
