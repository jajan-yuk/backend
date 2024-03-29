package entity

import "fmt"

type Product struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Price       int64     `json:"price"`
	Category    Category  `json:"category"`
	Variant     []Variant `json:"variant"`
}

func (p *Product) Validate() error {
	if p.Name == "" {
		return fmt.Errorf("product name is required")
	}

	if p.Description == "" {
		return fmt.Errorf("product description is required")
	}

	if p.Image == "" {
		return fmt.Errorf("product image is required")
	}

	if p.Price == 0 {
		return fmt.Errorf("product price is required")
	}

	if p.Category.ID == 0 {
		return fmt.Errorf("product category is required")
	}

	return nil
}

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (c *Category) Validate() error {
	if c.Name == "" {
		return fmt.Errorf("category name is required")
	}

	return nil
}

type Variant struct {
	ID               int64         `json:"id"`
	Name             string        `json:"name"`
	TotalPrice       int64         `json:"total_price"`
	CountVariantType int64         `json:"count_variant_type"`
	VariantTypes     []VariantType `json:"variant_types,omitempty"`
}

type VariantType struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Price     int64  `json:"price"`
	VariantID int64  `json:"variant_id"`
}
