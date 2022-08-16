package models

import (
	"time"
)

type Blog struct {
        ID                  uint              `json:"id" gorm:"primary_key"`
        Title               string            `json:"title"`
        Body                string            `json:"body"`
        Slug                string            `json:"slug"`
        CreatedAt           time.Time         `json:"created_at"`
        UpdatedAt           time.Time         `json:"updated_at"`
}
