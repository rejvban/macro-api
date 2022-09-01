package intake

import (
	"github.com/lucsky/cuid"
	"gorm.io/gorm"
)

type Intake struct {
	ID   string `gorm:"primary_key" json:"id"`
	Meal string `json:"meal"`
}

func (intake Intake) BeforeCreate(tx *gorm.DB) error {
	id := cuid.New()

	tx.Statement.SetColumn("ID", id)

	return nil
}
