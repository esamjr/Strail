package api

import (
	"net/http"
	"strail/db"
	"strail/models"
	"time"

	"github.com/labstack/echo"
)

// Get Schedule
func GetSchedule(c echo.Context) error {
	// Param := c.Param("username")
	db := db.ConnectDB()
	data := []models.Schedule{}
	// db.Where("username= ?", Param).Find(&data)
	db.Find(&data)
	if data == nil {
		return c.JSON(http.StatusNotFound, "Data Not Found")
	} else {
		return c.JSON(http.StatusOK, data)
	}
}

// Get ScheduleByID
func GetScheduleByID(c echo.Context) error {
	ID := c.Param("id")
	db := db.ConnectDB()
	data := []models.Schedule{}
	db.Where("id= ?", ID).Find(&data)

	if data == nil {
		return c.JSON(http.StatusNotFound, "Data Not Found")
	}

	return c.JSON(http.StatusOK, data)
}

// Set Schedule
func CreateSchedule(c echo.Context) error {
	var (
		// param = c.Param("userid")
		e  = new(models.Schedule)
		db = db.ConnectDB()
	)
	err := c.Bind(&e)
	if err != nil {
		return err
	}
	// db.Where("username= ?", param).Find(&e)
	db.Create(&e)
	return c.JSON(http.StatusCreated, e)

}

// Update Schedule
func UpdateSchedule(c echo.Context) error {
	var (
		ID    = c.Param("id")
		title = c.FormValue("title")
		times = c.FormValue("time")
		db    = db.ConnectDB()
	)
	Convert, _ := time.Parse(time.RFC3339Nano, times)
	data := models.Schedule{}
	db.Where("ID= ?", ID).First(&data)
	data.Title = title
	data.Time = Convert
	db.Save(&data)
	return c.JSON(http.StatusOK, data)
}

func DeleteSchedule(c echo.Context) error {
	var (
		ID = c.Param("id")
		db = db.ConnectDB()
	)

	data := models.Schedule{}
	err := db.Where("id= ?", ID).Delete(&data)
	if err == nil {
		return c.String(http.StatusNotFound, "Data Not Found")
	} else {
		return c.String(http.StatusOK, "Delete Success ID: "+ID)
	}

}
