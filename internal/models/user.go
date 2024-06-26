package models

import (
	"time"
)

type User struct {
	ID               int64     `json:"id"`
	Email            string    `json:"email"`
	Password         string    `json:"password"`
	FirstName        string    `json:"first_name"`
	LastName         string    `json:"last_name"`
	CreationDate     time.Time `json:"creation_date"`
	ModificationDate time.Time `json:"modification_date"`
}
