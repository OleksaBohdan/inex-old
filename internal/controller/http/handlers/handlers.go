package handlers

import (
	"fmt"
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

	n, err := usecase.CreateNote(note, h.Repo)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
	}

	return c.JSON(http.StatusCreated, n)
}

func (h Handler) CreateIncomeItem(c echo.Context) error {
	var incomeItem domain.IncomeItem

	err := c.Bind(&incomeItem)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid data"})
	}

	item, err := usecase.CreateIncomeItem(incomeItem, h.Repo)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
	}

	return c.JSON(http.StatusCreated, item)
}

func (h Handler) UpdateIncomeItem(c echo.Context) error {
	var incomeItem domain.IncomeItem

	err := c.Bind(&incomeItem)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid data"})
	}

	item, err := usecase.UpdateIncomeItem(incomeItem, h.Repo)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
	}

	return c.JSON(http.StatusCreated, item)
}

func (h Handler) ReadIncomeItems(c echo.Context) error {
	var user domain.User

	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid data"})
	}

	items, err := usecase.ReadIncomeItems(user, h.Repo)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
	}

	return c.JSON(http.StatusCreated, items)
}
