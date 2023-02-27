package handlers

import (
	"github.com/labstack/echo/v4"
	"inex/main/domain"
	"inex/main/internal/repository"
	"inex/main/internal/usecase"
	"net/http"
)

type Handler struct {
	Repo repository.InexRepo
}

func NewHandler(repo repository.InexRepo) *Handler {
	return &Handler{repo}
}

func (h Handler) CreateNote(c echo.Context) error {
	var note domain.Note

	err := c.Bind(&note)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid data"})
	}

	err = usecase.CreateNote(note, h.Repo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Note created successfully"})
}
