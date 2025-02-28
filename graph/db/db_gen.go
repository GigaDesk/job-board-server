// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package db

import (
	"github.com/GigaDesk/eardrum-server/graph/model"
	"gorm.io/gorm"
)

type AutoGqlDB struct {
	Db    *gorm.DB
	Hooks map[string]any
}

// Create a new AutoGqlDB
func NewAutoGqlDB(db *gorm.DB) AutoGqlDB {
	return AutoGqlDB{
		Db:    db,
		Hooks: make(map[string]any),
	}
}

// execute Gorm AutoMigrate with all @SQL Graphql Types
func (db *AutoGqlDB) Init() error {
	return db.Db.AutoMigrate(&model.School{}, &model.UnverifiedSchool{}, &model.Job{}, &model.Student{}, &model.Admin{}, &model.UnverifiedAdmin{})
}
