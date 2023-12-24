package domain

type Currency string

const (
	IDR Currency = "IDR"
)

type UserType string

const (
	CUSTOMER UserType = "CUSTOMER"
	MERCHANT UserType = "MERCHANT"
)

type Money struct {
	Currency Currency
	Value    float64
}
