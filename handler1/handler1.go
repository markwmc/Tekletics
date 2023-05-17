package handler1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// e.POST("/users", saveUser)
// e.GET("/users/:id", getUSer)
// e.PUT("/users/:id", updateUser)
// e.DELETE("/users/:id", deleteUser)

// e.GET("/users/:id", getUSer)
func getUser(c echo.Context) error {
	//user ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

// e.GET("/users/:id", getUSer)
func GetUser(c echo.Context) error {
	//user ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func updateUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
