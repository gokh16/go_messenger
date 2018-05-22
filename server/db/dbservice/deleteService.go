package dbservice
	
import (
	"../../models"
)

func DeleteGroup(groupName string){
	db := dbconnect()
	defer db.Close()
	db.Where("group_name = ?", groupName).Delete(&models.Group{})
}

