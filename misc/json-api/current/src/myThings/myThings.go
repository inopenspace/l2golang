// Package things defines the things that the application will use.  This
// package primarily defines struct types and their JSON serialization rules.
package myThings

// Characters that players will create.
type Character struct {
	ID    int    `json: "id"`
	Name  string `json: "name"`
	Race  string `json: "race"`
	Level int    `json: "level"`
}

// Spells which can be cast. Currently they are not tied to characters
type Spell struct {
	ID          int    `json: "id"`
	Level       int    `json: "level"`
	Name        string `json: "name"`
	Description string `json: "description"`
}