package models

import (
	"github.com/jinzhu/gorm"
)

//GroupType is a model to Database table.
type GroupType struct {
	gorm.Model

	Type string `json:"group_type"`
}
