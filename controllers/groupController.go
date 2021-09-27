package controllers

import (
	"net/http"

	"fast.bibabo.vn/repositories"
	groupService "fast.bibabo.vn/services/group_services"
	"github.com/gin-gonic/gin"
)

type groupController struct {
	groupService groupService.GroupService
}

func InstanceGroupController(groupRepo repositories.GroupRepository) *groupController {
	gs := groupService.InstanceGroupService(groupRepo)
	return &groupController{
		groupService: gs,
	}
}

func (gc *groupController) FindAll(c *gin.Context) {
	groups := gc.groupService.FindAll(0, 10)
	c.JSON(http.StatusOK, gin.H{"message": "Ok", "data": groups})
}
