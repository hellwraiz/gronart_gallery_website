package users

import (
	"gronart_gallery_website/internal/auth"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type DBHandler struct {
	db *sqlx.DB
}

func (h *DBHandler) loginUser(c *gin.Context) {
	var login auth.Login
	db := h.db
	if err := c.ShouldBindJSON(&login); isError(err, "JSON error", http.StatusBadRequest, c) {
		return
	}

	token, err := Login(db, &login)
	if isError(err, "DB error", http.StatusInternalServerError, c) {
		return
	}
	c.JSON(http.StatusOK, token)
}

func (h *DBHandler) getAccount(c *gin.Context) {
	db := h.db
	uuid := c.MustGet("user_id").(string)
	user, err := GetUserByUUID(db, uuid)
	if isError(err, "DB error", http.StatusInternalServerError, c) {
		return
	}
	c.JSON(http.StatusOK, user)

}

func (h *DBHandler) createUser(c *gin.Context) {
	db := h.db
	var user User

	if err := c.ShouldBindJSON(&user); isError(err, "JSON error", http.StatusBadRequest, c) {
		return
	}
	if err := CreateUser(db, &user); isError(err, "DB error", http.StatusInternalServerError, c) {
		return
	}
	c.JSON(http.StatusCreated, user)

}

func (h *DBHandler) deleteUser(c *gin.Context) {
	db := h.db

	uuid := c.Param("uuid")
	if err := DeleteUser(db, uuid); isError(err, "DB error", http.StatusInternalServerError, c) {
		return
	}
	c.Status(http.StatusAccepted)

}

func InitRoutes(db *sqlx.DB, api *gin.RouterGroup) {
	users := api.Group("users/")
	users.Use(AuthMiddleware)

	h := DBHandler{db: db}

	users.POST("register/", auth.AdminAuthMiddleware(), h.createUser)
	users.POST("login/", h.loginUser)
	users.GET(":uuid", h.getAccount)
	users.DELETE(":uuid", auth.AdminAuthMiddleware(), h.deleteUser)

}
