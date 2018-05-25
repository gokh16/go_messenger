<<<<<<< HEAD
package models

import (
	"github.com/jinzhu/gorm"
)

type UserRelation struct {
	gorm.Model

	RelationType RelationType

	RelatingUser   uint `sql:"type:int REFERENCES users(id)"`
	RelatedUser    uint `sql:"type:int REFERENCES users(id)"`
	RelationTypeID uint
}
=======
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
>>>>>>> 3661ec18fda6f6db02155e9be22dd834f0e1cd48
