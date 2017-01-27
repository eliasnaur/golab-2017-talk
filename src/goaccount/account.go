// START 1 OMIT
package goaccount

import (
	"fmt"
	"time"
)

type Account struct {
	balance      float64
	transactions []*Transaction
}

func NewAccount() *Account { return &Account{} }

type Transaction struct {
	amount float64
	time   time.Time
}

// END 1 OMIT
// START 2 OMIT
func (a *Account) FormatBalance() string {
	return fmt.Sprintf("€%.2f", a.balance)
}

func (a *Account) NewTransaction(amount float64) {
	a.transactions = append(a.transactions, &Transaction{
		amount: amount,
		time:   time.Now(),
	})
	a.balance += amount
}

func (a *Account) NumTransactions() int {
	return len(a.transactions)
}

func (a *Account) Transaction(pos int) *Transaction {
	return a.transactions[pos]
}

// END 2 OMIT

func (t *Transaction) FormatAmount() string {
	return fmt.Sprintf("€%.2f", t.amount)
}

func (t *Transaction) FormatTime() string {
	return t.time.Format("2006-02-01 15:04:05")
}
