package models

import (
	"testing"
	"time"
)

func TestQueryAll(t *testing.T) {
	q := QueryAll()
	expectedPrefix := ListQueryPrefix + "WHERE create_time < ?"
	if q.Q != expectedPrefix {
		t.Errorf("Expected query %s, got %s", expectedPrefix, q.Q)
	}
	if len(q.Args) != 1 {
		t.Errorf("Expected 1 argument, got %d", len(q.Args))
	}
}

func TestNewQuery(t *testing.T) {
	sql := ListQueryPrefix + "WHERE id = ?"
	q := NewQuery(sql, 1)
	if q.Q != sql {
		t.Errorf("Expected query %s, got %s", sql, q.Q)
	}
	if len(q.Args) != 1 || q.Args[0] != 1 {
		t.Errorf("Expected args [1], got %v", q.Args)
	}
}

func TestAttrsBuild(t *testing.T) {
	attrs := &Attrs{
		Attr: map[Field]any{
			FieldTitle: "New Title",
		},
	}
	built := attrs.Build()
	if built != "title = ?" {
		t.Errorf("Expected 'title = ?', got '%s'", built)
	}
}

func TestNewAttrsAndDropUnmutableFields(t *testing.T) {
	rawAttrs := map[Field]any{
		FieldTitle:   "Test Title",
		FieldID:      123, // Unmutable, should be dropped
		FieldSummary: "Test Summary",
	}
	attrs := NewAttrs(rawAttrs)

	if _, exists := attrs.Attr[FieldID]; exists {
		t.Errorf("Expected unmutable field FieldID to be dropped")
	}
	if attrs.Attr[FieldTitle] != "Test Title" {
		t.Errorf("Expected FieldTitle to be preserved")
	}
	if attrs.Attr[FieldSummary] != "Test Summary" {
		t.Errorf("Expected FieldSummary to be preserved")
	}
}

func TestIsMutableField(t *testing.T) {
	if !IsMutableField(FieldTitle) {
		t.Errorf("Expected FieldTitle to be mutable")
	}
	if IsMutableField(FieldID) {
		t.Errorf("Expected FieldID to be immutable")
	}
}

func TestEventEqual(t *testing.T) {
	now := time.Now()
	e1 := &Event{
		ID:         1,
		Title:      "Event 1",
		Summary:    "Summary 1",
		OccurTime:  now,
		CreateTime: now,
		Format:     FormatMarkdown,
		Source:     "source 1",
		IsDel:      0,
	}
	e2 := &Event{
		ID:         1,
		Title:      "Event 1",
		Summary:    "Summary 1",
		OccurTime:  now,
		CreateTime: now,
		Format:     FormatMarkdown,
		Source:     "source 1",
		IsDel:      0,
	}

	equal, field := e1.Equal(e2)
	if !equal {
		t.Errorf("Expected events to be equal, but mismatched at %s", field)
	}

	// Test mismatch scenario
	e3 := &Event{
		ID:         1,
		Title:      "Event 2", // Different title
		Summary:    "Summary 1",
		OccurTime:  now,
		CreateTime: now,
		Format:     FormatMarkdown,
		Source:     "source 1",
		IsDel:      0,
	}

	equal, field = e1.Equal(e3)
	if equal || field != FieldTitle {
		t.Errorf("Expected mismatch at FieldTitle, got equal=%v, field=%s", equal, field)
	}
}
