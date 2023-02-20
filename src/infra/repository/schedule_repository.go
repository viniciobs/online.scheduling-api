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

type IScheduleRepository interface {
	Get(ctx context.Context, filter *models.ScheduleFilter) ([]*models.Schedule, error)
	Create(ctx context.Context, schedule *models.Schedule) error
	Edit(ctx context.Context, schedule *models.Schedule) error
	DeleteBy(ctx context.Context, userId, modalityId *uuid.UUID) (isFound bool, err error)
}

type ScheduleRepository struct {
	DB *data.DB
}

func (sr *ScheduleRepository) Get(ctx context.Context, filter *models.ScheduleFilter) ([]*models.Schedule, error) {
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

	cursor, err := sr.collection().Find(ctx, query)
	if err != nil {
		return nil, err
	}

	result := []*models.Schedule{}

	for cursor.Next(ctx) {
		var schedule models.Schedule
		cursor.Decode(&schedule)

		result = append(result, &schedule)
	}

	return result, nil
}

func (sr *ScheduleRepository) Create(ctx context.Context, schedule *models.Schedule) error {
	_, err := sr.collection().InsertOne(ctx, schedule)

	return err
}

func (sr *ScheduleRepository) Edit(ctx context.Context, schedule *models.Schedule) error {
	filter := bson.M{
		"user-id":     &schedule.UserId,
		"modality-id": &schedule.ModalityId,
	}

	update := bson.M{
		"$set": bson.M{
			"availability": schedule.Availability,
		},
	}

	res, err := sr.collection().UpdateOne(ctx, filter, update)

	if res.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return err
}

func (sr *ScheduleRepository) DeleteBy(ctx context.Context, userId, modalityId *uuid.UUID) (isFound bool, err error) {
	query := bson.M{}

	if *userId != uuid.Nil {
		query["user-id"] = userId
	}

	if *modalityId != uuid.Nil {
		query["modality-id"] = modalityId
	}

	res, err := sr.collection().DeleteOne(ctx, query)

	if err != nil {
		return false, err
	}

	if res.DeletedCount <= 0 {
		return false, nil
	}

	return true, nil
}

func (sr *ScheduleRepository) collection() *mongo.Collection {
	return sr.DB.Client.
		Database(config.GetDBName()).
		Collection("SCHEDULES")
}
