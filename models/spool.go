package models

import (
	"time"
)

type Spool struct {
	ID uint `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Tag string `json:"tag"`
	Brand string `json:"brand"`
	Name string `json:"name"`
	Weight int `json:"weight"`
	SpoolWeight int `json:"spool_weight"`
	Color string `json:"color"`
	Material string `json:"material"`
	Notes string `json:"notes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
