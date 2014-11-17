package persister

import (
	"github.com/cjdell/go_angular_starter/model/entity"
	"reflect"
)

type SessionPersister struct {
	db     DB
	common *commonPersister
}

func NewSessionPersister(db DB) *SessionPersister {
	entityType := reflect.TypeOf(&entity.Session{}).Elem()
	return &SessionPersister{db, &commonPersister{db, entityType}}
}

func (self SessionPersister) GetByKey(key string) (*entity.Session, error) {
	session := &entity.Session{}

	params := QueryParameters{}
	params["key"] = key

	return session, self.common.getOne(session, "WHERE key = :key", params)
}

func (self SessionPersister) Insert(session *entity.Session) error {
	_, err := self.common.insert(session)
	return err
}

func (self SessionPersister) Delete(id int64) error {
	return self.common.delete(id)
}
