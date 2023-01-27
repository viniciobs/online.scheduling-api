package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/online.scheduling-api/config"
	"github.com/online.scheduling-api/src/models"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type IUserRepository interface {
	GetAllUsers() ([]*models.User, error)
	GetUserById(uuid *uuid.UUID) (*models.User, error)
	CreateNewUser(u *models.User) error
	ActivateUser(uuid *uuid.UUID) error
	EditUser(uuid *uuid.UUID, u *models.User) error
	DeleteUserById(uuid *uuid.UUID) (isFound bool, err error)
	ExistsByPhone(uuid *uuid.UUID, phone *string) (bool, error)
}

type UserRepository struct {
	Client *mongo.Client
}

func (ur *UserRepository) GetAllUsers() ([]*models.User, error) {
	cursor, err := ur.collection().Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	users := []*models.User{}

	for cursor.Next(context.TODO()) {
		var user models.User
		cursor.Decode(&user)

		users = append(users, &user)
	}

	return users, nil
}

func (ur *UserRepository) GetUserById(uuid *uuid.UUID) (*models.User, error) {
	var user *models.User

	err := ur.collection().
		FindOne(
			context.TODO(),
			&bson.M{"id": uuid}).
		Decode(&user)

	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	return user, err
}

func (ur *UserRepository) CreateNewUser(u *models.User) error {
	_, err := ur.collection().InsertOne(context.TODO(), u)

	return err
}

func (ur *UserRepository) ActivateUser(uuid *uuid.UUID) error {
	filter := bson.M{"id": uuid}
	update := bson.M{
		"$set": bson.M{
			"isActive": true,
		},
	}

	res, err := ur.collection().UpdateOne(context.TODO(), filter, update)

	if res.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return err
}

func (ur *UserRepository) EditUser(uuid *uuid.UUID, u *models.User) error {
	filter := bson.M{"id": uuid}
	update := bson.M{
		"$set": bson.M{
			"name":  u.Name,
			"phone": u.Phone,
			"role":  u.Role,
		},
	}

	res, err := ur.collection().UpdateOne(context.TODO(), filter, update)

	if res.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return err
}

func (ur *UserRepository) DeleteUserById(uuid *uuid.UUID) (isFound bool, err error) {
	res, err := ur.collection().
		DeleteOne(
			context.TODO(),
			&bson.M{"id": uuid})

	if err != nil {
		return false, err
	}

	if res.DeletedCount <= 0 {
		return false, nil
	}

	return true, nil
}

func (ur *UserRepository) ExistsByPhone(uuid *uuid.UUID, phone *string) (bool, error) {
	err := ur.collection().
		FindOne(
			context.TODO(),
			&bson.M{
				"id":    &bson.M{"$ne": uuid},
				"phone": phone,
			}).
		Err()

	if err == mongo.ErrNoDocuments {
		return false, nil
	}

	return err == nil, err
}

func (ur *UserRepository) collection() *mongo.Collection {
	return ur.Client.
		Database(config.GetDBName()).
		Collection(config.GetUsersCollection())
}
