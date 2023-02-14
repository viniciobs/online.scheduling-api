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
	Get(*models.UserFilter) ([]*models.User, error)
	GetUserById(uuid *uuid.UUID) (*models.User, error)
	CreateNewUser(u *models.User) error
	ActivateUser(uuid *uuid.UUID) error
	EditUser(uuid *uuid.UUID, u *models.User) error
	EditUserModalities(uuid *uuid.UUID, u *models.User) error
	DeleteUserById(uuid *uuid.UUID) (isFound bool, err error)
	ExistsBy(uuid *uuid.UUID, phone, login *string) (bool, error)
	Authenticate(login, passphrase string) (bool, *models.User)
	EditAuth(uuid *uuid.UUID, login, passphrase string) error
}

type UserRepository struct {
	Client *mongo.Client
}

func (ur *UserRepository) Get(filter *models.UserFilter) ([]*models.User, error) {
	query := bson.M{}

	if filter.UserId != uuid.Nil {
		query["id"] = filter.UserId
	}

	if filter.UserName != "" {
		query["name"] = bson.M{"$regex": filter.UserName, "$options": "i"}
	}

	if filter.ModalityId != uuid.Nil {
		query["modalities.id"] = filter.ModalityId
	}

	if filter.ModalityName != "" {
		query["modalities.name"] = bson.M{"$regex": filter.ModalityName, "$options": "i"}
	}

	cursor, err := ur.collection().Find(context.TODO(), query)
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

func (ur *UserRepository) EditUserModalities(uuid *uuid.UUID, u *models.User) error {
	filter := bson.M{"id": uuid}
	update := bson.M{
		"$set": bson.M{
			"modalities": u.Modalities,
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

func (ur *UserRepository) ExistsBy(uuid *uuid.UUID, phone, login *string) (bool, error) {
	err := ur.collection().
		FindOne(
			context.TODO(),
			bson.M{
				"id": bson.M{"$ne": uuid},
				"$or": []bson.M{
					{"phone": phone},
					{"login": login},
				},
			}).
		Err()

	if err == mongo.ErrNoDocuments {
		return false, nil
	}

	return err == nil, err
}

func (ur *UserRepository) Authenticate(login, passphrase string) (bool, *models.User) {
	var user *models.User

	err := ur.collection().
		FindOne(
			context.TODO(),
			&bson.M{
				"login":      login,
				"passphrase": passphrase,
			}).Decode(&user)

	return err == nil, user
}

func (ur *UserRepository) EditAuth(uuid *uuid.UUID, login, passphrase string) error {
	filter := bson.M{"id": uuid}
	update := bson.M{
		"$set": bson.M{
			"login":      login,
			"passphrase": passphrase,
		},
	}

	res, err := ur.collection().UpdateOne(context.TODO(), filter, update)

	if res.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return err
}

func (ur *UserRepository) collection() *mongo.Collection {
	return ur.Client.
		Database(config.GetDBName()).
		Collection("USERS")
}
