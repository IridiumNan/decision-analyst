package store

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/IridiumNan/decision-analyst/internal/models"
	_ "github.com/IridiumNan/decision-analyst/internal/models"
	"github.com/IridiumNan/decision-analyst/pkgs/utils"
	_ "github.com/mattn/go-sqlite3"
)

// SQLiteDB will auto commit before read any data
type SQLiteDB struct {
	tableName string

	db *sql.DB

	tx *sql.Tx
}

// prepareWrite do prepare for write, it ensure [SQLiteDB.tx] is not nil
func (d *SQLiteDB) prepareWrite() error {
	if d.tx == nil {
		t, err := d.db.Begin()
		if err != nil {
			return err
		}

		d.tx = t
	}

	return nil
}

// prepareRead ensure last tx has been commited
func (d *SQLiteDB) prepareRead() error {
	if d.tx != nil {
		err := d.tx.Commit()
		if err != nil {
			return err
		}
		d.tx = nil
	}

	return nil
}

const addSQL = `
	INSERT INTO %T
	(title, summary, occur_time, format, source)
	VALUES (?, ?, ?, ?, ?)`

// Add new event into transcation, call [SQLiteDB.Commit] to sync data
//
// impliment for [models.Database.Add]
func (d *SQLiteDB) Add(e *models.Event) error {
	err := d.prepareWrite()
	if err != nil {
		return fmt.Errorf("Add: error while prepare wirte: %s", err.Error())
	}

	sqlStr := d.replaceTable(addSQL)
	_, err = d.tx.Exec(sqlStr, e.Title, e.Summary, e.OccurTime, models.FormatToText(e.Format), e.Source)
	if err != nil {
		return fmt.Errorf("Add: error while add new event to transcation, err: %s", err.Error())
	}

	return nil
}

const delSQL = `
	UPDATE %T SET is_delete = 1 
	WHERE id = ?
	`

// Del event from database by id (soft-delete)
//
// impliment for [models.Database.Del]
func (d *SQLiteDB) Del(id int) error {
	err := d.prepareWrite()
	if err != nil {
		return fmt.Errorf("Del: error while prepare write: %s", err.Error())
	}

	sqlStr := d.replaceTable(delSQL)

	_, err = d.tx.Exec(sqlStr, id)
	if err != nil {
		return fmt.Errorf("Del: error while set event whose id = %d as deleted", id)
	}

	return nil
}

const updateSQL = `
	UPDATE %T SET %U WHERE id = ?`

// Update event mutable fields by id
//
// impliment for [models.Database.Update]
func (d *SQLiteDB) Update(id int, a *models.Attrs) error {
	err := d.prepareWrite()
	if err != nil {
		return fmt.Errorf("Update: error while prepare write: %s", err.Error())
	}
	sqlStr := strings.Replace(updateSQL, "%U", a.Build(), 1)
	sqlStr = d.replaceTable(sqlStr)

	args := utils.MapV(a.Attr)
	args = append(args, id)
	_, err = d.tx.Exec(sqlStr, args...)
	if err != nil {
		return fmt.Errorf("Update: failed to update, sqlStr: %s, err: %s", sqlStr, err.Error())
	}

	return nil
}

// rowToEvent convert all data on [sql.Rows] to event slice
// NOTE: it just scan for List fields:
// id
// title
// summary
// occur_time
// create_time
// format
func (d *SQLiteDB) rowsToEvent(r *sql.Rows) ([]*models.Event, error) {
	// query start with
	// SELECT id, title, summary, occur_time, create_time, format
	events := make([]*models.Event, 0, 10)
	errs := make([]error, 0, 10)

	for r.Next() {
		var id int
		var title string
		var summary string
		var occurTime time.Time
		var createTime time.Time
		var formatStr string
		var isDel int

		err := r.Scan(&id, &title, &summary, &occurTime, &createTime, &formatStr, &isDel)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		eve := &models.Event{
			ID:         id,
			Title:      title,
			Summary:    summary,
			OccurTime:  occurTime,
			CreateTime: createTime,
			Format:     models.FormatToType(formatStr),
			Source:     "",
			IsDel:      isDel,
		}

		events = append(events, eve)
	}

	var errMsg strings.Builder
	if len(errs) > 0 {
		errMsg.Grow(1 << 6)
		fmt.Fprintf(&errMsg, "rowToEvent: %d errors occured while scan", len(errs))
		for i, e := range errs {
			fmt.Fprintf(&errMsg, "err %d: %s", i+1, e.Error())
		}

		return events, errors.New(errMsg.String())
	}

	return events, nil
}

