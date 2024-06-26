package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/victorgomez09/vira-dapo/internal/database"
	"github.com/victorgomez09/vira-dapo/internal/models"
)

type CollectionHandler struct {
	Repository *database.DB
}

func (handler *CollectionHandler) Create(c echo.Context) error {
	var collection models.Collection
	err := c.Bind(collection)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Can't deserialize data",
		})
	}

	// TODO: set schema before create table
	var project models.Project
	row := handler.Repository.Db.QueryRowContext(context.TODO(), "SELECT * FROM projects WHERE id = $1", collection.Project)
	err = row.Scan(&project)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"error": "Project not found",
		})
	}
	_, err = handler.Repository.Tx.Exec("SET search_path=$1", project.Name)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"error": "Schema not found",
		})
	}

	_, err = handler.Repository.Db.ExecContext(context.TODO(), "INSERT INTO collections(name, index) VALUES($1, $2)", collection.Name, collection.Index)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Error creating collection",
		})
	}

	row = handler.Repository.Db.QueryRowContext(context.TODO(), "SELECT * FROM collections WHERE name = $1", collection.Name)
	err = row.Scan(collection)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Internal server error",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": collection,
	})
}
