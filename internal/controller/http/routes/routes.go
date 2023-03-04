package routes

import (
	"github.com/labstack/echo/v4"
	"inex/main/internal/controller/http/handlers"
	"inex/main/internal/repository"
)

func RegisterRoutes(e *echo.Echo, repo *repository.InexRepo) {
	handler := handlers.NewHandler(*repo)

	e.POST("/api/v1/note", handler.CreateNote)

	e.POST("/api/v1/income-item", handler.CreateIncomeItem)
	e.GET("/api/v1/income-item", handler.ReadIncomeItems)
	e.PUT("/api/v1/income-item", handler.UpdateIncomeItem)
}
