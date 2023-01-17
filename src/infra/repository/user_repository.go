package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/online.scheduling-api/config"
	"github.com/online.scheduling-api/src/business/models"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type IUserRepository interface {
	GetAllUsers() ([]*models.User, error)
	GetUserById(uuid *uuid.UUID) (*models.User, error)
	CreateNewUser(u *models.User) error
	UpdateUser(u *models.User) (isFound bool, err error)
	DeleteUserById(uuid *uuid.UUID) (isFound bool, err error)
	ExistsByPhone(phone string) (bool, error)
}

type UserRepository struct {
	Client *mongo.Client
}

func (ur *UserRepository) GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	cursor, err := ur.collection().Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var user models.User
		cursor.Decode(&user)

		users = append(users, &user)
	}

	return users, nil
}

func (ur *UserRepository) GetUserById(uuid *uuid.UUID) (*models.User, error) {
	// var user *models.User
	// // err := ur.collection().FindId(uuid).One(&user)

	// if err == mgo.ErrNotFound {
	// 	return nil, nil
	// }

	// return user, err
	return nil, nil
}

func (ur *UserRepository) CreateNewUser(u *models.User) error {
	_, err := ur.collection().InsertOne(context.TODO(), u)

	return err
}

func (ur *UserRepository) UpdateUser(u *models.User) (isFound bool, err error) {
	// err = ur.collection().UpdateId(u.Id, u)

	// if err == nil {
	// 	return true, nil
	// }

	// if err == mgo.ErrNotFound {
	// 	return false, nil
	// }

	// return false, mgo.ErrCursor

	return false, nil
}

func (ur *UserRepository) DeleteUserById(uuid *uuid.UUID) (isFound bool, err error) {
	// err = ur.collection().RemoveId(uuid)

	// if err == nil {
	// 	return true, nil
	// }

	// if err == mgo.ErrNotFound {
	// 	return false, nil
	// }

	// return false, err
	return false, nil
}

func (ur *UserRepository) ExistsByPhone(phone string) (bool, error) {
	return false, nil
}

func (ur *UserRepository) collection() *mongo.Collection {
	return ur.Client.Database(config.GetDBName()).Collection(config.GetUsersCollection())
}
