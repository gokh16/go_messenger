package models

import (
	"github.com/jinzhu/gorm"
)

type GroupType struct {
	gorm.Model

	Type string `json:"group_type"`
}
