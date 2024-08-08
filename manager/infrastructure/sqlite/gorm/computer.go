package gorm

import "gorm.io/gorm"

// Computer のモデル。
type Computer struct {
	gorm.Model

	// Computer の ID。
	ID string `gorm:"primaryKey"`
	// Computer の名前。
	Name string
	// Computer の MAC アドレス。
	MacAddress string `gorm:"unique"`
}
