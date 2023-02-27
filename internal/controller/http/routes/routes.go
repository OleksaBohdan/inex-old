package routes

import (
	"github.com/labstack/echo/v4"
	"inex/main/internal/controller/http/handlers"
	"inex/main/internal/repository"
)

func RegisterRoutes(e *echo.Echo, repo *repository.InexRepo) {
	handler := handlers.NewHandler(*repo)

	e.POST("/api/v1/note", handler.CreateNote)
}
