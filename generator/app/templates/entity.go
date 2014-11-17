package entity

type <%= entityNameSingularPascalCase %> struct {
	__table struct{} `db:"<%= entityNamePluralSnakeCase %>"`

	Id   int64  `db:"id"`
	Name string `db:"name"`
}

func (<%= entityNameSingularPascalCase %>) GetTypeName() string {
	return "<%= entityNameSingularPascalCase %>"
}

func (self *<%= entityNameSingularPascalCase %>) GetId() int64 {
	return self.Id
}

func (self *<%= entityNameSingularPascalCase %>) Merge(update *<%= entityNameSingularPascalCase %>, fields []string) error {
  if contains(fields, "Name") {
    self.Name = update.Name
  }

  return nil
}
