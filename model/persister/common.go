package persister

/*
	Boring SQL CRUD operations for all entities via reflection
*/

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/cjdell/go_angular_starter/model/entity"
	"reflect"
	"strconv"
	"strings"
)

type DB interface {
	Select(dest interface{}, query string, args ...interface{}) error
	Get(dest interface{}, query string, args ...interface{}) error
	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)
	NamedExec(query string, arg interface{}) (sql.Result, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
}

type commonPersister struct {
	db         DB
	entityType reflect.Type
}

func (self commonPersister) getAll(ents interface{}, where string, params ...interface{}) error {
	var (
		table   = entity.GetDbTable(self.entityType)
		columns = entity.GetDbColumns(self.entityType)
	)

	if where == "" {
		where = "1 = 1"
	}

	query := "SELECT e." + strings.Join(columns, ", e.") + " FROM " + table + " AS e WHERE " + where + " ORDER BY e.id ASC"

	return self.db.Select(ents, query, params...)
}

func (self commonPersister) getById(ent entity.Entity, id int64) error {
	var (
		table   = entity.GetDbTable(self.entityType)
		columns = entity.GetDbColumns(self.entityType)
	)

	query := "SELECT e." + strings.Join(columns, ", e.") + " FROM " + table + " AS e WHERE e.id = $1"

	return self.db.Get(ent, query, id)
}

func (self commonPersister) insert(ent entity.Entity) error {
	var (
		table   = entity.GetDbTable(self.entityType)
		columns = entity.GetDbColumns(self.entityType)
	)

	var insertColumns []string

	for _, field := range columns {
		if field != "id" {
			insertColumns = append(insertColumns, field)
		}
	}

	query := "INSERT INTO " + table + " (" + strings.Join(insertColumns, ", ") + ") VALUES (:" + strings.Join(insertColumns, ", :") + ") RETURNING id"

	var id int64

	r, err := self.db.NamedQuery(query, ent)

	if err != nil {
		return err
	}

	r.Next()
	r.Scan(&id)
	r.Close()

	ent.SetId(id)

	return nil
}

func (self commonPersister) update(ent entity.Entity) error {
	var (
		table   = entity.GetDbTable(self.entityType)
		columns = entity.GetDbColumns(self.entityType)
	)

	var fieldNamesPairs []string

	for _, field := range columns {
		if field != "id" {
			fieldNamesPairs = append(fieldNamesPairs, field+" = :"+field)
		}
	}

	query := "UPDATE " + table + " SET " + strings.Join(fieldNamesPairs, ", ") + " WHERE id = :id"

	rows, err := self.db.NamedExec(query, ent)

	if err != nil {
		return err
	}

	count, err := rows.RowsAffected()

	if count != 1 {
		return errors.New("Row updated count incorrect: " + strconv.FormatInt(count, 10))
	}

	return nil
}

func (self commonPersister) delete(id int64) error {
	var (
		table = entity.GetDbTable(self.entityType)
	)

	query := "DELETE FROM " + table + " WHERE id = $1"

	_, err := self.db.Exec(query, id)

	return err
}
