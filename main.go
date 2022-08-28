package main

import (
	"gomysql/controller"
	"gomysql/storage"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
  // Echo instance
  e := echo.New()
	storage.NewDB()
  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // Routes
  e.GET(("/users"), controller.GetUsers)
  e.GET(("/users/:id"), controller.GetUser)
  e.POST(("/users"), controller.CreateUser)
  e.PATCH(("/users/:id"), controller.UpdateUser)
  e.DELETE(("/users/:id"), controller.DeleteUser)


  // Start server
  e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}