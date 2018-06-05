package interfaces

//groupInterface as contract between ORM level and Service Level
type groupInterface interface {
	CreateGroup(groupName, groupOwner string, groupType uint) bool
	//Delete(group *models.Group)
}

//GI is the type of groupInterface
type GI groupInterface
