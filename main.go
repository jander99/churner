package main

import (
	"churner/data"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)
import "github.com/araddon/dateparse"


func readTransaction(line []string) data.Transaction {

	date, err := dateparse.ParseAny(line[0])
	if err != nil {
		panic(err)
	}

	f, err := strconv.ParseFloat(line[3], 64)
	if err != nil {
		panic(err)
	}

	t := data.Transaction{
		date,
		line[1],
		line[2],
		f,
		line[4],
		line[5],
		line[6],
		line[7],
		line[8],
	}

	return t
}


func main() {

	f, err := os.Open("resources/transactions.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close();

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}

	for _, line := range lines[1:] {
		trans := readTransaction(line)
		fmt.Println(trans)
	}
}