package models

type Collection struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Index   string `json:"index"`
	Project int64  `json:"project"`
}
