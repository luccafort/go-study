package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "Hello, Gopher World!")
	})

	// /users/
	e.POST("/users", createUser)
	e.GET("/users/:id", getUserByID)
	e.PUT("/users/:id", updateUserByID)
	e.DELETE("/users/:id", deleteUserByID)

	e.Logger.Fatal(e.Start(":1323"))
}

func createUser(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	fmt.Printf("name: %v, email: %v", name, email)
	return c.String(http.StatusOK, fmt.Sprintf("name: %v, email: %v", name, email))
}

func getUserByID(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, fmt.Sprintln(id))
}

func updateUserByID(c echo.Context) error {
	return nil
}

func deleteUserByID(c echo.Context) error {
	return nil
}
