package models

// Database defines all interface for raw event record needed

type Database interface {
	// Add function add a new event into database
	// It will throw error only when database read or write error
	Add(e *Event) error

	// Del function delete the event from database by [Event.ID]
	// it will mark as deleted
	Del(id int) error

	// Update update the event information
	// It will check this by id
	// So ensure that this event id is valid
	Update(e *Event) error

	// List load all events then filter them with function
	// If return false, this event will be dropped
	//
	// On the return slice, the [Event.Source] Field will all emptyStr
	List(func(e *Event) bool) ([]*Event, error)

	// Fetch get whole event information by event id
	Fetch(id string) (*Event, error)
}
