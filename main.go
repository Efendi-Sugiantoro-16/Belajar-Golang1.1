package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func getUser(c echo.Context) error {
	user := User{
		ID:    1,
		Name:  "Efendi Sugiantoro",
		Email: "efendi16@bmg.com",
	}
	return c.JSON(http.StatusOK, user)
}

func getUserById(c echo.Context) error {
	user := User{
		ID:    1,
		Name:  "Efendi Sugiantoro",
		Email: "efendi16@bmg.com",
	}
	return c.JSON(http.StatusOK, user)
}

func main() {
	e := echo.New()
	e.GET("/user", getUser)
	e.Logger.Fatal(e.Start(":8080"))
}
