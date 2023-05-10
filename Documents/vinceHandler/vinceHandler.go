package vinceHandler

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

// make this work
// e.GET("/GetData22/username/password", save)
func GetData22(c echo.Context) error {
	quote := new(Quote)

	username := c.Param("username")
	quote.Username = username
	password := c.Param("password")
	quote.Password = password

	log.Println("in GetData22", quote)

	return c.JSON(http.StatusOK, quote)
}
