package model

import "time"

type OrderRequest struct {
	UserUuid  string   `json:"user_uuid"`
	PartUuids []string `json:"part_uuids"`
}

func (o *OrderRequest) GetUserUuid() string {
	return o.UserUuid
}

func (o *OrderRequest) GetPartUuids() []string {
	return o.PartUuids
}

type OrderResponse struct {
	OrderUuid  string  `json:"order_uuid"`
	TotalPrice float64 `json:"total_price"`
}

// Part — основная сущность
type Part struct {
	UUID          string
	Name          string
	Description   string
	Price         float64
	StockQuantity int64
	Category      Category
	Dimensions    Dimensions
	Manufacturer  Manufacturer
	Tags          []string
	Metadata      map[string]*Value
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
type PartsFilter struct {
	Uuids                 []string
	Names                 []string
	Categories            []Category
	ManufacturerCountries []string
	ManufacturerNames     []string
	Tags                  []string
}

// Dimensions — размеры детали
type Dimensions struct {
	Length float64
	Width  float64
	Height float64
	Weight float64
}

// Manufacturer — информация о производителе
type Manufacturer struct {
	Name    string
	Country string
	Website string
}

// Value — универсальное значение (string/int/double/bool)
type Value struct {
	StringValue *string
	Int64Value  *int64
	DoubleValue *float64
	BoolValue   *bool
}

// Category — аналог enum Category из proto
type Category int

const (
	CategoryUnknown  Category = 0
	CategoryEngine   Category = 1
	CategoryFuel     Category = 2
	CategoryPorthole Category = 3
	CategoryWing     Category = 4
)

type GetPartRequest struct {
	InventoryUuid string
}

type GetPartResponse struct {
	Part Part
}

type ListPartsResponse struct {
	Parts []Part
}

type PaymentRequest struct {
	PaymentMethod PaymentMethod `json:"payment_method"`
}
type PaymetnResponse struct {
	TransactionUuid string `json:"transaction_uuid"`
}

type GetOrderResponse struct {
	OrderUuid       string   `json:"order_uuid"`
	UserUuid        string   `json:"user_uuid"`
	PartUuids       []string `json:"part_uuids"`
	TotalPrice      float64  `json:"total_price"`
	TransactionUuid string   `json:"transaction_uuid"`
	PaymentMethod   string   `json:"payment_method"`
	Status          string   `json:"status"`
}

type GetOrder struct {
	OrderUuid string `json:"order_uuid"`
}

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
