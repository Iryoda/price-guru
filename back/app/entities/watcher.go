package entities

import (
	"github.com/go-playground/validator/v10"
)

type STATUS string

const (
	SUCCESS STATUS = "SUCCESS"
	FAIL           = "FAIL"
	STARTED        = "STARTED"
)

type Watcher struct {
	Id     string `json:"id" bson:"_id,omitempty"`
	UserId string `json:"userId" bson:"userId"`
	Url    string `json:"url" bson:"url"`
	Node   string `json:"node" bson:"node"`
	Name   string `json:"name" bson:"name"`
	Status STATUS `json:"status" bson:"status"`
	Start  int    `json:"start" validate:"required,datetime"`
}

type CreateWatcher struct {
	Url    string `json:"url" validate:"required,url"`
	Name   string `json:"name" validate:"required,min=3"`
	Node   string `json:"node" validate:"required,html"`
	UserId string `json:"userId" validate:"required"`
	Start  int    `json:"start" validate:"omitempty,number,min=1,max=24"`
}

func NewWatcher(data CreateWatcher) (*Watcher, error) {
	validate := validator.New()

	if err := validate.Struct(data); err != nil {
		return nil, err
	}

	if data.Start == 0 {
		data.Start = 12
	}

	watcher := Watcher{
		Url:    data.Url,
		Node:   data.Node,
		UserId: data.UserId,
		Name:   data.Name,
		Start:  data.Start,
		Status: STARTED,
	}

	return &watcher, nil
}
