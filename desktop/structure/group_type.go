package structure

import (
	"github.com/jinzhu/gorm"
)

//GroupType is a model to Database table.
type GroupType struct {
	gorm.Model

	Type uint `json:"group_type"`
}
