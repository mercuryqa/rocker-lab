package model

import "time"

// Category — аналог enum Category из proto
type Category int

const (
	CategoryUnknown  Category = 0
	CategoryEngine   Category = 1
	CategoryFuel     Category = 2
	CategoryPorthole Category = 3
	CategoryWing     Category = 4
)

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

// Part — основная сущность
type Part struct {
	UUID          string
	Name          string
	Description   string
	Price         float64
	StockQuantity int64
	Category      Category
	Dimensions    *Dimensions
	Manufacturer  *Manufacturer
	Tags          []string
	Metadata      map[string]*Value
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type GetPartRequest struct {
	UUID string
}

type GetPartResponse struct {
	UUID          string
	Name          string
	Description   string
	Price         float64
	StockQuantity int64
	Category      Category
	Dimensions    *Dimensions
	Manufacturer  *Manufacturer
	Tags          []string
	Metadata      map[string]*Value
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type ListPartsRequest struct {
	// не описано
}

type ListPartsResponse struct {
	// не описано
}
