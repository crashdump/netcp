package router

import (
	"github.com/gin-gonic/gin"

	"github.com/crashdump/netcp/internal/controller"
)

func NewBlob(p *controller.BlobController, r *gin.RouterGroup) {

	blobRoute := r.Group("/blob")

	// swagger:operation GET /blob blob getBlobs
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
	//     description: blob list response
	//     schema:
	//	        "$ref": "#/definitions/BlobListResponse"
	//
	//   default:
	//      description: General error
	//      schema:
	//	        "$ref": "#/definitions/GeneralError"
	//
	blobRoute.GET("", p.List)

	// swagger:operation GET /blob/{id} blob getBlob
	//
	// Returns blob details of given blob id.
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
	//     description: "Blob id which is require to fetch blob details."
	//     required: true
	//     schema:
	//       type: string
	// responses:
	//   '200':
	//     description: blob get response
	//     schema:
	//	        "$ref": "#/definitions/BlobResponse"
	//
	//   default:
	//      description: General error
	//      schema:
	//	        "$ref": "#/definitions/GeneralError"
	//
	blobRoute.GET("/:id", p.Get)

	// swagger:operation POST /blob blob addBlob
	//
	// Insert given new blob details in people.
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
	//     description: "Blob object that needs to be added to the people"
	//     required: true
	//     schema:
	//          "$ref": "#/definitions/Blob"
	// responses:
	//   '200':
	//     description: blob add response
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
	blobRoute.POST("", p.Post)

	// swagger:operation PUT /blob/{id} blob updateBlob
	//
	// Update given blob details in people.
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
	//     description: "Blob id which is require to fetch blob details."
	//     required: true
	//     schema:
	//       type: string
	//   -
	//     in: "body"
	//     name: "body"
	//     description: "Blob object that needs to be update in the people"
	//     required: true
	//     schema:
	//          "$ref": "#/definitions/Blob"
	// responses:
	//   '200':
	//     description: blob update response
	//     schema:
	//	        "$ref": "#/definitions/BlobResponse"
	//
	//   default:
	//      description: General error
	//      schema:
	//	        "$ref": "#/definitions/GeneralError"
	//
	blobRoute.PUT("/:id", p.Put)

	// swagger:operation DELETE /blob/{id} blob deleteBlob
	//
	// Delete blob details of given blob id.
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
	//     description: "Blob id which is require to delete blob details."
	//     required: true
	//     schema:
	//       type: string
	// responses:
	//   '200':
	//     description: blob delete response
	//     schema:
	//       type: string
	//
	//   default:
	//      description: General error
	//      schema:
	//	        "$ref": "#/definitions/GeneralError"
	//
	blobRoute.DELETE("/:id", p.Delete)
}