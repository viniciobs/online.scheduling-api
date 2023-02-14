package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/online.scheduling-api/config"
	"github.com/online.scheduling-api/src/infra/data"
	"github.com/online.scheduling-api/src/models"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type IModalityRepository interface {
	GetModalities(filter *models.ModalityFilter) ([]models.Modality, error)
	GetModalityById(uuid *uuid.UUID) (*models.Modality, error)
	CreateNewModality(m *models.Modality) error
	EditModality(uuid *uuid.UUID, m *models.Modality) error
	DeleteModalityById(uuid *uuid.UUID) (isFound bool, err error)
	ExistsByName(uuid *uuid.UUID, name *string) (bool, error)
}

type ModalityRepository struct {
	DB *data.DB
}

func (mr *ModalityRepository) GetModalities(filter *models.ModalityFilter) ([]models.Modality, error) {
	query := bson.M{}

	if len(filter.Ids) > 0 {
		query["id"] = bson.M{"$in": filter.Ids}
	}

	if filter.Name != "" {
		query["name"] = bson.M{"$regex": filter.Name, "$options": "i"}
	}

	cursor, err := mr.collection().Find(context.TODO(), query)
	if err != nil {
		return nil, err
	}

	modalities := []models.Modality{}

	for cursor.Next(context.TODO()) {
		var modality models.Modality
		cursor.Decode(&modality)

		modalities = append(modalities, modality)
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

func (mr *ModalityRepository) DeleteModalityById(uuid *uuid.UUID) (isFound bool, err error) {
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

func (mr *ModalityRepository) ExistsByName(uuid *uuid.UUID, name *string) (bool, error) {
	err := mr.collection().
		FindOne(
			context.TODO(),
			&bson.M{
				"id":   &bson.M{"$ne": uuid},
				"name": name,
			}).
		Err()

	if err == mongo.ErrNoDocuments {
		return false, nil
	}

	return err == nil, err
}

func (mr *ModalityRepository) collection() *mongo.Collection {
	return mr.DB.Client.
		Database(config.GetDBName()).
		Collection("MODALITIES")
}
