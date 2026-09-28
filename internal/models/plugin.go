package models

import (
	"encoding/json"
	"time"
)

type PluginName string

type RenderKind string

type Facet struct {
	PluginName PluginName

	Kind RenderKind

	Data json.RawMessage

	Text string

	CreateAt time.Time
}

type Plugin interface {
	// Name return the plugin name for calling on manager
	Name() PluginName

	// Process function process [Event] with all original data read-only
	// It returns [Facet] and error
	// The Facet is one part and analysis direction about this event
	//
	// if process return non nil error, the result should be dropped
	//
	// WARN: All plugins should not change any original data on the database
	// Plugin can only alter their own table on database
	Process(e *Event) (*Facet, error)
}
