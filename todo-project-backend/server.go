package main

import (
	"net/http"
	db "todo-project-backend/db"
	u "todo-project-backend/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db.DB()
	db.Migrate()

	e := echo.New()

	// implements CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, `{"response": "success"}`)
	})

	e.GET("/login/:email", u.GetUserLogin) // api for login page

	e.POST("/login/:email", u.CreateUser) // api for login page

	e.GET("/todos/:email", u.GetUserTodos) // api for todo page

	e.POST("/todos/:email", u.Save) // api for todo page (save=update)

	// may not use in project
	e.PUT("/users/:id", u.Update)
	e.DELETE("/users/:id", u.Delete)

	e.Logger.Fatal(e.Start(":1234")) // create port
}
