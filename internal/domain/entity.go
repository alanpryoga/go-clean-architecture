package domain

import "time"

type User struct {
	ID                string
	Name              string
	Phone             string
	BankAccountNumber string
	UserType          UserType
}

func (u User) IsEligibleToTopUp() bool {
	return u.UserType == CUSTOMER
}

func (u User) IsEligibleToPayment() bool {
	return u.UserType == CUSTOMER
}

func (u User) IsEligibleToWithdraw() bool {
	return u.UserType == MERCHANT
}

type Wallet struct {
	ID      string
	UserID  string
	Balance Money
}

func (w Wallet) IncreaseBalance(amount Money) error {
	if w.Balance.Currency != amount.Currency {
		return ErrCurrencyMismatch
	}

	w.Balance.Value += amount.Value
	return nil
}

func (w Wallet) DecreaseBalance(amount Money) error {
	if w.Balance.Currency != amount.Currency {
		return ErrCurrencyMismatch
	}

	if amount.Value > w.Balance.Value {
		return ErrBalanceInsufficient
	}

	w.Balance.Value -= amount.Value
	return nil
}

type Transaction struct {
	ID               string
	CreditedWalletID string
	DebitedWalletID  string
	Amount           Money
	Timestamp        time.Time
}
