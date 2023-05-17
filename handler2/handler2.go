package handler2

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// e.GET("/users2/:id", getUser2)
func GetUser2(c echo.Context) error {
	//user ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

// e.GET("/show", show)
func Show2(c echo.Context) error {
	//Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

// e.POST("/save", save)
func Save2(c echo.Context) error {
	//Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}

// JSON response
type Quote struct {
	Username string `json: "username"`
	Password string `json:"password"`
}

func GetData22(c echo.Context) error {
	quote := new(Quote)

	username := c.Param("username")
	quote.Username = username
	password := c.Param("password")
	quote.Password = password

	log.Println("in GetData22", quote)

	return c.JSON(http.StatusOK, quote)
}
