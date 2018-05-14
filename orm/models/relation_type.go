package models

import (
	"github.com/jinzhu/gorm"
)

type Relation_Type struct {
	gorm.Model

	Type string
}