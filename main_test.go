package main

import (
	"churner/data"
	"testing"
	"time"
)

func TestReadTransaction(t *testing.T) {

	expectedTransaction := data.Transaction{
		Date: time.Date(2019, 4, 27, 0, 0, 0, 0, time.UTC),
		Description: "Five Guys",
		O_Description: "FIVE GUYS",
		Amount: 13.12,
		TransactionType: "debit",
		Category: "Fast Food",
		AccountName: "Blue Cash Preferred",
		Labels: "",
		Notes: "",
	}

	line := []string{"4/27/2019","Five Guys","FIVE GUYS","13.12","debit","Fast Food","Blue Cash Preferred","",""}

	trans := readTransaction(line)

	if trans != expectedTransaction {
		t.Errorf("Transaction was incorrect, got %s, want %s", trans.String(), expectedTransaction.String())
	}
}