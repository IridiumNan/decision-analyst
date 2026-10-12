package store

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/IridiumNan/decision-analyst/internal/models"
	_ "github.com/IridiumNan/decision-analyst/internal/models"
	_ "github.com/mattn/go-sqlite3"
)

const addSQL = `
	INSERT INTO Event
	(title, summary, occur_time, format, source)
	VALUES (?, ?, ?, ?, ?)`

const delSql = `
	UPDATE Event SET is_delete = 1
	WHERE id = ?`

// SQLiteDB will auto commit before read any data
type SQLiteDB struct {
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
			slog.Error("error while commit transcation", "err", err)
		}
	}

	return nil
}

// Add new event into transcation, call [SQLiteDB.Commit] to sync data
//
// impliment for [models.Database.Add]
func (d *SQLiteDB) Add(e *models.Event) error {
	err := d.prepareWrite()
	if err != nil {
		return fmt.Errorf("Add: error while prepare wirte: %s", err.Error())
	}

	_, err = d.tx.Exec(addSQL, e.Title, e.Summary, e.OccurTime.String(), models.FormatToText(e.Format), e.Source)
	if err != nil {
		return fmt.Errorf("Add: error while add new event to transcation, err: %s", err.Error())
	}

	return nil
}

// Del event from database by id (soft-delete)
//
// impliment for [models.Database.Del]
func (d *SQLiteDB) Del(id int) error {
	err := d.prepareWrite()
	if err != nil {
		return fmt.Errorf("Del: error while prepare write: %s", err.Error())
	}

	_, err = d.tx.Exec(delSql, id)
	if err != nil {
		return fmt.Errorf("Del: error while set event whose id = %s as deleted", id)
	}

	return nil
}

func (d *SQLiteDB) Update(id int, attr map[models.Field]any) {
	models.DropUnmutableFields(attr)

	// TODO: write update sql and logical
}

func (d *SQLiteDB) Commit() error

// NewSQLiteDB create a new db that impliment the [models.Database]
// if sqlite file not exist, it will create a new one
func NewSQLiteDB(sqliteFile string) (*SQLiteDB, error) {
	db, err := sql.Open("sqlite3", sqliteFile)
	if err != nil {
		return nil, fmt.Errorf("error while opening a new sqlite database, err: %s", err.Error())
	}

	return &SQLiteDB{
		db: db,
	}, nil
}
