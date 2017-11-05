package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func main(){
	fmt.Println("welcome to the server")
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "yallo from the web side!")
	})
	e.Start(":8000")
}