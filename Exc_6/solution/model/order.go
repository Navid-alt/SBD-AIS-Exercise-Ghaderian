package model

import (
	"fmt"
)

const (
	orderFilename = "order_%d.md"

	// todo create markdown emplate, fields should be able to be populated with fmt.Sprintf
	markdownTemplate = `
# order: %d
| Created at      | Drink ID | Amount |
|-----------------|----------|--------|
| %-15s | %-8d | %-6d |
Thanks for drinking with us!
`
)

type Order struct {
	Base
	Amount uint64 `json:"amount"`
	// Relationships
	// foreign key
	DrinkID uint  `json:"drink_id" gorm:"not null"`
	Drink   Drink `json:"drink"`
}

func (o *Order) ToMarkdown() string {
	formattedTime := o.CreatedAt.Format("Jan 02 15:04:05")
	return fmt.Sprintf(markdownTemplate, o.ID, formattedTime, o.DrinkID, o.Amount)
}

func (o *Order) GetFilename() string {
	return fmt.Sprintf(orderFilename, o.ID)
}
