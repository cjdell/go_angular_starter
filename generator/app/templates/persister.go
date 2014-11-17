package persister

import (
	"github.com/cjdell/go_angular_starter/model/entity"
	"reflect"
)

type <%= entityNameSingularPascalCase %>Persister struct {
	db     DB
	common *commonPersister
}

func New<%= entityNameSingularPascalCase %>Persister(db DB) *<%= entityNameSingularPascalCase %>Persister {
	entityType := reflect.TypeOf(&entity.<%= entityNameSingularPascalCase %>{}).Elem()
	return &<%= entityNameSingularPascalCase %>Persister{db, &commonPersister{db, entityType}}
}

func (self <%= entityNameSingularPascalCase %>Persister) GetAll(limit *Limit) ([]*entity.<%= entityNameSingularPascalCase %>, error) {
	<%= entityNamePluralCamelCase %> := []*entity.<%= entityNameSingularPascalCase %>{}
	return <%= entityNamePluralCamelCase %>, self.common.getAll(&<%= entityNamePluralCamelCase %>, limit, "", nil)
}

func (self <%= entityNameSingularPascalCase %>Persister) GetById(id int64) (*entity.<%= entityNameSingularPascalCase %>, error) {
	<%= entityNameSingularCamelCase %> := &entity.<%= entityNameSingularPascalCase %>{}
	return <%= entityNameSingularCamelCase %>, self.common.getOne(<%= entityNameSingularCamelCase %>, "WHERE id = :id", NewQueryParametersWithId(id))
}

func (self <%= entityNameSingularPascalCase %>Persister) Insert(<%= entityNameSingularCamelCase %> *entity.<%= entityNameSingularPascalCase %>) (int64, error) {
	return self.common.insert(<%= entityNameSingularCamelCase %>)
}

func (self <%= entityNameSingularPascalCase %>Persister) Update(<%= entityNameSingularCamelCase %> *entity.<%= entityNameSingularPascalCase %>) error {
	return self.common.update(<%= entityNameSingularCamelCase %>)
}

func (self <%= entityNameSingularPascalCase %>Persister) Delete(id int64) error {
	return self.common.delete(id)
}
