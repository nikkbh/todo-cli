package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Task   string
	Status status
}

type status string

const (
	Pending status = "pending"
	Done    status = "done"
)
