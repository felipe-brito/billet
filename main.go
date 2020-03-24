package main

import (
	"fmt"
	"github.com/felipe-brito/billet/entity"
	"github.com/felipe-brito/billet/factory"
	"github.com/felipe-brito/billet/interfaces"
	"time"
)

func main() {
	bank, _ := factory.GetBank(interfaces.Bradesco)

	b := entity.Beneficiary{
		Agency:      "2374",
		Account:     "76492",
		TypeAccount: "2",
		Wallet:      "9",
		Agreement:   "4937949",
	}

	d := entity.Document{
		FebrabanType: "DM",
		Value:        100,
		OurNumber:    "2019797906", //inteiro quebra, colocar string
		//Date:         time.Date(2020, 03, 11, 0, 0, 0, 0, time.UTC),
	}

	da := entity.Dates{
		DateDue:     time.Date(2020, 03, 10, 0, 0, 0, 0, time.UTC),
		DateOfIssue: time.Date(2020, 03, 10, 0, 0, 0, 0, time.UTC),
	}

	bi, e := bank.
		WithBeneficiary(b).
		WithDocument(d).
		WithDate(da).
		Build()

	fmt.Println(e)
	fmt.Println(bi.Digitable)
	fmt.Println(bi.Barcode)
}
