package models

import (
	"time"
)

type Project struct {
	ID               int64     `json:"id"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	Users            []*User   `json:"users"`
	CreationDate     time.Time `json:"creation_date"`
	ModificationDate time.Time `json:"modification_date"`
}
