package api

import (
	"net/http"

	validator "github.com/online.scheduling-api/src/api/validators"
	service "github.com/online.scheduling-api/src/business/services"

	"github.com/gin-gonic/gin"
	"github.com/online.scheduling-api/src/business/models"
	"github.com/online.scheduling-api/src/helpers"
)

func GetAll(c *gin.Context) {
	users, err := service.GetAllUsers()

	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}

	c.IndentedJSON(http.StatusOK, &users)
}

func GetById(c *gin.Context) {
	id, err := helpers.GetUUID(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	user, err := service.GetUserById(id)
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}
	if user == nil {
		c.IndentedJSON(http.StatusNotFound, &user)
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func Create(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	if err := validator.Validate(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	isDuplicated, err := service.CreateNewUser(&user)
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, err)
		return
	}
	if isDuplicated {
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "Usuário já cadastrado"})
		return
	}

	c.IndentedJSON(http.StatusCreated, &user)
}

func Update(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	if err := validator.Validate(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	if err := service.UpdateUser(&user); err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, err)
		return
	}

	c.IndentedJSON(http.StatusNoContent, nil)
}

func Delete(c *gin.Context) {
	id, err := helpers.GetUUID(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	found, err := service.DeleteUserById(id)
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, err)
		return
	}
	if !found {
		c.IndentedJSON(http.StatusNotFound, nil)
		return
	}

	c.IndentedJSON(http.StatusNoContent, nil)
}
