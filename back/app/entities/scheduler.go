package entities

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Scheduler struct {
	Id      string    `json:"id"`
	Start   time.Time `json:"start"`
	UserId  string    `json:"userId"`
	LastRun time.Time `json:"lastRun"`
}

type CreateScheduler struct {
	Start time.Time `json:"start" validate:"required,datetime"`
}

func NewScheduler(data CreateScheduler) (*Scheduler, error) {
	validate := validator.New()

	if err := validate.Struct(data); err != nil {
		return nil, err
	}

	return &Scheduler{
		Start:   data.Start,
		LastRun: time.Now(),
	}, nil
}
