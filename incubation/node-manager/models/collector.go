package models

import "gorm.io/gorm"

type Collector struct {
	gorm.Model
	Name            string          `json:"name"`
	Details         string          `json:"details"`
	OperatingSystem string          `json:"operating_system"`
	Introduction    string          `json:"introduction"`
	Configurations  []Configuration `gorm:"foreignKey:CollectorID"`
}
