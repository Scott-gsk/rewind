package models

import "github.com/WeOps-Lab/rewind/lib/pkgs/database"

func init() {
	database.RegisterModel(
		&CloudRegion{},
	)
	database.RegisterModel(
		&Variable{},
	)
	database.RegisterModel(
		&Node{},
	)
	database.RegisterModel(
		&Collector{},
	)
	database.RegisterModel(
		&Configuration{},
	)
}
