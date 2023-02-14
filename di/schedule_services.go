package di

import (
	"github.com/online.scheduling-api/constants"
	api "github.com/online.scheduling-api/src/api/handlers"
	"github.com/online.scheduling-api/src/infra/repository"
	"github.com/online.scheduling-api/src/services"
	"github.com/sarulabs/di"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetSchedulingServices() []di.Def {
	return []di.Def{
		{
			Name:  constants.SCHEDULE_REPOSITORY,
			Scope: di.Request,
			Build: func(ctn di.Container) (interface{}, error) {
				return &repository.ScheduleRepository{
					Client: ctn.Get(constants.DB_SERVICE).(*mongo.Client),
				}, nil
			},
		},
		{
			Name:  constants.SCHEDULE_SERVICE,
			Scope: di.Request,
			Build: func(ctn di.Container) (interface{}, error) {
				return &services.ScheduleService{
					ScheduleRepository: ctn.Get(constants.SCHEDULE_REPOSITORY).(*repository.ScheduleRepository),
					UserRespository:    ctn.Get(constants.USER_REPOSITORY).(*repository.UserRepository),
					ModalityRepository: ctn.Get(constants.MODALITY_REPOSITORY).(*repository.ModalityRepository),
				}, nil
			},
		},
		{
			Name:  constants.SCHEDULE_HANDLER,
			Scope: di.Request,
			Build: func(ctn di.Container) (interface{}, error) {
				return &api.SchedulesHandler{
					ScheduleService: ctn.Get(constants.SCHEDULE_SERVICE).(*services.ScheduleService),
				}, nil
			},
		},
	}
}
