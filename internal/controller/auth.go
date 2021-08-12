package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/crashdump/netcp/internal/model"
)

// AuthController is the resource for the Oauth model
type AuthController struct {
}

// GET /auth/servers
func (v AuthController) List(c *gin.Context) {
	c.JSON(http.StatusOK, []model.Auth{{
		Type: "oauth2",
		Config: model.AuthConfig{
			// TODO: Read from the database
			Domain:   "netcp-dev.eu.auth0.com",
			ClientID: "9mRCkgSZkDdJjECOcJivRXdTF0crNbDZ",
		},
	},
	})
}
