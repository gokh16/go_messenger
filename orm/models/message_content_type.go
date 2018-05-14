package models

import (
	"github.com/jinzhu/gorm"
)

type Message_Content_Type struct {
	gorm.Model

	Type string
}