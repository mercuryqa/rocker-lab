package model

type Order struct {
	OrderUuid       string
	UserUuid        string
	PartUuids       []string
	TotalPrice      float64
	TransactionUuid string
	PaymentMethod   PaymentMethod
	Status          OrderStatus
}

type PaymentMethod string

const (
	Unknown       PaymentMethod = "UNKNOWN"
	Card          PaymentMethod = "CARD"
	Sbp           PaymentMethod = "SBP"
	CreditCard    PaymentMethod = "CREDIT_CARD"
	InvestorMoney PaymentMethod = "INVESTOR_MONEY"
)

type OrderStatus string

const (
	PendingPayment OrderStatus = "PENDING_PAYMENT"
	Paid           OrderStatus = "PAID"
	Cancelled      OrderStatus = "CANCELLED"
)
