package handlers

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"inex/main/domain"
	"inex/main/internal/repository"
	"inex/main/pkg/postgres"
	"net/http"
)

type Handler struct {
	Pg *postgres.Postgres
}

func NewHandler(pg *postgres.Postgres) *Handler {
	return &Handler{pg}
}

func (h Handler) CreateNote(c echo.Context) error {
	var note domain.Note

	err := c.Bind(&note)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid data"})
	}

	inex := repository.New(h.Pg)

	err = inex.CreateNote(context.Background(), note)
	if err != nil {
		return fmt.Errorf("seeder - seedNotes - inex.CreateNote: %w", err)
	}

	fmt.Println(note)
	return c.JSON(http.StatusCreated, map[string]string{"message": "Note created successfully"})
}
