package models

import (
	"github.com/google/uuid"
)

type User struct {
	Id         uuid.UUID  `bson:"id"`
	Name       string     `bson:"name"`
	Phone      string     `bson:"phone"`
	Role       Role       `bson:"role"`
	IsActive   bool       `bson:"isActive"`
	Modalities []Modality `bson:"modalities"`
	Login      string     `bson:"login"`
	Passphrase string     `bson:"passphrase"`
}

func MapNewUserFrom(name, login, passphrase, phone string, role Role, active bool) User {
	return User{
		Id:         uuid.New(),
		Name:       name,
		Login:      login,
		Passphrase: passphrase,
		Phone:      phone,
		Role:       role,
		IsActive:   active,
	}
}

func MapUserFrom(name, phone string, role Role, active bool) User {
	return User{
		Id:       uuid.New(),
		Name:     name,
		Phone:    phone,
		Role:     role,
		IsActive: active,
	}
}
