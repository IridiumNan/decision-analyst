package models

import (
	"fmt"
	"log"
	"strings"
)

// Database defines all interface for raw event record needed

// Query is struct used for [Database.List]
type Query struct {
	Q string

	Args []any
}

const ListQueryPrefix = "SELECT id, title, summary, occur_time, create_time, format, is_delete FROM %T "

// NewQuery creating a new [Query] for [Database.List]
//
// it must begin with
// [ListQueryPrefix]
func NewQuery(sql string, args ...any) Query {
	// check if start with select
	if !strings.HasPrefix(sql, ListQueryPrefix) {
		log.Fatalf("error while creating new query, fix your code !! : sql not start with select, got: %s", sql)
	}

	c := strings.Count(sql, "?")
	if len(args) != c {
		log.Fatalf("args count not match, expected: %d, got: %d", c, len(args))
	}

	return Query{
		Q:    sql,
		Args: args,
	}
}

type Attrs struct {
	Attr map[Field]any
}

// Build attr string as
// attr1 = ?, attr2 = ? ...
func (a *Attrs) Build() string {
	var r strings.Builder
	r.Grow(1 << 5)

	for a := range a.Attr {
		fmt.Fprintf(&r, "%s = ?,", string(a))
	}

	// remove last comma
	return strings.TrimRight(r.String(), ",")
}

// NewAttrs create new attrs for update
// it call [DropUnmutableFields] to check fields
func NewAttrs(attr map[Field]any) *Attrs {
	DropUnmutableFields(attr)

	return &Attrs{
		Attr: attr,
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
	Update(id int, a *Attrs) error

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
	Fetch(id int) (*Event, error)

	// Commit all write actions
	Commit() error

	// Close database
	Close() error
}
