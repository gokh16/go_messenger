package models

import (
	"github.com/jinzhu/gorm"
)

type RelationType struct {
	gorm.Model

	Type string `json:"relation_type"`
}
