package builder

import (
	"errors"
	"github.com/felipe-brito/billet/entity"
	"github.com/felipe-brito/billet/generator"
	"github.com/felipe-brito/billet/utils"
	"strconv"
)

const (
	barcodeNumberMaxSize = 44
)

//	Define
//
//	@currency
//	@document
//	@date
//	@bankId
//	@adjunct
//	@barcode
//	@digitGenerator
type BarcodeBuilderImpl struct {
	currency       string
	document       entity.Document
	date           entity.Dates
	bankId         string
	adjunct        string
	barcode        string
	digitGenerator generator.DigitGenerator
}

// Define
type BarcodeBuilder interface {
	WithDigitGenerator(generator.DigitGenerator) BarcodeBuilder
	WithSpecialCurrency(int) BarcodeBuilder
	WithBankId(string) BarcodeBuilder
	WithDocument(entity.Document) BarcodeBuilder
	WithDate(entity.Dates) BarcodeBuilder
	WithAdjunct(string) BarcodeBuilder
	Barcode() (string, error)
	Digitable() string
}

// NewBarcodeBuilder
func NewBarcodeBuilder() BarcodeBuilder {
	return &BarcodeBuilderImpl{}
}

//WithDigitGenerator
// @digitGenerator
func (b *BarcodeBuilderImpl) WithDigitGenerator(digitGenerator generator.DigitGenerator) BarcodeBuilder {
	b.digitGenerator = digitGenerator
	return b
}

// WithSpecialCurrency
// @currency
func (b *BarcodeBuilderImpl) WithSpecialCurrency(currency int) BarcodeBuilder {
	b.currency = strconv.Itoa(currency)
	return b
}

// WithBankId
// @bankId
func (b *BarcodeBuilderImpl) WithBankId(bankId string) BarcodeBuilder {
	b.bankId = bankId
	return b
}

// WithDocument
// @document
func (b *BarcodeBuilderImpl) WithDocument(document entity.Document) BarcodeBuilder {
	b.document = document
	return b
}

//	WithDate method for add date
//	@date
func (b *BarcodeBuilderImpl) WithDate(date entity.Dates) BarcodeBuilder {
	b.date = date
	return b
}

// WithAdjunct
// @adjunct
func (b *BarcodeBuilderImpl) WithAdjunct(adjunct string) BarcodeBuilder {
	b.adjunct = adjunct
	return b
}

// Generate
func (b *BarcodeBuilderImpl) Barcode() (string, error) {

	builder := utils.NewStringBuilder()

	dateDueFactor, err := b.date.DateDueFactor()

	if err != nil {
		return "", err
	}

	b.barcode = builder.
		Append(b.bankId).
		Append(b.currency).
		AppendInt(dateDueFactor).
		Append(b.document.GetFormattedValue()).
		Append(b.adjunct).
		Builder()

	digit := b.digitGenerator.Module11(b.barcode)

	b.barcode = builder.InsertInt(4, digit).Builder()

	if err = b.validateSizeOfBarcode(); err != nil {
		return "", err
	}

	return b.barcode, nil
}

// Digitable
func (b *BarcodeBuilderImpl) Digitable() string {

	stringBuilder := utils.NewStringBuilder()

	str := stringBuilder.
		Append(b.barcode[0:3]).
		Append(b.barcode[3:4]).
		Append(b.barcode[19:24]).Builder()

	TypeCheckerBlock1 := strconv.Itoa(b.digitGenerator.Module10(str))

	str = stringBuilder.
		Append(TypeCheckerBlock1).
		Append(b.barcode[24:34]).Builder()

	TypeCheckerBlock2 := strconv.Itoa(b.digitGenerator.Module10(str[10:20]))

	str = stringBuilder.
		Append(TypeCheckerBlock2).
		Append(b.barcode[34:]).Builder()

	TypeCheckerBlock3 := strconv.Itoa(b.digitGenerator.Module10(str[21:31]))

	str = stringBuilder.
		Append(TypeCheckerBlock3).
		Append(b.barcode[4:5]).
		Append(b.barcode[5:9]).
		Append(b.barcode[9:19]).
		Builder()

	return b.format(str)
}

func (b *BarcodeBuilderImpl) format(input string) string {
	st := utils.NewStringBuilder()

	st.
		Append(input).
		Insert(5, ".").
		Insert(11, "  ").
		Insert(18, ".").
		Insert(25, "  ").
		Insert(32, ".").
		Insert(39, "  ").
		Insert(42, "  ")

	return st.Builder()
}

// validateSizeOfBarcode
func (b *BarcodeBuilderImpl) validateSizeOfBarcode() error {
	if len(b.barcode) != barcodeNumberMaxSize {
		return errors.New("error in barcode generation. Number of digits is different from 44")
	}
	return nil
}
