package models

import "gorm.io/gorm"

type Collector struct {
	gorm.Model
	Name                    string                   `json:"name"`
	Details                 string                   `json:"details"`
	OperatingSystem         string                   `json:"operating_system"`
	Introduction            string                   `json:"introduction"`
	CollectorConfigurations []CollectorConfiguration `gorm:"foreignKey:CollectorID"`
}

func (Collector) TableName() string {
	return "collectors"
}
