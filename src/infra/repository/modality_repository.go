package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/online.scheduling-api/config"
	"github.com/online.scheduling-api/src/models"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type IModalityRepository interface {
	GetAllModalities() ([]*models.Modality, error)
	GetModalityById(uuid *uuid.UUID) (*models.Modality, error)
	CreateNewModality(m *models.Modality) error
	EditModality(uuid *uuid.UUID, m *models.Modality) error
	DeleteSModalityById(uuid *uuid.UUID) (isFound bool, err error)
}

type ModalityRepository struct {
	Client *mongo.Client
}

func (mr *ModalityRepository) GetAllModalities() ([]*models.Modality, error) {
	cursor, err := mr.collection().Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	modalities := []*models.Modality{}

	for cursor.Next(context.TODO()) {
		var modality models.Modality
		cursor.Decode(&modality)

		modalities = append(modalities, &modality)
	}

	return modalities, nil
}

func (mr *ModalityRepository) GetModalityById(uuid *uuid.UUID) (*models.Modality, error) {
	var modality *models.Modality

	err := mr.collection().
		FindOne(
			context.TODO(),
			&bson.M{"id": uuid}).
		Decode(&modality)

	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	return modality, err
}

func (mr *ModalityRepository) CreateNewModality(m *models.Modality) error {
	_, err := mr.collection().InsertOne(context.TODO(), m)

	return err
}

func (mr *ModalityRepository) EditModality(uuid *uuid.UUID, m *models.Modality) error {
	filter := bson.M{"id": uuid}
	update := bson.M{
		"$set": bson.M{
			"name":        m.Name,
			"description": m.Description,
		},
	}

	res, err := mr.collection().UpdateOne(context.TODO(), filter, update)

	if res.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return err
}

func (mr *ModalityRepository) DeleteSModalityById(uuid *uuid.UUID) (isFound bool, err error) {
	res, err := mr.collection().
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

func (ur *ModalityRepository) collection() *mongo.Collection {
	return ur.Client.
		Database(config.GetDBName()).
		Collection(config.GetModalitiesCollection())
}
