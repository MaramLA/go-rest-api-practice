package models

import "time"


type Event struct {
	ID int
	Name string `binding:"required"`
	Description string `binding:"required"`
	Location string `binding:"required"`
	DataTime time.Time 
	UserId int 
}

var events = []Event{}

func (e *Event) Save(
	// name, description, location string, userId int
	){
	// e.Name = name
	// e.Description = description
	// e.Location = location
	// e.DataTime = time.Now()
	// e.UserId = userId

	events = append(events, *e)
}

func GetAllEvents() []Event{
	return events
}