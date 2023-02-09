package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/online.scheduling-api/config"
	"github.com/online.scheduling-api/src/models"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type IScheduleRepository interface {
	Get(*models.ScheduleFilter) ([]*models.Schedule, error)
	Create(*models.Schedule) error
	Edit(*models.Schedule) error
	DeleteBy(userId, modalityId *uuid.UUID) (isFound bool, err error)
}

type ScheduleRepository struct {
	Client *mongo.Client
}

func (sr *ScheduleRepository) Get(filter *models.ScheduleFilter) ([]*models.Schedule, error) {
	query := bson.M{}

	if filter.ModalityId != uuid.Nil {
		query["modality-id"] = filter.ModalityId
	}

	if filter.ModalityName != "" {
		query["modality-name"] = bson.M{"$regex": filter.ModalityName, "$options": "i"}
	}

	if filter.UserId != uuid.Nil {
		query["user-id"] = filter.UserId
	}

	if filter.UserName != "" {
		query["user-name"] = bson.M{"$regex": filter.UserName, "$options": "i"}
	}

	if filter.Available {
		query["availability.reserved-to"] = uuid.Nil
	}

	if filter.ReservedTo != uuid.Nil {
		query["availability.reserved-to"] = filter.ReservedTo
	}

	cursor, err := sr.collection().Find(context.TODO(), query)
	if err != nil {
		return nil, err
	}

	result := []*models.Schedule{}

	for cursor.Next(context.TODO()) {
		var schedule models.Schedule
		cursor.Decode(&schedule)

		result = append(result, &schedule)
	}

	return result, nil
}

func (sr *ScheduleRepository) Create(schedule *models.Schedule) error {
	_, err := sr.collection().InsertOne(context.TODO(), schedule)

	return err
}

func (sr *ScheduleRepository) Edit(schedule *models.Schedule) error {
	filter := bson.M{
		"user-id":     &schedule.UserId,
		"modality-id": &schedule.ModalityId,
	}

	update := bson.M{
		"$set": bson.M{
			"availability": schedule.Availability,
		},
	}

	res, err := sr.collection().UpdateOne(context.TODO(), filter, update)

	if res.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return err
}

func (sr *ScheduleRepository) DeleteBy(userId, modalityId *uuid.UUID) (isFound bool, err error) {
	res, err := sr.collection().
		DeleteOne(
			context.TODO(),
			&bson.M{
				"user-id":     userId,
				"modality-id": modalityId,
			})

	if err != nil {
		return false, err
	}

	if res.DeletedCount <= 0 {
		return false, nil
	}

	return true, nil
}

func (sr *ScheduleRepository) collection() *mongo.Collection {
	return sr.Client.
		Database(config.GetDBName()).
		Collection("SCHEDULES")
}
