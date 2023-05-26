// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Intent struct {
	UserID *string          `json:"user_id,omitempty"`
	Items  []*IntentProduct `json:"items,omitempty"`
}

type IntentProduct struct {
	ID          string   `json:"id"`
	Title       *string  `json:"title,omitempty"`
	Price       *float64 `json:"price,omitempty"`
	Category    *string  `json:"category,omitempty"`
	Description *string  `json:"description,omitempty"`
	Image       *string  `json:"image,omitempty"`
}

type NewIntent struct {
	UserID *string             `json:"user_id,omitempty"`
	Items  []*NewIntentProduct `json:"items,omitempty"`
}

type NewIntentProduct struct {
	Title       *string  `json:"title,omitempty"`
	Price       *float64 `json:"price,omitempty"`
	Category    *string  `json:"category,omitempty"`
	Description *string  `json:"description,omitempty"`
	Image       *string  `json:"image,omitempty"`
}