package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/victorgomez09/vira-dapo/internal/database"
	"github.com/victorgomez09/vira-dapo/internal/models"
	"github.com/victorgomez09/vira-dapo/internal/utils"
)

type UserHandler struct {
	Repository *database.DB
}

func (handler *UserHandler) Createuser(c echo.Context) error {
	var user models.User

	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Error parsing new user",
		})
	}

	hashedPass, err := utils.HashPassword(user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Error hashing user password",
		})
	}

	user.Password = hashedPass
	_, err = handler.Repository.Db.ExecContext(context.TODO(), "INSERT INTO users(email, password, first_name, last_name, creation_date) VALUES($1, $2, $3, $4, $5)", user.Email, user.Password, user.FirstName, user.LastName, time.Now())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Error inserting new user",
		})
	}

	row := handler.Repository.Db.QueryRowContext(context.TODO(), "SELECT * FROM users WHERE email = $1", user.Email)
	err = row.Scan(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Error getting new user",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"user": user,
	})
}
