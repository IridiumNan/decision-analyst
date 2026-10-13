package store

import (
	"os"
	"testing"
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

	testDBReadWriteFetchIntegration(d, t)

	testDBSimpleAddDelFetch(d, t)
}
