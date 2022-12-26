package entity

import (
	"gorm.io/gorm"
)

type PaymentOption struct {
	gorm.Model
	PaymentName string
}

func (PaymentOption) TableName() string {
	return "payment_options"
}