package api

import (
	"net/http"
	"strail/db"
	"strail/models"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

// GetShoppingList Functuion
func GetShoppingList(c echo.Context) error {
	db := db.ConnectDB()
	list := []models.ShoppingList{}
	db.Find(&list)
	if list == nil {
		return c.JSON(http.StatusNotFound, "Data Not Found")
	} else {
		return c.JSON(http.StatusOK, list)
	}
}

// GetShoppingListByID Function
func GetShoppingListByID(c echo.Context) error {
	ID := c.Param("id")
	db := db.ConnectDB()
	data := []models.ShoppingList{}
	db.Where("id= ?", ID).Find(&data)

	if data == nil {
		return c.JSON(http.StatusNotFound, "Data Not Found")
	}

	return c.JSON(http.StatusOK, data)
}

// CreateShoppingList Function
func CreateShoppingList(c echo.Context) error {
	var (
		v     = validator.New()
		name  = c.FormValue("name")
		total = c.FormValue("total")
	)
	db := db.ConnectDB()
	ConvID, err := strconv.Atoi(total)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	data := models.ShoppingList{Name: name, Total: ConvID}
	err = v.Struct(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	db.Create(&data)
	return c.JSON(http.StatusOK, data)
}

// UpdateShoppingList Function
func UpdateShoppingList(c echo.Context) error {
	var (
		v     = validator.New()
		ID    = c.Param("id")
		name  = c.FormValue("name")
		total = c.FormValue("total")
		db    = db.ConnectDB()
	)

	ConvID, err := strconv.Atoi(total)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "Total is missing")
	}
	data := models.ShoppingList{}
	db.Where("ID= ?", ID).First(&data)
	data.Name = name
	data.Total = ConvID
	err = v.Struct(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	db.Save(&data)

	return c.JSON(http.StatusOK, data)
}

// DeleteShoppingList Function
func DeleteShoppingList(c echo.Context) error {
	var (
		ID = c.Param("id")
		db = db.ConnectDB()
	)

	data := models.ShoppingList{}
	err := db.Where("id= ?", ID).Delete(&data)
	if err == nil {
		return c.String(http.StatusNotFound, "Data Not Found")
	} else {
		return c.String(http.StatusOK, "Delete Success ID: "+ID)
	}

}
