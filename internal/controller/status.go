package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// StatusController is the resource for the status model
type StatusController struct {
}

// GET /auths
func (v StatusController) Get(c *gin.Context) {
	c.JSON(http.StatusOK, struct {
		status string
	}{
		"ok",
	})
}
