package models

import (
	"time"
)

type User struct {
	Email             string     `db:"email"`
	Name              string     `db:"name"`
	CPF               string     `db:"cpf"`
	CNPJ              string     `db:"cnpj"`
	StateRegistration string     `db:"state_registration"`
	ContactName       string     `db:"contact_name"`
	MobileNumber      string     `db:"mobile_number"`
	Password          string     `db:"password"`
	Groups            UserGroups `db:"group"`
	Status            uint       `db:"status"`
	CreatedAt         time.Time  `db:"created_at"`
	ChangedAt         time.Time  `db:"changed_at"`
	IsDeleted         bool       `db:"is_deleted"`
}

type UserGroups uint8

const (
	GroupRoot UserGroups = 1 << iota
	GroupAdmin
	GorupSeller
)

// Set.
func (g *UserGroups) Set(flag UserGroups) {
	// r := *g | flag
	// g = &r
	*g = *g | flag
}

// Clear.
func (g *UserGroups) Clear(flag UserGroups) {
	*g = *g &^ flag
}

// Toggle.
func (g *UserGroups) Toggle(flag UserGroups) {
	*g = *g ^ flag
}

// Has.
func (g UserGroups) Has(flag UserGroups) bool {
	return (g & flag) != 0
}

// Has.
func (g UserGroups) Hav() bool {
	// *g & flag
	return false
}

// Is Admin.
func (g UserGroups) IsAdmin() bool {
	return g&GroupAdmin != 0
}
