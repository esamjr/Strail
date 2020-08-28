package api

import (
	"fmt"
	"net/http"
	"strail/db"
	"strail/helper"
	"strail/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
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
		username = c.FormValue("username")
		password = c.FormValue("password")
	)
	db := db.ConnectDB()
	hash, _ := helper.HashPassword(password)
	users := models.User{Username: username, Password: hash}
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
		v        = validator.New()
		ID       = c.Param("id")
		username = c.FormValue("username")
		db       = db.ConnectDB()
	)

	users := models.User{}
	db.Where("ID= ?", ID).Find(&users)
	users.Username = username
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

func Login(c echo.Context) error {
	var (
		v        = validator.New()
		username = c.FormValue("username")
		password = c.FormValue("password")
	)
	db := db.ConnectDB()
	users := models.User{Username: username, Password: password}
	err := v.Struct(users)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if errs := db.Where("username = ?", username).First(&users).Error; errs != nil {
		return c.JSON(http.StatusNotFound, "Username Not Found")
	}
	expiresAt := time.Now().Add(time.Minute * 100000).Unix()

	errf := bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return c.JSON(http.StatusNotFound, "Invalid login credentials. Please try again")
	}

	tk := &models.Token{
		UserID:   users.ID,
		Username: users.Username,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"user":     users.Username,
		"password": users.Password,
		"message":  "You Are Logged",
		"token":    tokenString,
	})

}
