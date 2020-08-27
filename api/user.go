package api

import (
	"net/http"
	"strail/db"
	"strail/helper"
	"strail/models"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

// Display All Users
func GetUser(c echo.Context) error {
	db := db.ConnectDB()
	users := []models.User{}
	db.Find(&users)
	return c.JSON(http.StatusOK, users)
}

// GET User By ID
func GetUserByID(c echo.Context) error {
	ID := c.Param("id")
	db := db.ConnectDB()
	users := []models.User{}
	db.Where("id= ?", ID).Find(&users)

	if users == nil {
		return c.JSON(http.StatusNotFound, "Data Not Found")
	}

	return c.JSON(http.StatusOK, users)
}

// CreateUser
func CreateUser(c echo.Context) error {
	var (
		v        = validator.New()
		name     = c.FormValue("name")
		username = c.FormValue("username")
		password = c.FormValue("password")
	)
	db := db.ConnectDB()
	hash, _ := helper.HashPassword(password)
	users := models.User{Name: name, Username: username, Password: hash}
	err := v.Struct(users)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	db.Create(&users)
	return c.JSON(http.StatusOK, users)
}

// UpdateUser
func UpdateUser(c echo.Context) error {

	var (
		v    = validator.New()
		ID   = c.Param("id")
		name = c.FormValue("name")
		db   = db.ConnectDB()
	)

	users := models.User{}
	db.Where("ID= ?", ID).Find(&users)
	users.Name = name
	err := v.Struct(users)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	db.Save(&users)
	return c.JSON(http.StatusOK, users)
}

// DeleteUser By ID
func DeleteUser(c echo.Context) error {
	var (
		ID = c.Param("id")
		db = db.ConnectDB()
	)

	users := models.User{}
	err := db.Where("id= ?", ID).Delete(&users)
	if err == nil {
		return c.String(http.StatusNotFound, "User Not Found")
	} else {
		return c.String(http.StatusOK, "Delete Success ID: "+ID)
	}

}
