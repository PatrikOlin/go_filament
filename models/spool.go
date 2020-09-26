package models

import (
	"time"

	_ "github.com/jinzhu/gorm"  
)							  
							  
type Spool struct {
	ID uint `json:"id" gorm:"primary_key, AUTO_INCREMENT"`
	Tag string `json:"tag"`
	Brand string `json:"brand"`
	Name string `json:"name"`
	Weight int `json:"weight"`
	SpoolWeight int `json:"spool_weight"`
	Color string `json:"color"`
	Material string `json:"material"`
	NozzleTemp int `json:"nozzle_temp"`
	PlateTemp int `json:"plate_temp"`
	PricePerKg int `json:"price_per_kg"`
	Superpowers []Superpower `json:"superpowers" gorm:"many2many:spools_superpowers"`
	Notes string `json:"notes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type Superpower struct {
	ID uint `json:"id" gorm:"primary_key, AUTO_INCREMENT"`
	Name string `json:"superpower" gorm:"<-:false"`
}							  
