package models

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Todo struct {
	ID          string    `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name        string    `json:"name"`
	Completed   bool      `json:"completed"`
	Due         time.Time `json:"due"`
	Description string    `json:"description"`
}

func (m *Todo) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4().String())
	return nil
}
