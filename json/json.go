package json

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

// JSON response
type Quote struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Time     string `json:"time"`
}

// e.GET("/GetData22/:username/:password", GetData22)
func GetData22(c echo.Context) error {
	quote := new(Quote)

	username := c.Param("username")
	quote.Username = username
	password := c.Param("password")
	quote.Password = password
	quote.Time = time.Now().Format(time.RFC3339)

	log.Println("in GetData22", quote)

	return c.JSON(http.StatusOK, quote)
}

func main() {
	e := echo.New()
	e.GET("/GetData22/:username/:password", GetData22)
	e.Logger.Fatal(e.Start(":8080"))
}
