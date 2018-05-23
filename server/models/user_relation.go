package models

import (
	"github.com/jinzhu/gorm"
)

type UserRelation struct {
	gorm.Model

	RelationType RelationType

	RelatingUser   int `sql:"type:int REFERENCES users(id)"`
	RelatedUser    int `sql:"type:int REFERENCES users(id)"`
	RelationTypeID int
}
