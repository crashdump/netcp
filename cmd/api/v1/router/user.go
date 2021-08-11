package router

import (
	"github.com/gin-gonic/gin"

	"github.com/crashdump/netcp/internal/controller"
)

func NewUser(p *controller.UserController, r *gin.RouterGroup) {

	userRoute := r.Group("/user")

	// swagger:operation GET /user user getUsers
	//
	// Returns list of all people.
	//
	// ---
	// consumes:
	//   - "application/json"
	// produces:
	//   - "application/json"
	// responses:
	//   '200':
	//     description: user list response
	//     schema:
	//	        "$ref": "#/definitions/UserListResponse"
	//
	//   default:
	//      description: General error
	//      schema:
	//	        "$ref": "#/definitions/GeneralError"
	//
	userRoute.GET("", p.List)

	// swagger:operation GET /user/{id} user getUser
	//
	// Returns user details of given user id.
	//
	// ---
	// consumes:
	//   - "application/json"
	// produces:
	//   - "application/json"
	// parameters:
	//   -
	//     in: "path"
	//     name: "id"
	//     description: "User id which is require to fetch user details."
	//     required: true
	//     schema:
	//       type: string
	// responses:
	//   '200':
	//     description: user get response
	//     schema:
	//	        "$ref": "#/definitions/UserResponse"
	//
	//   default:
	//      description: General error
	//      schema:
	//	        "$ref": "#/definitions/GeneralError"
	//
	userRoute.GET("/:id", p.Get)

	// swagger:operation POST /user user addUser
	//
	// Insert given new user details in people.
	//
	// ---
	// consumes:
	//   - "application/json"
	// produces:
	//   - "application/json"
	// parameters:
	//   -
	//     in: "body"
	//     name: "body"
	//     description: "User object that needs to be added to the people"
	//     required: true
	//     schema:
	//          "$ref": "#/definitions/User"
	// responses:
	//   '200':
	//     description: user add response
	//     schema:
	//       type: object
	//       required:
	//         - id
	//       properties:
	//         id:
	//           type: string
	//
	//   default:
	//      description: General error
	//      schema:
	//	        "$ref": "#/definitions/GeneralError"
	//
	userRoute.POST("", p.Post)

	// swagger:operation PUT /user/{id} user updateUser
	//
	// Update given user details in people.
	//
	// ---
	// consumes:
	//   - "application/json"
	// produces:
	//   - "application/json"
	// parameters:
	//   -
	//     in: "path"
	//     name: "id"
	//     description: "User id which is require to fetch user details."
	//     required: true
	//     schema:
	//       type: string
	//   -
	//     in: "body"
	//     name: "body"
	//     description: "User object that needs to be update in the people"
	//     required: true
	//     schema:
	//          "$ref": "#/definitions/User"
	// responses:
	//   '200':
	//     description: user update response
	//     schema:
	//	        "$ref": "#/definitions/UserResponse"
	//
	//   default:
	//      description: General error
	//      schema:
	//	        "$ref": "#/definitions/GeneralError"
	//
	userRoute.PUT("/:id", p.Put)

	// swagger:operation DELETE /user/{id} user deleteUser
	//
	// Delete user details of given user id.
	//
	// ---
	// consumes:
	//   - "application/json"
	// produces:
	//   - "application/json"
	// parameters:
	//   -
	//     in: "path"
	//     name: "id"
	//     description: "User id which is require to delete user details."
	//     required: true
	//     schema:
	//       type: string
	// responses:
	//   '200':
	//     description: user delete response
	//     schema:
	//       type: string
	//
	//   default:
	//      description: General error
	//      schema:
	//	        "$ref": "#/definitions/GeneralError"
	//
	userRoute.DELETE("/:id", p.Delete)
}