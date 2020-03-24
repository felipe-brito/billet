package generator

import (
	"fmt"
	"github.com/felipe-brito/billet/utils"
	"strconv"
)

const (
	minMultiplierModule10 = 1
	maxMultiplierModule10 = 2
	minMultiplierModule11 = 2
	maxMultiplierModule11 = 9
	module11              = 11
	module10              = 10
)

// Define
type DigitGeneratorImp struct {
}

// Define
type DigitGenerator interface {
	Module11(string) int
	Module10(string) int
	ITypeCheckerInOurNumber(string) int
}

// Module11
func (d *DigitGeneratorImp) Module11(input string) int {
	digit := utils.NewDigit()

	result := digit.
		WithNumber(input).
		WithMultipliers(minMultiplierModule11, maxMultiplierModule11).
		WithComplement(true).
		WithModule(module11).
		WithOneToOne(false).
		WithReplacement("1", []int{0, 1, 10, 11}).
		Calc()

	aux, _ := strconv.Atoi(result)
	return aux
}

func (d *DigitGeneratorImp) Module10(input string) int {
	digit := utils.NewDigit()

	result := digit.
		WithNumber(input).
		WithMultipliersInOrder(maxMultiplierModule10, minMultiplierModule10).
		WithComplement(true).
		WithModule(module10).
		WithOneToOne(true).
		WithReplacement("0", []int{10}).
		Calc()

	aux, _ := strconv.Atoi(result)
	return aux
}

func (d *DigitGeneratorImp) ITypeCheckerInOurNumber(string) int {
	fmt.Println("Implement me!")
	return 1
}
