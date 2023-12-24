package transaction

type Money struct {
	Currency string
	Value    float64
}

type TopUpRequest struct {
	Phone  string
	Amount Money
}

type PaymentRequest struct {
	CustomerPhone string
	MerchantPhone string
	Amount        Money
}

type WithdrawRequest struct {
	Phone  string
	Amount Money
}
