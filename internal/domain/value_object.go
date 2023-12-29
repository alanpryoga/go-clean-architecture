package domain

type Currency string

const (
	IDR Currency = "IDR"
)

type Money struct {
	Currency Currency
	Value    int64
}
