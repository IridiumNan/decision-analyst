package models

import (
	"log"
	"strings"
)

// Database defines all interface for raw event record needed

type Query struct {
	Q string

	Args []any
}

// NewQuery creating a new [Query] for [Database.List]
func NewQuery(sql string, args ...any) Query {
	c := strings.Count(sql, "?")
	if len(args) != c {
		log.Fatalf("args count not match, expected: %d, got: %d", c, len(args))
	}

	return Query{
		Q:    sql,
		Args: args,
	}
}

type Database interface {
	// Add function add a new event into database
	// It will throw error only when database read or write error
	Add(e *Event) error

	// Del function delete the event from database by [Event.ID]
	// it will mark as deleted
	Del(id int) error

	// Update update the event information
	// So ensure that this event id is valid
	//
	Update(id int, attr map[Field]any) error

	// List query events by [Query]
	//
	// if f(e) return false, this event will be dropped
	// id no need to filter any event, just use nil as f
	//
	// the condition is string behind WHERE
	// and all events will sorted by time increment
	//
	// NOTE: By default, the [Event.Source] will be empty, if you want to get them, use [Fetch]
	List(q Query, f func(e *Event) bool) ([]*Event, error)

	// Fetch get whole event information by event id
	Fetch(id string) (*Event, error)

	// Commit all write actions
	Commit() error

	// Close database
	Close() error
}
