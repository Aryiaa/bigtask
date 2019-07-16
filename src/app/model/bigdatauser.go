package model

import "time"


type BigdataUser struct {
	ID int `gorm:"primary_key"`
	Name string `gorm:"type:varchar(64);not null"`
	Email string `gorm:"type:varchar(64);unique;not null"`
	AvatarUrl string `gorm:"type:varchar(255);"`
	UpdatedAt time.Time
	CreatedAt time.Time
	Permission int `gorm:"default:0;"`
}

func (BigdataUser) Tablename() string {
	return "bigdata_user"
}


