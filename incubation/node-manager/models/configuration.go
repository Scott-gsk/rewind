package models

import "gorm.io/gorm"

type Configuration struct {
	gorm.Model
	Name            string    `json:"name"`
	OperatingSystem string    `json:"operating_system"`
	CollectorID     uint      `json:"collector_id"`
	Collector       Collector `gorm:"foreignKey:CollectorID"`
	Nodes           []Node    `gorm:"many2many:configuration_nodes;"`
}

type ConfigurationNode struct {
	gorm.Model
	ConfigurationID uint `json:"configuration_id"`
	Configuration   Configuration
	NodeID          uint `json:"node_id"`
	Node            Node
}
