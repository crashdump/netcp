package route

import (
	"github.com/gin-gonic/gin"

	"github.com/crashdump/netcp/internal/controller"
)

func NewStatus(p *controller.StatusController, r *gin.RouterGroup) {
	authRoute := r.Group("/status")

	// swagger:operation GET /auth auth getAuths
	//
	// Returns list of all auth methods.
	//
	// ---
	// consumes:
	//   - "application/json"
	// produces:
	//   - "application/json"
	// responses:
	//   '200':
	//     description: auth list response
	//     schema:
	//	        "$ref": "#/definitions/AuthListResponse"
	//
	//   default:
	//      description: General error
	//      schema:
	//	        "$ref": "#/definitions/GeneralError"
	//
	authRoute.GET("", p.Get)
}