package main

import (
	"time"
	// "github.com/jinzhu/gorm"
)

type Records []Record
type Record struct {
	// If valid belowe, add columns id, created_at, updated_at and deleted_at automatically.
	// gorm.Model

	ID           int        `gorm:"AUTO_INCREMENT" json:"id"` // set id to auto incrementable
	Name         string     `json:"name"`
	Age          int        `json:"age"`
	Birthday     *time.Time `json:"birthday"`
	Email        string     `gorm:"type:varchar(100);unique_index" json:"email"`
	Role         string     `gorm:"size:255" json:"role"`                 // set field size to 255
	MemberNumber *string    `gorm:"unique;not null" json:"member_number"` // set member number to unique and not null
	Address      string     `gorm:"index:addr" json:"address"`            // create index with name `addr` for address
	IgnoreMe     int        `gorm:"-" json:"ignore_me"`                   // ignore this field
}

// Specify file name
var dbFileName = "analysis.db"
