package model

type PaymentMethod string

const (
	PaymentMethodUnknown       PaymentMethod = "UNKNOWN"
	PaymentMethodCard          PaymentMethod = "CARD"
	PaymentMethodSBP           PaymentMethod = "SBP"
	PaymentMethodCreditCard    PaymentMethod = "CREDIT_CARD"
	PaymentMethodInvestorMoney PaymentMethod = "INVESTOR_MONEY"
)

type PayOrderRequest struct {
	OrderUuid     string
	UserUuid      string
	PaymentMethod PaymentMethod
}

type PayOrderResponse struct {
	TransactionUuid string
}
