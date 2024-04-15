package models

import "gorm.io/gorm"

/*
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	all thease fields will be added by gorm.Model
*/

type Post struct {
	gorm.Model 
	Title string
	Body string	
}