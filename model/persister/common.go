package persister

/*
	Boring SQL CRUD operations for all entities via reflection
*/

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/cjdell/go_angular_starter/model/entity"
	"github.com/jmoiron/sqlx"
	//"log"
	"reflect"
	"strconv"
	"strings"
)

type DB interface {
	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)
	NamedExec(query string, arg interface{}) (sql.Result, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
}

type commonPersister struct {
	db         DB
	entityType reflect.Type
}

type Limit struct {
	Limit  int64
	Offset int64
}

type QueryParameters map[string]interface{}

func NewQueryParametersWithId(id int64) QueryParameters {
	params := QueryParameters{}
	params["id"] = id
	return params
}

type NotFoundError struct {
}

func (NotFoundError) Error() string { return "Not found" }

func (self commonPersister) getAll(ents interface{}, limit *Limit, where string, params QueryParameters) error {
	var (
		table   = entity.GetDbTable(self.entityType)
		columns = entity.GetDbColumns(self.entityType)
	)

	if params == nil {
		params = make(QueryParameters)
	}

	limitClause := ""

	if limit != nil {
		limitClause = fmt.Sprintf(" LIMIT %d OFFSET %d", limit.Limit, limit.Offset)
	}

	query := "SELECT e." + strings.Join(columns, ", e.") + " FROM " + table + " AS e " + where + " ORDER BY e.id ASC" + limitClause

	//log.Println(query)

	rows, err := self.db.NamedQuery(query, map[string]interface{}(params))

	if err != nil {
		return err
	}

	defer rows.Close()

	return sqlx.StructScan(rows, ents)
}

func (self commonPersister) getOne(ent entity.Entity, where string, params QueryParameters) error {
	var (
		table   = entity.GetDbTable(self.entityType)
		columns = entity.GetDbColumns(self.entityType)
	)

	if params == nil {
		params = make(QueryParameters)
	}

	query := "SELECT e." + strings.Join(columns, ", e.") + " FROM " + table + " AS e " + where + " LIMIT 1"

	//log.Println(query)

	rows, err := self.db.NamedQuery(query, map[string]interface{}(params))

	if err != nil {
		return err
	}

	defer rows.Close()

	next := rows.Next()

	if !next {
		return NotFoundError{}
	}

	return rows.StructScan(ent)
}

func (self commonPersister) insert(ent entity.Entity) (int64, error) {
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
		return -1, err
	}

	r.Next()

	err = r.Scan(&id)

	if err != nil {
		return -1, err
	}

	r.Close()

	return id, nil
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
