package services

import (
	"github.com/cjdell/go_angular_starter/model/entity"
	"github.com/cjdell/go_angular_starter/model/persister"
)

type <%= entityNameSingularPascalCase %>Service struct {
	persister *persister.<%= entityNameSingularPascalCase %>Persister
	user      *entity.User
}

// A struct composed of the record coupled with additional computed information
type <%= entityNameSingularPascalCase %>Info struct {
	*entity.<%= entityNameSingularPascalCase %>
}

// A struct to hold additional modifications that don't fit into the record data structure
type <%= entityNameSingularPascalCase %>Save struct {
	*entity.<%= entityNameSingularPascalCase %>

	Changes struct {
		Fields []string
	}
}

func New<%= entityNameSingularPascalCase %>Service(db persister.DB, user *entity.User) *<%= entityNameSingularPascalCase %>Service {
	return &<%= entityNameSingularPascalCase %>Service{persister.New<%= entityNameSingularPascalCase %>Persister(db), user}
}

func (self *<%= entityNameSingularPascalCase %>Service) GetAll(limit *persister.Limit) ([]*<%= entityNameSingularPascalCase %>Info, error) {
	var err error
	var <%= entityNamePluralCamelCase %> []*entity.<%= entityNameSingularPascalCase %>

	if <%= entityNamePluralCamelCase %>, err = self.persister.GetAll(limit); err != nil {
		return nil, err
	}

	<%= entityNameSingularCamelCase %>Infos := make([]*<%= entityNameSingularPascalCase %>Info, len(<%= entityNamePluralCamelCase %>), len(<%= entityNamePluralCamelCase %>))

	for i, <%= entityNameSingularCamelCase %> := range <%= entityNamePluralCamelCase %> {
		<%= entityNameSingularCamelCase %>Infos[i], _ = self.<%= entityNameSingularCamelCase %>Info(<%= entityNameSingularCamelCase %>)
	}

	return <%= entityNameSingularCamelCase %>Infos, nil
}

func (self *<%= entityNameSingularPascalCase %>Service) GetOne(id int64) (*<%= entityNameSingularPascalCase %>Info, error) {
	var err error
	var <%= entityNameSingularCamelCase %> *entity.<%= entityNameSingularPascalCase %>

	if <%= entityNameSingularCamelCase %>, err = self.persister.GetById(id); err != nil {
		return nil, err
	}

	return self.<%= entityNameSingularCamelCase %>Info(<%= entityNameSingularCamelCase %>)
}

func (self *<%= entityNameSingularPascalCase %>Service) Insert(<%= entityNameSingularCamelCase %>Save *<%= entityNameSingularPascalCase %>Save) (*<%= entityNameSingularPascalCase %>Info, error) {
	var err error

	<%= entityNameSingularCamelCase %> := <%= entityNameSingularCamelCase %>Save.<%= entityNameSingularPascalCase %>

	if err = self.beforeSave(<%= entityNameSingularCamelCase %>, <%= entityNameSingularCamelCase %>Save); err != nil {
		return nil, err
	}

	if <%= entityNameSingularCamelCase %>.Id, err = self.persister.Insert(<%= entityNameSingularCamelCase %>); err != nil {
		return nil, err
	}

	if err = self.afterSave(<%= entityNameSingularCamelCase %>, <%= entityNameSingularCamelCase %>Save); err != nil {
		return nil, err
	}

	return self.GetOne(<%= entityNameSingularCamelCase %>.Id)
}

func (self *<%= entityNameSingularPascalCase %>Service) Update(<%= entityNameSingularCamelCase %>Save *<%= entityNameSingularPascalCase %>Save) (*<%= entityNameSingularPascalCase %>Info, error) {
	var err error
	var <%= entityNameSingularCamelCase %> *entity.<%= entityNameSingularPascalCase %>

	if <%= entityNameSingularCamelCase %>, err = self.persister.GetById(<%= entityNameSingularCamelCase %>Save.Id); err != nil {
		return nil, err
	}

	if err = self.beforeSave(<%= entityNameSingularCamelCase %>, <%= entityNameSingularCamelCase %>Save); err != nil {
		return nil, err
	}

	if err = self.persister.Update(<%= entityNameSingularCamelCase %>); err != nil {
		return nil, err
	}

	if err = self.afterSave(<%= entityNameSingularCamelCase %>, <%= entityNameSingularCamelCase %>Save); err != nil {
		return nil, err
	}

	return self.GetOne(<%= entityNameSingularCamelCase %>.Id)
}

func (self *<%= entityNameSingularPascalCase %>Service) Delete(id int64) error {
	return self.persister.Delete(id)
}

// Wrap <%= entityNameSingularPascalCase %> into <%= entityNameSingularPascalCase %>Info - Add computed properties here
func (self *<%= entityNameSingularPascalCase %>Service) <%= entityNameSingularCamelCase %>Info(<%= entityNameSingularCamelCase %> *entity.<%= entityNameSingularPascalCase %>) (*<%= entityNameSingularPascalCase %>Info, error) {
	return &<%= entityNameSingularPascalCase %>Info{<%= entityNameSingularCamelCase %>}, nil
}

// Handle <%= entityNameSingularPascalCase %>Save - before saving to the database
func (self *<%= entityNameSingularPascalCase %>Service) beforeSave(<%= entityNameSingularCamelCase %> *entity.<%= entityNameSingularPascalCase %>, <%= entityNameSingularCamelCase %>Save *<%= entityNameSingularPascalCase %>Save) error {
	return <%= entityNameSingularCamelCase %>.Merge(<%= entityNameSingularCamelCase %>Save.<%= entityNameSingularPascalCase %>, <%= entityNameSingularCamelCase %>Save.Changes.Fields)
}

// Handle <%= entityNameSingularPascalCase %>Save - after saving to the database
func (self *<%= entityNameSingularPascalCase %>Service) afterSave(<%= entityNameSingularCamelCase %> *entity.<%= entityNameSingularPascalCase %>, <%= entityNameSingularCamelCase %>Save *<%= entityNameSingularPascalCase %>Save) error {
	return nil
}
