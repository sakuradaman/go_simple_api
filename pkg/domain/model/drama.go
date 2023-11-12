package model

import "gorm.io/gorm"

type Drama struct {
	gorm.Model

	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"type:varchar(100);not null;unique_index"`
	Year int    `gorm:"type:int;not null"`
}
