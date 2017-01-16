package models

// Member represents a member record.
type Member struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
