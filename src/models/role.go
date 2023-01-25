package models

type Role int16

const (
	Admin    Role = 1
	Worker   Role = 2
	Customer Role = 3
)

func (r Role) GetDescription() string {
	switch r {
	case Admin:
		return "Administrador(a)"
	case Worker:
		return "Funcionário(a)"
	case Customer:
		return "Cliente"
	default:
		return ""
	}
}
