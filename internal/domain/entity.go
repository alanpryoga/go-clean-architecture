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

func (w *Wallet) AddBalance(amount Money) error {
	newBalance, err := w.Balance.Add(amount)
	if err != nil {
		return err
	}

	w.Balance = newBalance
	return nil
}

func (w *Wallet) SubBalance(amount Money) error {
	newBalance, err := w.Balance.Sub(amount)
	if err != nil {
		return err
	}

	w.Balance = newBalance
	return nil
}

type Transaction struct {
	ID               string
	CreditedWalletID string
	DebitedWalletID  string
	Amount           Money
	Timestamp        time.Time
}
