package entity

import "time"

type Account struct {
	ID        string
	Balance   float64
	UpdatedAt time.Time
}

func NewAccount(id string, balance float64) *Account {
	account := &Account{
		ID:        id,
		Balance:   balance,
		UpdatedAt: time.Now(),
	}
	return account
}

func (a *Account) UpdateBalance(balance float64) {
	a.Balance = balance
}
