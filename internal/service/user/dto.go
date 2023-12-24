package user

type RegisterCustomerRequest struct {
	Name  string
	Phone string
}

type RegisterMerchantRequest struct {
	Name              string
	Phone             string
	BankAccountNumber string
}

type User struct {
	ID   string
	Name string
}
