package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"github.com/crashdump/netcp/internal/model"
	"github.com/crashdump/netcp/internal/repository"
	"time"
)

// BlobController is the resource for the Blob model
type BlobController struct {
	blobRepository repository.BlobRepository
}

var (
	MsgSucessBlobCreated = "blob created successfully"
	MsgSucessBlobDeleted = "blob deleted successfully"
	//MsgErrorInvalidID         = "invalid blob id"
	//MsgErrorNotFound          = "blob not found"
	//MsgErrorFieldMissingName  = "field name is mandatory"
	MsgErrorFieldMissingOS = "field operating_system is mandatory"
)

// GET /blobs
func (v BlobController) List(c *gin.Context) {
	blobs := v.blobRepository.FindAll()

	c.JSON(http.StatusOK, gin.H{"blobs": blobs})
}

// GET /blob/:id
func (v BlobController) Get(c *gin.Context) {
	//id, err := uuid.Parse(c.Param("id"))
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": MsgErrorInvalidID})
	//}
	//
	//blob := models.Blob{ ID: id }
	//err = models.GetDB().Find(&blob).Error
	//
	//switch err {
	//case gorm.ErrRecordNotFound:
	//	c.JSON(http.StatusNotFound, gin.H{"error": MsgErrorNotFound})
	//case nil:
	//	c.JSON(http.StatusOK, gin.H{"result": blob})
	//default:
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//}
}

// POST /blob
func (v BlobController) Post(c *gin.Context) {
	var blob model.Blob
	blob.UploadedAt = time.Now()

	err := c.BindJSON(&blob)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	v.blobRepository.Save(blob)

	blob = v.blobRepository.FindByID(blob.ID)

	c.JSON(http.StatusOK, blob)
}

// PUT /blob
func (v BlobController) Put(c *gin.Context) {
	//var blob models.Blob
	//c.Bind(&blob)
	//
	//if blob.Name == "" {
	//	c.JSON(http.StatusUnprocessableEntity, gin.H{"status": MsgErrorFieldMissingName})
	//}
	//
	//if blob.Email == "" {
	//	c.JSON(http.StatusUnprocessableEntity, gin.H{"status": MsgErrorFieldMissingEmail})
	//}
	//
	//err := models.GetDB().Create(&blob).Error // INSERT INTO "blobs" (name) VALUES (blob.Name);
	//
	//switch err {
	//case gorm.ErrRecordNotFound:
	//	c.JSON(http.StatusNotFound, gin.H{"error": MsgErrorNotFound})
	//case nil:
	//	c.JSON(http.StatusOK, gin.H{"success": blob})
	//default:
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//}
}

// DELETE /blob/:id
func (v BlobController) Delete(c *gin.Context) {
	//// Get id blob
	//id := c.Params.ByName("id")
	//
	//var blob models.Blob
	//err := models.GetDB().First(&blob, id).Error // SELECT * FROM blobs WHERE id = $id;
	//
	//switch err {
	//case gorm.ErrRecordNotFound:
	//	c.JSON(http.StatusNotFound, gin.H{"error": MsgErrorNotFound})
	//case nil:
	//	models.GetDB().Delete(&blob) // DELETE FROM blobs WHERE id = blob.ID
	//	c.JSON(http.StatusOK, gin.H{"success": MsgSucessBlobDeleted + id})
	//default:
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//}
}