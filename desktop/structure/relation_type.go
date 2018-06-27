package structure

import (
	"github.com/jinzhu/gorm"
)

//RelationType is a model to Database table
type RelationType struct {
	gorm.Model

	Type string `json:"relation_type"`
}