package models

import (
	"github.com/jinzhu/gorm"
)

type MessageContentType struct {
	gorm.Model

	Type string `json:"message_content_type"`
}
