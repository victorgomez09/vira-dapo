package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/victorgomez09/vira-dapo/internal/database"
	"github.com/victorgomez09/vira-dapo/internal/models"
)

type ProjectHandler struct {
	Repository *database.DB
}

func (handler *ProjectHandler) FindAll(c echo.Context) error {
	var projects []models.Project

	rows, err := handler.Repository.Db.QueryContext(context.TODO(), "SELECT * FROM projects")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Error getting projects",
		})
	}

	for rows.Next() {
		var project models.Project
		err = rows.Scan(&project)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error": "Error serializing project",
			})
		}

		projects = append(projects, project)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": projects,
	})
}

func (handler *ProjectHandler) FindById(c echo.Context) error {
	id := c.Param("id")

	var project models.Project
	row := handler.Repository.Db.QueryRowContext(context.TODO(), "SELECT * FROM projects WHERE id = $1", id)
	err := row.Scan(&project)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Error getting projects",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": project,
	})
}

func (handler *ProjectHandler) Create(c echo.Context) error {
	var project models.Project
	err := c.Bind(&project)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Can't deserialize data",
		})
	}

	_, err = handler.Repository.Db.ExecContext(context.TODO(), "INSERT INTO projects(name, description, creation_date) VALUES($1, $2, $3)", project.Name, project.Description, time.Now())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Error creating project",
		})
	}

	row := handler.Repository.Db.QueryRowContext(context.TODO(), "SELECT * FROM projects WHERE name = $1", project.Name)
	err = row.Scan(&project)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"error": "Error getting new project",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": project,
	})
}
