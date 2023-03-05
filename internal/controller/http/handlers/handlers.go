package handlers

import (
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
	"inex/main/domain"
	"inex/main/internal/repository"
	"inex/main/internal/usecase"
	"net/http"
	"sync"
)

var (
	lock = sync.Mutex{}
)

type Handler struct {
	Repo repository.InexRepo
}

func NewHandler(repo repository.InexRepo) *Handler {
	return &Handler{repo}
}

func (h Handler) CreateNote(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	var note domain.Note

	err := c.Bind(&note)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	n, err := usecase.CreateNote(note, h.Repo)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusCreated, n)
}

func (h Handler) CreateIncomeItem(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	var incomeItem domain.IncomeItem

	err := c.Bind(&incomeItem)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	item, err := usecase.CreateIncomeItem(incomeItem, h.Repo)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusCreated, item)
}

func (h Handler) ReadIncomeItems(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	var user domain.User

	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	user.ID = id

	items, err := usecase.ReadIncomeItems(user, h.Repo)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, items)
}

func (h Handler) UpdateIncomeItem(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	var incomeItem domain.IncomeItem

	err := c.Bind(&incomeItem)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	item, err := usecase.UpdateIncomeItem(incomeItem, h.Repo)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, item)
}

func (h Handler) DeleteIncomeItem(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	var incomeItem domain.IncomeItem

	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	incomeItem.ID = id

	err = usecase.DeleteIncomeItem(incomeItem, h.Repo)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h Handler) CreateCostItem(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	var costItem domain.CostItem

	err := c.Bind(&costItem)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	item, err := usecase.CreateCostItem(costItem, h.Repo)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusCreated, item)
}

func (h Handler) ReadCostItems(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	var user domain.User

	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	user.ID = id

	items, err := usecase.ReadCostItems(user, h.Repo)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, items)
}

func (h Handler) UpdateCostItem(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	var costItem domain.CostItem

	err := c.Bind(&costItem)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	item, err := usecase.UpdateCostItem(costItem, h.Repo)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, item)
}

func (h Handler) DeleteCostItem(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	var costItem domain.CostItem

	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	costItem.ID = id

	err = usecase.DeleteCostItem(costItem, h.Repo)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h Handler) CreateIncome(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	var income domain.Income

	err := c.Bind(&income)
	if err != nil {
		fmt.Println(err)
		return c.NoContent(http.StatusBadRequest)
	}

	i, err := usecase.CreateIncome(income, h.Repo)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusCreated, i)
}
