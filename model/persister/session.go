package persister

import (
	"github.com/jmoiron/sqlx"
	"github.com/cjdell/go_angular_starter/model/entity"
)

type SessionPersister struct {
	db *sqlx.DB
}

func NewSessionPersister(db *sqlx.DB) *SessionPersister {
	return &SessionPersister{db}
}

func (self SessionPersister) GetByKey(key string) error {
	query := `SELECT s.* FROM sessions AS s WHERE key = :key`

	session := &entity.Session{}
	err := self.db.Get(session, query, map[string]interface{}{"key": key})

	if err != nil {
		return err
	}

	return nil
}

func (self SessionPersister) Insert(session *entity.Session) error {
	query := `INSERT INTO sessions (user_id, api_key) VALUES (:user_id, :api_key)`

	var id int64

	r, err := self.db.NamedQuery(query, session)

	if err != nil {
		return err
	}

	r.Next()
	r.Scan(&id)

	session.Id = id

	return nil
}

func (self SessionPersister) Delete(id int64) error {
	query := `DELETE FROM sessions WHERE id = :id`

	_, err := self.db.NamedExec(query, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	return nil
}
