package domain

import "time"

type UserType string

const (
	Customer UserType = "Customer"
	Merchant UserType = "Merchant"
)

type User struct {
	ID                string
	Name              string
	Phone             string
	BankAccountNumber string
	UserType          UserType
}

func (u User) IsEligibleToTopUp() bool {
	return u.UserType == Customer
}

func (u User) IsEligibleToPayment() bool {
	return u.UserType == Customer
}

func (u User) IsEligibleToWithdraw() bool {
	return u.UserType == Merchant
}

type Wallet struct {
	ID      string
	UserID  string
	Balance Money
}

func (w Wallet) IncreaseBalance(amount Money) error {
	_, err := w.Balance.Add(amount)
	return err
}

func (w Wallet) DecreaseBalance(amount Money) error {
	_, err := w.Balance.Sub(amount)
	return err
}

type Transaction struct {
	ID               string
	CreditedWalletID string
	DebitedWalletID  string
	Amount           Money
	Timestamp        time.Time
}
