package models

import (
	"github.com/jinzhu/gorm"
)

type User_Relation struct {
	gorm.Model
	
	Relation_Type Relation_Type

	Relating_user uint `sql:"type:int REFERENCES users(id)"`
	Related_user uint `sql:"type:int REFERENCES users(id)"`
	Relation_typeID uint 
}