package event_builder

import (
	"bytes"
	"errors"
	"fmt"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/IridiumNan/decision-analyst/internal/models"
)

type buildConfig struct {
	Title  string `toml:"title"`
	Format string `toml:"format"`
}

type EventBuilder struct {
	// confPart is the first part of event.md
	// it sets title and format
	confPart []byte

	conf buildConfig

	// summaryPart is the seconde part of event.md
	// all content will be treated as the summary
	summaryPart []byte

	// sourcePart is the third port of event.md
	// all content remaining will be treated as source
	sourcePart []byte
}

func NewEventBuilder(data []byte) (*EventBuilder, error) {
	b := &EventBuilder{}

	err := b.parse(data)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// parse read from the raw data
// User [EventBuilder.Build] to get event
func (b *EventBuilder) parse(data []byte) error {
	err := b.separate(data)
	if err != nil {
		return fmt.Errorf("parse: err when separate: %s", err.Error())
	}

	err = b.parseHeaderConf()
	if err != nil {
		return fmt.Errorf("parse: err while parse header: %s", err.Error())
	}

	return nil
}

// separate cut different part for different fields
func (b *EventBuilder) separate(data []byte) error {
	var found bool
	_, data, found = bytes.Cut(data, []byte("+++\n"))
	if !found {
		return errors.New("cur prefix +++, the first +++ not found")
	}
	b.confPart, data, found = bytes.Cut(data, []byte("+++\n"))
	if !found {
		return errors.New("toml header config sep +++ not found")
	}

	b.summaryPart, data, found = bytes.Cut(data, []byte("---\n"))

	if !found {
		return errors.New("summary end sep --- not found")
	}

	// TODO: Write content part parse

	return nil
}

// parseHeaderConf get [buildConfig] from [EventBuilder.confPart]
func (b *EventBuilder) parseHeaderConf() error {
	err := toml.Unmarshal(b.confPart, b.conf)
	if err != nil {
		return fmt.Errorf("error while unmarshal header toml conf, err: %s", err.Error())
	}

	return nil
}

func (b *EventBuilder) Build(occurTime *time.Time) (*models.Event, error) {
	ft := models.FormatToType(b.conf.Format)
	if ft == models.FormatInvalid {
		return nil, fmt.Errorf("error while convert format text to type, format text: %s", b.conf.Format)
	}
	return &models.Event{
		Title:     b.conf.Title,
		Summary:   string(b.summaryPart),
		OccurTime: *occurTime,
		Format:    ft,
		Source:    string(b.sourcePart),
		IsDel:     0,
	}, nil
}
