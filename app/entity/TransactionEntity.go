package entity

import (
	"time"
	"gorm.io/gorm"
)

// swagger:model Transaction
type Transaction struct {
	ID               int            `gorm:"AUTO_INCREMENT" json:"id"`
	Uuid 			 string	    	`gorm:"unique" json:"uuid"`
	UserId 			 int    		`json:"user_id"`
	// example: "1970-01-01 00:00:00"
	DeviceTimestamp	 *time.Time  	`json:"device_timestamp,omitempty" default:"1970-01-01 00:00:00"`
	TotalAmount 	 int    		`json:"total_amount"`
	PaidAmount 		 int    		`json:"paid_amount"`
	ChangeAmount 	 int    		`json:"change_amount"`
	PaymentMethod	 string    		`json:"payment_method"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `json:"deleted_at"`
}

type TransactionCreate struct {
	ID               int            `json:"-"`
	Uuid 			 string	    	`gorm:"unique" json:"uuid"`
	UserId 			 int    		`json:"user_id"`
	// example: "1970-01-01 00:00:00"
	DeviceTimestamp	 *time.Time  	`json:"device_timestamp,omitempty" default:"1970-01-01 00:00:00"`
	TotalAmount 	 int    		`json:"total_amount"`
	PaidAmount 		 int    		`json:"paid_amount"`
	ChangeAmount 	 int    		`json:"change_amount"`
	PaymentMethod	 string    		`json:"payment_method"`
}

type TransactionDelete struct {
	ID int `json:"id"`
}