// List get events with [models.Event.Source] = ""
// impliment for [models.Database.List]
func (d *SQLiteDB) List(q models.Query, f func(e *models.Event) bool) ([]*models.Event, error) {
	err := d.prepareRead()
	if err != nil {
		return nil, fmt.Errorf("List: error while prepare read, err: %s", err.Error())
	}

	sqlStr := d.replaceTable(q.Q)
	rows, err := d.db.Query(sqlStr, q.Args...)
	if err != nil {
		errAppend := strings.Repeat("%v, ", len(q.Args))

		args := make([]any, 0, len(q.Args)+1)
		args = append(args, sqlStr)
		args = append(args, q.Args...)

		return nil, fmt.Errorf("List: error while query, sql: %s, "+errAppend, args...)
	}
	defer rows.Close()

	events, err := d.rowsToEvent(rows)

	r := make([]*models.Event, 0, len(events))
	for _, eve := range events {
		if f(eve) {
			r = append(r, eve)
		}
	}

	return r, err
}

const fetchSQL = `
	SELECT title, summary, occur_time, create_time, format, source, is_delete FROM %T
	WHERE id = ?`

// Fetch the impliment for [models.Database.Fetch]
func (d *SQLiteDB) Fetch(id int) (*models.Event, error) {
	err := d.prepareRead()
	if err != nil {
		return nil, fmt.Errorf("Fetch: error while prepare read, err: %s", err.Error())
	}

	sqlStr := d.replaceTable(fetchSQL)
	r := d.db.QueryRow(sqlStr, id)

	var title string
	var summary string
	var occurTime time.Time
	var createTime time.Time
	var formatStr string
	var source string
	var isDel int

	if err := r.Scan(&title, &summary, &occurTime, &createTime, &formatStr, &source, &isDel); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Fetch: no row found for id: %d, err: %s", id, err.Error())
		}
		return nil, fmt.Errorf("Fetch: error while scanning: %s", err.Error())
	}

	event := &models.Event{
		ID:         id,
		Title:      title,
		Summary:    summary,
		OccurTime:  occurTime,
		CreateTime: createTime,
		Format:     models.FormatToType(formatStr),
		Source:     source,
		IsDel:      isDel,
	}

	return event, nil
}

// Commit  this transcation
func (d *SQLiteDB) Commit() error {
	if d.tx == nil {
		return nil
	}

	err := d.tx.Commit()
	if err != nil {
		return fmt.Errorf("Commit: error while commit tx: %s", err.Error())
	}

	return nil
}

// Close cimmit then close this db connection
func (d *SQLiteDB) Close() error {
	if err := d.Commit(); err != nil {
		return err
	}
	return d.db.Close()
}

// NewSQLiteDB create a new db that impliment the [models.Database]
// if sqlite file not exist, it will create a new one
//
// Default table name is Event
func NewSQLiteDB(sqliteFile string) (*SQLiteDB, error) {
	db, err := sql.Open("sqlite3", sqliteFile)
	if err != nil {
		return nil, fmt.Errorf("error while opening a new sqlite database, err: %s", err.Error())
	}

	return &SQLiteDB{
		db: db,
	}, nil
}

// SetTableName reset the table name
func (d *SQLiteDB) SetTableName(tableName string) {
	d.tableName = tableName
}

// replaceTable replace the %T on sqlStr with [SQLiteDB.tableName]
func (d *SQLiteDB) replaceTable(sqlStr string) string {
	return strings.Replace(sqlStr, "%T", d.tableName, 1)
}
