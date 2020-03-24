package entity

import (
	"errors"
	"time"
)

var (
	dateDueFixed = time.Date(1997, 10, 07, 0, 0, 0, 0, time.UTC)
)

// Define
// @DateDue
// @DateOfIssue
type Dates struct {
	DateDue     time.Time
	DateOfIssue time.Time
}

// Define
func (d *Dates) DateDueFactor() (int, error) {
	if factor := int(d.DateDue.Sub(dateDueFixed).Hours() / 24); factor <= 0 {
		return 0, errors.New("expiration date should be in the future")
	} else {
		return factor, nil
	}
}
