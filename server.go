package main

import (
	_"fmt"
	"net/http"
	"github.com/labstack/echo"
)

type Msg struct {
	Message string `json:"message"`
}

func GreetingsWithParams(c echo.Context) error {
	params := c.Param("name")
	return c.JSON(http.StatusOK, Msg{
		Message: "Sei laaaaaaa " + params,
	})
}

func GreetingsWithQuery(c echo.Context) error {
	query := c.QueryParam("name")
	return c.JSON(http.StatusOK, Msg{
		Message: "Hakuna Matata " + query,
	})
}

func Greetings(c echo.Context) error {
	return c.JSON(http.StatusOK, Msg{
		Message: "Hello World",
	})
}

func main (){
	e := echo.New()
	e.GET("/message", Greetings)
	e.GET("/message/:name", GreetingsWithParams)
	e.GET("/boladona-query", GreetingsWithQuery)
	e.Logger.Fatal(e.Start(":3000"))
}







