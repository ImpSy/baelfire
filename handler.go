package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type handler struct {
	db *db
}

func (h handler) createTarget(c echo.Context) error {
	t := new(target)
	if err := c.Bind(t); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	if err := h.db.createTarget(t); err != nil {
		return c.JSON(http.StatusConflict, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, t)
}

func (h handler) listTargets(c echo.Context) error {
	names := h.db.listTarget()
	return c.JSON(http.StatusOK, names)
}

func (h handler) getTarget(c echo.Context) error {
	name := c.Param("name")
	t, err := h.db.getTarget(name)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, t)
}

func (h handler) deleteTarget(c echo.Context) error {
	name := c.Param("name")
	h.db.deleteTarget(name)
	return c.NoContent(http.StatusOK)
}

func (h handler) getTargetVersion(c echo.Context) error {
	name := c.Param("name")

	t, err := h.db.getTarget(name)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}

	version, err := t.getVersion()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"target":  t.Name,
		"version": version,
	})
}
