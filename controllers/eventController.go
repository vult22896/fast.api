package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type EventController struct{}

type PropertiesRequest struct {
	PostId     int `json:"postId"`
	CategoryId int `json:"categoryId"`
	GroupId    int `json:"groupId"`
	TopicId    int `json:"topicId"`
}
type BatchRequest struct {
	Event      string            `json:"event"`
	Properties PropertiesRequest `json:"properties"`
}

type EventRequest struct {
	Batch []BatchRequest `json:"batch"`
}

func (e *EventController) Index(c *gin.Context) {
	var input EventRequest
	if err := c.ShouldBindBodyWith(&input, binding.JSON); err != nil {
		c.JSON(http.StatusOK, gin.H{"err": err})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"input": input})
}
