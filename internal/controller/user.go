package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	"github.com/crashdump/netcp/internal/model"
	"github.com/crashdump/netcp/internal/repository"
)

// UserController is the resource for the User model
type UserController struct {
	userRepository repository.UserRepository
}

var (
	MsgSucessUserCreated      = "user created successfully"
	MsgSucessUserDeleted      = "user deleted successfully"
	MsgErrorInvalidID         = "invalid user id"
	MsgErrorNotFound          = "user not found"
	MsgErrorFieldMissingName  = "field name is mandatory"
	MsgErrorFieldMissingEmail = "field email is mandatory"
)

// GET /users
func (v UserController) List(c *gin.Context) {
	users := v.userRepository.FindAll()

	c.JSON(http.StatusOK, gin.H{"users": users})
}

// GET /user/:id
func (v UserController) Get(c *gin.Context) {
	//id, err := uuid.Parse(c.Param("id"))
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": MsgErrorInvalidID})
	//}
	//
	//user := models.User{ ID: id }
	//err = models.GetDB().Find(&user).Error
	//
	//switch err {
	//case gorm.ErrRecordNotFound:
	//	c.JSON(http.StatusNotFound, gin.H{"error": MsgErrorNotFound})
	//case nil:
	//	c.JSON(http.StatusOK, gin.H{"result": user})
	//default:
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//}
}

// POST /user
func (v UserController) Post(c *gin.Context) {
	var user model.User

	err := c.BindJSON(&user)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	if user.Name == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": MsgErrorFieldMissingName})
		return
	}

	if user.Email == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": MsgErrorFieldMissingEmail})
		return
	}

	v.userRepository.Save(user)

	c.JSON(http.StatusOK, gin.H{"status": MsgSucessUserCreated})
}

// PUT /user
func (v UserController) Put(c *gin.Context) {
	//var user models.User
	//c.Bind(&user)
	//
	//if user.Name == "" {
	//	c.JSON(http.StatusUnprocessableEntity, gin.H{"status": MsgErrorFieldMissingName})
	//}
	//
	//if user.Email == "" {
	//	c.JSON(http.StatusUnprocessableEntity, gin.H{"status": MsgErrorFieldMissingEmail})
	//}
	//
	//err := models.GetDB().Create(&user).Error // INSERT INTO "users" (name) VALUES (user.Name);
	//
	//switch err {
	//case gorm.ErrRecordNotFound:
	//	c.JSON(http.StatusNotFound, gin.H{"error": MsgErrorNotFound})
	//case nil:
	//	c.JSON(http.StatusOK, gin.H{"success": user})
	//default:
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//}
}

// DELETE /user/:id
func (v UserController) Delete(c *gin.Context) {
	//// Get id user
	//id := c.Params.ByName("id")
	//
	//var user models.User
	//err := models.GetDB().First(&user, id).Error // SELECT * FROM users WHERE id = $id;
	//
	//switch err {
	//case gorm.ErrRecordNotFound:
	//	c.JSON(http.StatusNotFound, gin.H{"error": MsgErrorNotFound})
	//case nil:
	//	models.GetDB().Delete(&user) // DELETE FROM users WHERE id = user.ID
	//	c.JSON(http.StatusOK, gin.H{"success": MsgSucessUserDeleted + id})
	//default:
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//}
}