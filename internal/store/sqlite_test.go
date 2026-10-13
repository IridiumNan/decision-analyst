package store

import (
	"os"
	"testing"
	"time"

	"github.com/IridiumNan/decision-analyst/internal/models"
)

const testDBFile = "main.db"

func TestSqliteDB(t *testing.T) {
	err := os.Remove(testDBFile)
	if err != nil && !os.IsNotExist(err) {
		t.Errorf("error while removing old db file, path: %s, err: %s", testDBFile, err.Error())
	}
	d, err := NewSQLiteDB(testDBFile)
	if err != nil {
		t.Errorf("error while creating a new sqlite db, err: %s", err.Error())
	}

	defer d.Close()

	testDatabaseReadWriteFetchIntegration(d, t)
}

func dbTest1(d models.Database, t *testing.T) {
	e := &models.Event{
		Title:     "Hello World",
		Summary:   "This is a test event for hello world",
		OccurTime: time.Date(2025, time.October, 30, 0, 0, 0, 0, models.Local),
		Format:    models.FormatMarkdown,
		Source:    "this is test source",
	}

	firstID := 1

	exp := &models.Event{
		ID:        firstID,
		Title:     "Hello World",
		Summary:   "This is a test event for hello world",
		OccurTime: time.Date(2025, time.October, 30, 0, 0, 0, 0, models.Local),
		Format:    models.FormatMarkdown,
		Source:    "this is test source",
	}

	err := d.Add(e)
	if err != nil {
		t.Errorf("error while adding new event, err: %s", err.Error())
	}

	// fetch will get the total source
	got, err := d.Fetch(firstID)

	exp.CreateTime = got.CreateTime
	if err != nil {
		t.Errorf("error while fetching first event")
	}

	if equal, field := exp.Equal(got); !equal {
		t.Errorf("first added event not match, field: %s", field)
	}

	err = d.Del(firstID)
	if err != nil {
		t.Errorf("error while del event with id = %d, err: %s", firstID, err.Error())
	}

	got, err = d.Fetch(firstID)
	if err != nil {
		t.Errorf("error while fetching first event after del")
	}
	exp.IsDel = 1

	if equal, field := exp.Equal(got); !equal {
		t.Errorf("first added event after del not match, field: %s", field)
	}
}
