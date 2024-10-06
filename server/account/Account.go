package account

import (
	"math/rand"
)

type Account struct {
	ID      int64
	Name    string
	Balance float64
	accType string
}

func getUUID() int64 {
  var uuid int64
  for i := 0; i < 10; i++ {
    uuid += rand.Int63()
  }
  return uuid
}

func New(name string, balance float64) *Account {
	account := Account{
		ID:      getUUID(),
		Name:    name,
		Balance: balance,
	}
	return &account
}

func (a *Account) checkBalance() float64 {
	return a.Balance
}

func (a *Account) deposit(money float64) {
	a.Balance += money
}

func (a *Account) withdraw(money float64) {
	if a.accType == "debit" &&
		a.Balance >= money &&
		money > 0 &&
		money <= 10000 {
		a.Balance -= money
	}
}

func (a *Account) settings(name string) {
	if name != "" {
		a.Name = name
	}
}
