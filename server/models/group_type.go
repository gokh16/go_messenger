package models

import (
	"github.com/jinzhu/gorm"
)

type Group_Type struct {
	gorm.Model

	Type string
}