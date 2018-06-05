package models

import (
	"github.com/jinzhu/gorm"
)

//MessageContentType is a model to Database table.
type MessageContentType struct {
	gorm.Model

	Type string `json:"message_content_type"`
}
