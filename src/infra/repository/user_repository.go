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
	UpdateUser(uuid *uuid.UUID, u *models.User) (isFound bool, err error)
	DeleteUserById(uuid *uuid.UUID) (isFound bool, err error)
	ExistsByPhone(phone string) (bool, error)
}

type UserRepository struct {
	Client *mongo.Client
}

func (ur *UserRepository) GetAllUsers() ([]*models.User, error) {
	users := []*models.User{}

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
	var user *models.User

	filter := &bson.M{"id": uuid}

	err := ur.collection().FindOne(context.TODO(), filter).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	return user, err
}

func (ur *UserRepository) CreateNewUser(u *models.User) error {
	_, err := ur.collection().InsertOne(context.TODO(), u)

	return err
}

func (ur *UserRepository) UpdateUser(uuid *uuid.UUID, u *models.User) (isFound bool, err error) {
	// err = ur.gocollection().UpdateId(uuid, u)

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
	filter := &bson.M{"id": uuid}
	res, err := ur.collection().DeleteOne(context.TODO(), filter)

	if err != nil {
		return false, err
	}

	if res.DeletedCount <= 0 {
		return false, nil
	}

	return true, nil
}

func (ur *UserRepository) ExistsByPhone(phone string) (bool, error) {
	filter := &bson.M{"phone": phone}

	result := ur.collection().FindOne(context.TODO(), filter)
	err := result.Err()

	if err == mongo.ErrNoDocuments {
		return false, nil
	}

	return true, err
}

func (ur *UserRepository) collection() *mongo.Collection {
	return ur.Client.Database(config.GetDBName()).Collection(config.GetUsersCollection())
}
