package models

import (
	"time"
)

type Event struct {
	ID int64
	Name string `binding:"required"`
	Description string `binding:"required"`
	Location string `binding:"required"`
	DateTime time.Time `binding:"required"`
	UserId int
}

var events = []Event{
	{
		ID: 1,
		Name: "test prod",
		Description: "test description2",
		Location: "test location",
		DateTime: time.Now(),
		UserId: 1,
	},
}

func GetAllEvents() ([]Event, error) {
	return events, nil
}
