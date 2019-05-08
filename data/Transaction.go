package data

import (
	"fmt"
	"time"
)

type Transaction struct {
	Date time.Time
	Description string
	O_Description string
	Amount float64
	TransactionType string
	Category string
	AccountName string
	Labels string
	Notes string
}

func (t Transaction) String() string {
	return fmt.Sprintf("%s %s %s %.2f %s %s %s %s %s",
		t.Date.String(),
		t.Description,
		t.O_Description,
		t.Amount,
		t.TransactionType,
		t.Category,
		t.AccountName,
		t.Labels,
		t.Notes)
}