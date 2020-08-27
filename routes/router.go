package routes

import (
	"net/http"
	"strail/api"

	"github.com/labstack/echo"
)

// Routes
func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome To Strail API")
	})

	// User Endpoint
	e.GET("/user", api.GetUser)
	e.GET("/user/get/:id", api.GetUserByID)
	e.POST("/user", api.CreateUser)
	e.PUT("/user/:id", api.UpdateUser)
	e.DELETE("/user/del/:id", api.DeleteUser)

	// Shopping List Endpoint
	e.GET("/shopping", api.GetShoppingList)
	e.GET("/shopping:id/", api.GetShoppingListByID)
	e.POST("/shopping", api.CreateShoppingList)
	e.PUT("/shopping/:id", api.UpdateShoppingList)
	e.DELETE("/shopping/delete/:id", api.DeleteShoppingList)

	return e

}
