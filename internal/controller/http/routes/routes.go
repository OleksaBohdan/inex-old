package routes

import (
	"github.com/labstack/echo/v4"
	"inex/main/internal/controller/http/handlers"
	"inex/main/pkg/postgres"
)

func RegisterRoutes(e *echo.Echo, pg *postgres.Postgres) {
	handler := handlers.NewHandler(pg)

	e.POST("/api/v1/note", handler.CreateNote)
}
