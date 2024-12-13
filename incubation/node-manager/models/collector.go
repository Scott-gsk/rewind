package models

import "gorm.io/gorm"

type Collector struct {
	gorm.Model
	Name            string          `json:"name"`
	Details         []byte          `json:"details"` // []byte to store Markdown content
	OperatingSystem string          `json:"operating_system"`
	Introduction    string          `json:"introduction"`
	Configurations  []Configuration `gorm:"foreignKey:CollectorID"`
}
