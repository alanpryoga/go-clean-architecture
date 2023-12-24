package domain

import "errors"

var (
	ErrUserNotFound      = errors.New("user is not exists")
	ErrUserAlreadyExists = errors.New("user is already exists")

	ErrWalletNotFound = errors.New("wallet is not exists")

	ErrCurrencyMismatch    = errors.New("currency is not match with the one used")
	ErrBalanceInsufficient = errors.New("balance is not sufficient")
)
