package domain

import "strings"

type Product struct {
	ID          uint
	UUID        string
	SKU         string
	Name        string
	Description string
}

func NewProductName(response string) (string, error) {
	return strings.TrimSpace(response), nil
}
