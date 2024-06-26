package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/victorgomez09/vira-dapo/api/dtos"
	"github.com/victorgomez09/vira-dapo/internal/auth"
	"github.com/victorgomez09/vira-dapo/internal/database"
	"github.com/victorgomez09/vira-dapo/internal/models"
	"github.com/victorgomez09/vira-dapo/internal/utils"
)

type AuthHandler struct {
	Repository *database.DB
}

func (handler *AuthHandler) Login(c echo.Context) error {
	var dto dtos.LoginDto
	err := c.Bind(dto)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Can't deserialize data")
	}

	var user models.User
	row := handler.Repository.Db.QueryRowContext(c.Request().Context(), "SELECT * FROM users WHERE email = $1", dto.Email)
	errSelect := row.Scan(&user)
	if errSelect != nil {
		return c.JSON(http.StatusNotFound, "User not found")
	}

	if !utils.CheckPasswordHash(dto.Password, user.Password) {
		return c.JSON(http.StatusUnauthorized, "Password not match")
	}

	token, err := auth.CreateToken(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Can't create token")
	}

	return c.JSON(http.StatusUnauthorized, echo.Map{
		"token": token,
	})
}
