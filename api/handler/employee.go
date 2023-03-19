package handler

import (
	"golang-employee-crud-example/model"
	"golang-employee-crud-example/repository"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type EmployeeHandler struct {
	repo repository.EmployeeRepository
}

func NewEmployeeHandler(repo repository.EmployeeRepository) *EmployeeHandler {
	return &EmployeeHandler{repo: repo}
}

func (h *EmployeeHandler) CreateEmployee(c echo.Context) error {
	var e model.Employee
	err := c.Bind(&e)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	e.ID = uuid.New()
	id, err := h.repo.Create(c.Request().Context(), &e)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	e.ID = id
	return c.JSON(http.StatusCreated, e)
}

func (h *EmployeeHandler) GetEmployees(c echo.Context) error {
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")

	employees, err := h.repo.GetAll(c.Request().Context(), limit, offset)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, employees)
}

func (h *EmployeeHandler) GetEmployeeByID(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	e, err := h.repo.GetByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	if e == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "employee not found"})
	}

	return c.JSON(http.StatusOK, e)
}

func (h *EmployeeHandler) UpdateEmployee(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	var e model.Employee
	err = c.Bind(&e)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	e.ID = id

	err = h.repo.Update(c.Request().Context(), &e)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *EmployeeHandler) DeleteEmployee(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	err = h.repo.Delete(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}
