package services

import (
	"go_angular_starter/model/entity"
	"go_angular_starter/model/persister"
)

func GenerateFullyQualifiedNames(db persister.DB, categoryId int64) error {
	categoryPersister := persister.NewCategoryPersister(db)

	category, err := categoryPersister.GetById(categoryId)

	if err != nil {
		return err
	}

	var parent *entity.Category = nil

	if category.ParentId != 0 {
		parent, err = categoryPersister.GetById(category.ParentId)
	}

	return generateFullyQualifiedNamesRecursive(db, category, parent)
}

func generateFullyQualifiedNamesRecursive(db persister.DB, category *entity.Category, parent *entity.Category) error {
	categoryPersister := persister.NewCategoryPersister(db)

	if parent != nil {
		category.FqName = parent.FqName + " > " + category.Name
	} else {
		category.FqName = category.Name
	}

	categoryPersister.Update(category)

	children, err := categoryPersister.GetChildren(category.Id)

	if err != nil {
		return err
	}

	for _, child := range children {
		err = generateFullyQualifiedNamesRecursive(db, child, category)

		if err != nil {
			return err
		}
	}

	return nil
}
