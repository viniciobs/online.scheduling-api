package models

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	UserId       uuid.UUID      `bson:"user-id"`
	CreatedAt    time.Time      `bson:"created-at"`
	DueValue     float32        `bson:"due-value"`
	Transactions []*Transaction `bson:"transactions"`
}

type Transaction struct {
	Value float32   `bson:"value"`
	Time  time.Time `bson:"Time"`
}

func (p *Payment) IsFullPaid() bool {
	dueValue := p.DueValue

	for _, t := range p.Transactions {
		dueValue -= t.Value
	}

	return dueValue == 0.0
}
