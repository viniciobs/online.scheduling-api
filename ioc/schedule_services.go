package ioc

import (
	api "github.com/online.scheduling-api/src/api/handlers"
	"github.com/online.scheduling-api/src/infra/repository"
	"github.com/online.scheduling-api/src/services"
	"github.com/sarulabs/di"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetSchedulingServices() []di.Def {
	return []di.Def{
		{
			Name:  "schedule-repository",
			Scope: di.Request,
			Build: func(ctn di.Container) (interface{}, error) {
				return &repository.ScheduleRepository{
					Client: ctn.Get("mongo").(*mongo.Client),
				}, nil
			},
		},
		{
			Name:  "schedule-service",
			Scope: di.Request,
			Build: func(ctn di.Container) (interface{}, error) {
				return &services.ScheduleService{
					ScheduleRepository: ctn.Get("schedule-repository").(*repository.ScheduleRepository),
					UserRespository:    ctn.Get("user-repository").(*repository.UserRepository),
					ModalityRepository: ctn.Get("modality-repository").(*repository.ModalityRepository),
				}, nil
			},
		},
		{
			Name:  "schedule-handler",
			Scope: di.Request,
			Build: func(ctn di.Container) (interface{}, error) {
				return &api.SchedulesHandler{
					ScheduleService: ctn.Get("schedule-service").(*services.ScheduleService),
				}, nil
			},
		},
	}
}
