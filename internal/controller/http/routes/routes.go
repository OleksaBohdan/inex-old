package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"inex/main/internal/controller/http/handlers"
	"inex/main/internal/repository"
)

func RegisterRoutes(e *echo.Echo, repo *repository.InexRepo) {
	handler := handlers.NewHandler(*repo)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/api/v1/note", handler.CreateNote)

	e.POST("/api/v1/income-item", handler.CreateIncomeItem)
	e.GET("/api/v1/income-item/:id", handler.ReadIncomeItems)
	e.PUT("/api/v1/income-item", handler.UpdateIncomeItem)
	e.DELETE("/api/v1/income-item/:id", handler.DeleteIncomeItem)

	e.POST("/api/v1/cost-item", handler.CreateCostItem)
	e.GET("/api/v1/cost-item/:id", handler.ReadCostItems)
	e.PUT("/api/v1/cost-item", handler.UpdateCostItem)
	e.DELETE("/api/v1/cost-item/:id", handler.DeleteCostItem)

}
