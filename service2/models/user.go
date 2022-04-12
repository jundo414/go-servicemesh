package models

type User struct {
	ID     string `json:"ID"`
	Name   string `json:"Name"`
	Gender string `json:"Gender"`
	Born   string `json:"Born"`
}
