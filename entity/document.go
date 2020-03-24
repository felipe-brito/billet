package entity

import (
	"fmt"
	"regexp"
)

// Define a document type
// @Value ticket value
// @OurNumber our Number
// @FebrabanType
// @TypeOurNumber
type Document struct {
	Value         float64
	OurNumber     string
	TypeOurNumber string
	FebrabanType  string
}

// GetFormattedValue
func (d *Document) GetFormattedValue() string {
	formattedValue := fmt.Sprintf("%011.2f", d.Value)
	regex := regexp.MustCompile("[\\W]")
	return regex.ReplaceAllString(formattedValue, "")
}