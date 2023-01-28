package ioc

import (
	"github.com/online.scheduling-api/src/helpers"
	"github.com/sarulabs/di"
)

var Services = helpers.Concat(
	[][]di.Def{
		GetDataServices(),
		GetUserServices(),
		GetModalityServices(),
	})
