package domain

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

type Account struct {
	Id          string `json:"id"`
	URL         string `json:"url"`
	raw_balance string
	Year        int                 `json:"year"`
	Month       int                 `json:"month"`
	Products    map[string]*Product `json:"products"`
}

func NewAccount(year, month, raw_balance string) *Account {
	y, err := strconv.Atoi(year)
	if err != nil {
		log.Fatal("Error on name path_file \"year\"")
	}
	m, err := strconv.Atoi(month)
	if err != nil {
		log.Fatal("Error on name path_file \"month\"")
	}
	return &Account{
		Year:        y,
		Month:       m,
		raw_balance: raw_balance,
		Products:    make(map[string]*Product),
	}
}

func (a Account) String() string {
	return fmt.Sprintf("Account: %s Total_Products %d", a.Id, len(a.Products))
}

type Product struct {
	Id           string `json:"id"`
	Transactions []*Transaction
	Total        float32 `json:"total"`
}

func NewProduct(id string, total float32) *Product {
	return &Product{
		Id:           id,
		Transactions: []*Transaction{},
		Total:        total,
	}
}

func (p Product) String() string {
	return fmt.Sprintf("Product: %s Total: %f Transactions: %d:", p.Id, p.Total, len(p.Transactions))
}

type Transaction struct {
	Code        string
	Description string
	Amount      float32
	Date        *time.Time
}

func NewTransaction(code, description string, amount float32, date *time.Time) *Transaction {
	return &Transaction{
		Code:        code,
		Description: description,
		Amount:      amount,
		Date:        date,
	}
}

func (t Transaction) String() string {
	return fmt.Sprintf("Transaction: %s: %f: \n", t.Date, t.Amount)
}
