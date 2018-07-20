package structure

import (
	"github.com/jinzhu/gorm"
)

//UserRelation is a model to Database table
type UserRelation struct {
	gorm.Model

	RelationType RelationType

	RelatingUser   uint `sql:"type:int REFERENCES users(id)"`
	RelatedUser    uint `sql:"type:int REFERENCES users(id)"`
	RelationTypeID uint
}