package utils

import (
	"strconv"
)

type Digit struct {
	numbers           []int
	multipliers       map[int]int
	complement        bool
	module            int
	sumOneToOne       bool
	replacement       map[int]string
	multiplierInOrder bool
}

type DigitInterface interface {
	WithNumber(string) DigitInterface
	WithMultipliers(int, int) DigitInterface
	WithMultipliersInOrder(...int) DigitInterface
	WithComplement(bool) DigitInterface
	WithModule(int) DigitInterface
	WithOneToOne(bool) DigitInterface
	WithReplacement(string, []int) DigitInterface
	Calc() string
}

func NewDigit() DigitInterface {
	return &Digit{}
}

func (d *Digit) WithNumber(input string) DigitInterface {
	d.numbers = make([]int, len(input))
	sort := d.extractNumber(input)
	d.reverse(sort)
	return d
}

func (d *Digit) reverse(sort []int) {
	aux := 0
	for i := len(sort) - 1; i >= 0; i-- {
		d.numbers[aux] = sort[i]
		aux++
	}
}

func (d *Digit) extractNumber(input string) []int {
	sort := make([]int, len(input))
	for index, value := range input {
		n, _ := strconv.Atoi(string(value))
		sort[index] = n
	}
	return sort
}

func (d *Digit) WithMultipliers(begin int, end int) DigitInterface {
	d.multipliers = make(map[int]int)
	index := 0
	for i := begin; i <= end; i++ {
		d.multipliers[index] = i
		index++
	}
	return d
}
func (d *Digit) WithMultipliersInOrder(values ...int) DigitInterface {
	d.multipliers = make(map[int]int)
	index := 0
	for _, value := range values {
		d.multipliers[index] = value
		index++
	}
	return d
}
func (d *Digit) WithComplement(complement bool) DigitInterface {
	d.complement = complement
	return d
}

func (d *Digit) WithModule(module int) DigitInterface {
	d.module = module
	return d
}

func (d *Digit) WithOneToOne(sum bool) DigitInterface {
	d.sumOneToOne = sum
	return d
}

func (d *Digit) WithReplacement(value string, keys []int) DigitInterface {
	if d.replacement == nil {
		d.replacement = make(map[int]string)
	}
	for _, v := range keys {
		d.replacement[v] = value
	}
	return d
}

func (d *Digit) GetMultiplier(key int) int {
	if value, exist := d.multipliers[key]; !exist {
		return 0
	} else {
		return value
	}
}

func (d *Digit) GetReplacement(key int) string {
	if value, exist := d.replacement[key]; !exist {
		return "0"
	} else {
		return value
	}
}

func (d *Digit) NextMultiplier(multiplier int) int {
	multiplier++
	if multiplier == len(d.multipliers) {
		return 0
	}
	return multiplier
}

func (d *Digit) SumDigit(value int) int {
	return (value / 10) + (value % 10)
}

func (d *Digit) ContainsKey(key int) bool {
	_, exist := d.replacement[key]
	return exist
}

func (d *Digit) Calc() string {
	sum, turnMultiplier := 0, 0
	for _, v := range d.numbers {
		multiplier := d.GetMultiplier(turnMultiplier)
		amount := v * multiplier
		sum = d.sum(sum, amount)
		turnMultiplier = d.NextMultiplier(turnMultiplier)
	}

	result := sum % d.module

	if d.complement {
		result = d.module - result
	}

	if d.ContainsKey(result) {
		return d.GetReplacement(result)
	}

	return strconv.Itoa(result)
}

func (d *Digit) sum(sum int, total int) int {
	if d.sumOneToOne {
		sum += d.SumDigit(total)
	} else {
		sum += total
	}
	return sum
}
