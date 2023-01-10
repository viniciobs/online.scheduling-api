package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUUID(c *gin.Context) (*uuid.UUID, error) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return nil, err
	}

	return &id, nil
}
