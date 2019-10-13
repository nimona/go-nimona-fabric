package rel

import (
	"database/sql"
	"time"

	"nimona.io/pkg/errors"
)

const migrationsTable string = `
CREATE TABLE IF NOT EXISTS Migrations (
	ID INTEGER NOT NULL PRIMARY KEY,
	LastIndex INTEGER,
	Datetime INT
)`

type DB struct {
	db *sql.DB
}

type migrationRow struct {
	id        int
	LastIndex int
	Datetime  string
}

func New(db *sql.DB) (*DB, error) {
	ndb := &DB{
		db: db,
	}

	err := ndb.createMigrationTable()
	err = ndb.runMigrations()

	return ndb, err
}

func (d *DB) Close() error {
	return d.db.Close()
}

func (d *DB) createMigrationTable() error {
	_, err := d.db.Exec(migrationsTable)
	if err != nil {
		return errors.Wrap(err, errors.New("could not create migrations table"))
	}

	return nil
}

func (d *DB) runMigrations() error {
	tx, err := d.db.Begin()
	if err != nil {
		return errors.Wrap(err, errors.New("could not start transaction"))
	}

	for index, mig := range migrations {

		rows, err := tx.Query("select ID, LastIndex, Datetime from Migrations order by id desc limit 1")
		if err != nil {
			tx.Rollback()
			return errors.Wrap(err, errors.New("could not run migration"))
		}

		mgr := migrationRow{}

		for rows.Next() {
			rows.Scan(&mgr.id, &mgr.LastIndex, &mgr.Datetime)
		}

		if mgr.id > 0 && mgr.LastIndex >= index {
			continue
		}

		_, err = tx.Exec(mig)
		if err != nil {
			tx.Rollback()
			return errors.Wrap(err, errors.New("could not run migration"))
		}

		stmt, err := tx.Prepare(
			"INSERT INTO Migrations(LastIndex, Datetime) VALUES(?, ?)")
		if err != nil {
			tx.Rollback()
			return errors.Wrap(err, errors.New("could not insert to migrations table"))
		}

		_, err = stmt.Exec(index, time.Now().Unix())
		if err != nil {
			tx.Rollback()
			return errors.Wrap(err, errors.New("could not insert to migrations table"))
		}
	}

	tx.Commit()

	return nil
}
