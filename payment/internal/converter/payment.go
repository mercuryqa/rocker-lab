package converter

import (
	"github.com/mercuryqa/rocket-lab/payment/internal/model"
	generatedPaymentV1 "github.com/mercuryqa/rocket-lab/payment/pkg/proto/payment_v1"
)

func ToModelPaymentMethod(pm generatedPaymentV1.PaymentMethod) model.PaymentMethod {
	switch pm {
	case generatedPaymentV1.PaymentMethod_CARD:
		return model.PaymentMethodCard
	case generatedPaymentV1.PaymentMethod_SBP:
		return model.PaymentMethodSBP
	case generatedPaymentV1.PaymentMethod_CREDIT_CARD:
		return model.PaymentMethodCreditCard
	case generatedPaymentV1.PaymentMethod_INVESTOR_MONEY:
		return model.PaymentMethodInvestorMoney
	default:
		return "" // пустое значение, не UNKNOWN
	}
}
