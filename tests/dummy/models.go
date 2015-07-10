package db

import (
	"time"
)

// this is a configration model
type Configuration struct {
	ConfigKey   string `gorm:"primary_key:yes" sql:"type:varchar(20);not null;unique"`
	ConfigValue string `sql:"type:varchar(255);`
	Comment     string `sql:"type:varchar(255);`
}
