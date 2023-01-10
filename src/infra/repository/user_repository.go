package repository

import (
	"github.com/google/uuid"
	"github.com/online.scheduling-api/src/business/models"
)

func GetAllUsers() (*[]models.User, error) {
	return nil, nil
}

func GetUserById(uuid *uuid.UUID) (*models.User, error) {
	return nil, nil
}

func CreateNewUser(u *models.User) error {
	return nil
}

func UpdateUser(u *models.User) error {
	return nil
}

func DeleteUserById(uuid *uuid.UUID) error {
	return nil
}

func ExistsByPhone(phone string) (bool, error) {
	return false, nil
}
