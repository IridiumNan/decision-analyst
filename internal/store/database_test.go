package store

import (
	"testing"
	"time"

	"github.com/IridiumNan/decision-analyst/internal/models"
	_ "github.com/mattn/go-sqlite3"
)

// Below is an example of how you can structure your real CRUD test.
// Replace `NewSQLiteDatabase` with your actual constructor function.

func testDBReadWriteFetchIntegration(db models.Database, t *testing.T) {
	now := time.Now().UTC().Truncate(time.Second)

	// 1. Test Add (Write)
	newEvent := &models.Event{
		Title:     "Integration Test Event",
		Summary:   "Testing real SQLite read and write",
		OccurTime: now,
		Format:    models.FormatMarkdown,
		Source:    "Test Source Content",
		IsDel:     0,
	}

	// Example Add call:
	err := db.Add(newEvent)
	if err != nil {
		t.Fatalf("Failed to add event: %v", err)
	}

	// 2. Test Fetch (Read by ID)
	// Assuming the inserted event got ID = 1
	fetchedEvent, err := db.Fetch(1)
	if err != nil {
		t.Fatalf("Failed to fetch event: %v", err)
	}
	if fetchedEvent.Title != newEvent.Title {
		t.Errorf("Expected title %s, got %s", newEvent.Title, fetchedEvent.Title)
	}

	// 3. Test Update
	attrs := models.NewAttrs(map[models.Field]any{
		models.FieldTitle: "Updated Integration Title",
	})
	t.Logf("the build str: %s;  attr: %v", attrs.Build(), attrs.Attr)
	err = db.Update(1, attrs)
	if err != nil {
		t.Fatalf("Failed to update event: %v", err)
	}

	// 4. Test List Query
	query := models.NewQuery(models.ListQueryPrefix+"WHERE id = ?", 1)
	t.Logf("query with sqlstr: %s", query)
	events, err := db.List(query, nil)
	if err != nil {
		t.Fatalf("Failed to list events: %v", err)
	}
	if len(events) != 1 || events[0].Title != "Updated Integration Title" {
		t.Errorf("List query did not return expected updated record")
	}

	// 5. Test Del (Soft Delete)
	err = db.Del(1)
	if err != nil {
		t.Fatalf("Failed to soft delete event: %v", err)
	}

	// Verify soft delete flag (IsDel == 1)
	deletedEvent, err := db.Fetch(1)
	if err != nil {
		t.Fatalf("Failed to fetch after delete: %v", err)
	}
	if deletedEvent.IsDel != 1 {
		t.Errorf("Expected IsDel to be 1, got %d", deletedEvent.IsDel)
	}
}

func testDBSimpleAddDelFetch(d models.Database, t *testing.T) {
	e := &models.Event{
		Title:     "Hello World",
		Summary:   "This is a test event for hello world",
		OccurTime: time.Date(2025, time.October, 30, 0, 0, 0, 0, models.Local),
		Format:    models.FormatMarkdown,
		Source:    "this is test source",
	}

	firstID := 2

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
