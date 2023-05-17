package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func getCats(c echo.Context) error {
	//fmt.Println("getCats is firing...")
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")
	// catName := "harry"
	// catType := "fluffy"
	return c.String(http.StatusOK, fmt.Sprintf("your cat name is %s\nand his type is: %s\n", catName, catType))
}

func getCats2(c echo.Context) error {
	// fmt.Println("getCats is firing...")
	catName := c.Param("name")
	// catType := c.Param("type")
	// catName := "harry"
	catType := "fluffy"
	return c.String(http.StatusOK, fmt.Sprintf("your cat name is %s\nand his type is: %s\n", catName, catType))
}

func GetData22(c echo.Context) error {
	username := c.Param("username")
	password := c.Param("password")
	return c.String(http.StatusOK, fmt.Sprintf("your username is %s\nand password is %s\n", username, password))
}

func main() {
	e := echo.New()
	e.GET("/", hello)
	e.GET("/cats", getCats)
	e.GET("/cats2/:name", getCats2)
	// e.GET("/v1/pwlu/:username/:password", handler2.GetData22)
	e.GET("/v1/pwlu/:username/:password", GetData22)
	e.Logger.Fatal(e.Start(":1323"))

}

// e.POST("/users", saveUser)
// e.GET("/users/:id", getUSer)
// e.PUT("/users/:id", updateUser)
// e.DELETE("/users/:id", deleteUser)

// func getUser(c echo.Context) error {
// 	//user ID from path `users/:id`
// 	id := c.Param("id")
// 	return c.String(http.StatusOK, id)
// }

//FETCH A FILE USING curl AND DISPLAY ITS CONTENT
// $ curl https://www.digitalocean.com/robots.txt
// $ curl www.google.com/

//FETCH A FILE USING curl AND SAVE TO THE SAME NAME AS DOWNLOADED
//$ curl -O https://www.digitalocean.com/robots.txt

// CALL THE NEW SAVED FILE
//$ cat robots.txt

// FETCH AND DOWNLOAD A NEW FILE TO A NEW LOCALLY NAMED FILE
//$ curl -o do-bots.txt https://www.digitalocean.com/robots.txt

//CALL THE NEW SAVED FILE
//$ cat do-bots.txt

//FETCH AND DISPLAY THE REQUEST HEADERS RATHER THAN THE CONTENTS OF THE FILE
//$ curl -I www.digitalocean.com/robots.txt

// FETCH AND FOLLOW REDIRECTS
//$ curl -L www.digitalocean.com/robots.txt

// FETCH AND FOLLOW REDIRECTS AND DOWNLOAD CONTENT TO A NEW LOCALLY NAMED FILE
//$ curl -L -o do-bots.txt www.digitalocean.com/robots.txt
