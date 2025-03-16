// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package db

import (
	"context"

	"github.com/GigaDesk/eardrum-server/graph/model"
	"gorm.io/gorm"
)

type QueryName string
type GetName string
type AddName string
type UpdateName string
type DeleteName string
type Many2ManyAddName string
type Many2ManyDeleteName string

const (
	GetAdmin                 GetName    = "GetAdmin"
	QueryAdmin               QueryName  = "QueryAdmin"
	AddAdmin                 AddName    = "AddAdmin"
	UpdateAdmin              UpdateName = "UpdateAdmin"
	DeleteAdmin              DeleteName = "DeleteAdmin"
	GetEmployee              GetName    = "GetEmployee"
	QueryEmployee            QueryName  = "QueryEmployee"
	AddEmployee              AddName    = "AddEmployee"
	UpdateEmployee           UpdateName = "UpdateEmployee"
	DeleteEmployee           DeleteName = "DeleteEmployee"
	GetEmployer              GetName    = "GetEmployer"
	QueryEmployer            QueryName  = "QueryEmployer"
	AddEmployer              AddName    = "AddEmployer"
	UpdateEmployer           UpdateName = "UpdateEmployer"
	DeleteEmployer           DeleteName = "DeleteEmployer"
	GetJob                   GetName    = "GetJob"
	QueryJob                 QueryName  = "QueryJob"
	AddJob                   AddName    = "AddJob"
	UpdateJob                UpdateName = "UpdateJob"
	DeleteJob                DeleteName = "DeleteJob"
	GetUnapprovedJob         GetName    = "GetUnapprovedJob"
	QueryUnapprovedJob       QueryName  = "QueryUnapprovedJob"
	AddUnapprovedJob         AddName    = "AddUnapprovedJob"
	UpdateUnapprovedJob      UpdateName = "UpdateUnapprovedJob"
	DeleteUnapprovedJob      DeleteName = "DeleteUnapprovedJob"
	GetUnverifiedAdmin       GetName    = "GetUnverifiedAdmin"
	QueryUnverifiedAdmin     QueryName  = "QueryUnverifiedAdmin"
	AddUnverifiedAdmin       AddName    = "AddUnverifiedAdmin"
	UpdateUnverifiedAdmin    UpdateName = "UpdateUnverifiedAdmin"
	DeleteUnverifiedAdmin    DeleteName = "DeleteUnverifiedAdmin"
	GetUnverifiedEmployee    GetName    = "GetUnverifiedEmployee"
	QueryUnverifiedEmployee  QueryName  = "QueryUnverifiedEmployee"
	AddUnverifiedEmployee    AddName    = "AddUnverifiedEmployee"
	UpdateUnverifiedEmployee UpdateName = "UpdateUnverifiedEmployee"
	DeleteUnverifiedEmployee DeleteName = "DeleteUnverifiedEmployee"
	GetUnverifiedEmployer    GetName    = "GetUnverifiedEmployer"
	QueryUnverifiedEmployer  QueryName  = "QueryUnverifiedEmployer"
	AddUnverifiedEmployer    AddName    = "AddUnverifiedEmployer"
	UpdateUnverifiedEmployer UpdateName = "UpdateUnverifiedEmployer"
	DeleteUnverifiedEmployer DeleteName = "DeleteUnverifiedEmployer"
)

// Modelhooks
type AutoGqlHookM interface {
	model.Admin | model.Employee | model.Employer | model.Job | model.UnapprovedJob | model.UnverifiedAdmin | model.UnverifiedEmployee | model.UnverifiedEmployer
}

// Filter Hooks
type AutoGqlHookF interface {
	model.AdminFiltersInput | model.EmployeeFiltersInput | model.EmployerFiltersInput | model.JobFiltersInput | model.UnapprovedJobFiltersInput | model.UnverifiedAdminFiltersInput | model.UnverifiedEmployeeFiltersInput | model.UnverifiedEmployerFiltersInput
}

// Many2Many Hooks
type AutoGqlHookM2M interface {
}

// Order Hooks
type AutoGqlHookQueryO interface {
	model.AdminOrder | model.EmployeeOrder | model.EmployerOrder | model.JobOrder | model.UnapprovedJobOrder | model.UnverifiedAdminOrder | model.UnverifiedEmployeeOrder | model.UnverifiedEmployerOrder
}

// Input Hooks
type AutoGqlHookI interface {
	model.AdminInput | model.EmployeeInput | model.EmployerInput | model.JobInput | model.UnapprovedJobInput | model.UnverifiedAdminInput | model.UnverifiedEmployeeInput | model.UnverifiedEmployerInput
}

// Update Hooks
type AutoGqlHookU interface {
	model.UpdateAdminInput | model.UpdateEmployeeInput | model.UpdateEmployerInput | model.UpdateJobInput | model.UpdateUnapprovedJobInput | model.UpdateUnverifiedAdminInput | model.UpdateUnverifiedEmployeeInput | model.UpdateUnverifiedEmployerInput
}

// Update Payload Hooks
type AutoGqlHookUP interface {
	model.UpdateAdminPayload | model.UpdateEmployeePayload | model.UpdateEmployerPayload | model.UpdateJobPayload | model.UpdateUnapprovedJobPayload | model.UpdateUnverifiedAdminPayload | model.UpdateUnverifiedEmployeePayload | model.UpdateUnverifiedEmployerPayload
}

// Delete Payload Hooks
type AutoGqlHookDP interface {
	model.DeleteAdminPayload | model.DeleteEmployeePayload | model.DeleteEmployerPayload | model.DeleteJobPayload | model.DeleteUnapprovedJobPayload | model.DeleteUnverifiedAdminPayload | model.DeleteUnverifiedEmployeePayload | model.DeleteUnverifiedEmployerPayload
}

// Add Payload Hooks
type AutoGqlHookAP interface {
	model.AddAdminPayload | model.AddEmployeePayload | model.AddEmployerPayload | model.AddJobPayload | model.AddUnapprovedJobPayload | model.AddUnverifiedAdminPayload | model.AddUnverifiedEmployeePayload | model.AddUnverifiedEmployerPayload
}

// Add a getHook
// useable for
//   - GetAdmin
//   - GetEmployee
//   - GetEmployer
//   - GetJob
//   - GetUnapprovedJob
//   - GetUnverifiedAdmin
//   - GetUnverifiedEmployee
//   - GetUnverifiedEmployer
func AddGetHook[T AutoGqlHookM, I any](db *AutoGqlDB, name GetName, implementation AutoGqlHookGet[T, I]) {
	db.Hooks[string(name)] = implementation
}

// Add a queryHook
// useable for
//   - QueryAdmin
//   - QueryEmployee
//   - QueryEmployer
//   - QueryJob
//   - QueryUnapprovedJob
//   - QueryUnverifiedAdmin
//   - QueryUnverifiedEmployee
//   - QueryUnverifiedEmployer
func AddQueryHook[M AutoGqlHookM, F AutoGqlHookF, O AutoGqlHookQueryO](db *AutoGqlDB, name QueryName, implementation AutoGqlHookQuery[M, F, O]) {
	db.Hooks[string(name)] = implementation
}

// Add a addHook
// useable for
//   - AddAdmin
//   - AddEmployee
//   - AddEmployer
//   - AddJob
//   - AddUnapprovedJob
//   - AddUnverifiedAdmin
//   - AddUnverifiedEmployee
//   - AddUnverifiedEmployer
func AddAddHook[M AutoGqlHookM, I AutoGqlHookI, AP AutoGqlHookAP](db *AutoGqlDB, name AddName, implementation AutoGqlHookAdd[M, I, AP]) {
	db.Hooks[string(name)] = implementation
}

// Add a updateHook
// useable for
//   - UpdateAdmin
//   - UpdateEmployee
//   - UpdateEmployer
//   - UpdateJob
//   - UpdateUnapprovedJob
//   - UpdateUnverifiedAdmin
//   - UpdateUnverifiedEmployee
//   - UpdateUnverifiedEmployer
func AddUpdateHook[M AutoGqlHookM, U AutoGqlHookU, UP AutoGqlHookUP](db *AutoGqlDB, name UpdateName, implementation AutoGqlHookUpdate[U, UP]) {
	db.Hooks[string(name)] = implementation
}

// Add a Many2Many Add Hook
// useable for
func AddMany2ManyAddHook[U AutoGqlHookM2M, UP AutoGqlHookUP](db *AutoGqlDB, name Many2ManyAddName, implementation AutoGqlHookMany2ManyAdd[U, UP]) {
	db.Hooks[string(name)] = implementation
}

// Add a Many2Many Delete Hook
// useable for
func AddMany2ManyDeleteHook[U AutoGqlHookM2M, DP AutoGqlHookDP](db *AutoGqlDB, name Many2ManyDeleteName, implementation AutoGqlHookMany2ManyDelete[U, DP]) {
	db.Hooks[string(name)] = implementation
}

// Add a updateHook
// useable for
//   - DeleteAdmin
//   - DeleteEmployee
//   - DeleteEmployer
//   - DeleteJob
//   - DeleteUnapprovedJob
//   - DeleteUnverifiedAdmin
//   - DeleteUnverifiedEmployee
//   - DeleteUnverifiedEmployer
func AddDeleteHook[F AutoGqlHookF, DP AutoGqlHookDP](db *AutoGqlDB, name DeleteName, implementation AutoGqlHookDelete[F, DP]) {
	db.Hooks[string(name)] = implementation
}

// Interface description of a getHook
// Simple you can use DefaultGetHook and only implement the hooks you need:
//
//	type MyGetHook struct {
//	   DefaultGetHook[model.Todo, model.TodoInput, model.AddTodoPayload]
//	}
//	func (m MyGetHook) BeforeCallDb(ctx context.Context, db *gorm.DB, data []model.Todo) (*gorm.DB, []model.Todo, error) {
//	   //do some stuff
//	   return db, data, nil
//	}
type AutoGqlHookGet[obj AutoGqlHookM, identifier any] interface {
	Received(ctx context.Context, dbHelper *AutoGqlDB, id ...identifier) (*gorm.DB, error) // Direct after Resolver is call
	BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error)                       // Direct before call Database
	AfterCallDb(ctx context.Context, data *obj) (*obj, error)                              // After database call with resultset from database
	BeforeReturn(ctx context.Context, data *obj, db *gorm.DB) (*obj, error)                // short before return the data
}

// Default get hook implementation
// Simple you can use and only implement the hooks you need:
//
//	type MyGetHook struct {
//	   DefaultGetHook[model.Todo, int64]
//	}
//	func (m MyGetHook) BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, []model.Todo, error) {
//	   //do some stuff
//	   return db, data, nil
//	}
type DefaultGetHook[obj AutoGqlHookM, identifier any] struct{}

// Direct after Resolver is call
func (d DefaultGetHook[obj, identifier]) Received(ctx context.Context, dbHelper *AutoGqlDB, id ...identifier) (*gorm.DB, error) {
	return dbHelper.Db, nil
}

// Direct before call Database
func (d DefaultGetHook[obj, identifier]) BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error) {
	return db, nil
}

// After database call with resultset from database
func (d DefaultGetHook[obj, identifier]) AfterCallDb(ctx context.Context, data *obj) (*obj, error) {
	return data, nil
}

// short before return the data
func (d DefaultGetHook[obj, identifier]) BeforeReturn(ctx context.Context, data *obj, db *gorm.DB) (*obj, error) {
	return data, nil
}

// Interface description of a query Hook
// Simple you can use DefaultQueryHook and only implement the hooks you need:
//
//	type MyQueryHook struct {
//	   DefaultQueryHook[model.Todo, model.TodoFiltersInput, model.TodoOrder]
//	}
//	func (m MyQueryHook) BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, []model.Todo, error) {
//	   //do some stuff
//	   return db, nil
//	}
type AutoGqlHookQuery[obj AutoGqlHookM, filter AutoGqlHookF, order AutoGqlHookQueryO] interface {
	Received(ctx context.Context, dbHelper *AutoGqlDB, filter *filter, order *order, first, offset *int) (*gorm.DB, *filter, *order, *int, *int, error) // Direct after Resolver is call
	BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error)                                                                                    // Direct before call Database
	AfterCallDb(ctx context.Context, data []*obj) ([]*obj, error)                                                                                       // After database call with resultset from database
	BeforeReturn(ctx context.Context, data []*obj, db *gorm.DB) ([]*obj, error)                                                                         // short before return the data
}

// Default query hook implementation
// Simple you can use DefaultQueryHook and only implement the hooks you need:
//
//	type MyQueryHook struct {
//	   DefaultQueryHook[model.Todo, model.TodoFiltersInput, model.TodoOrder]
//	}
//	func (m MyQueryHook) BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, []model.Todo, error) {
//	   //do some stuff
//	   return db, nil
//	}
type DefaultQueryHook[obj AutoGqlHookM, filter AutoGqlHookF, order AutoGqlHookQueryO] struct{}

// Direct after Resolver is call
func (d DefaultQueryHook[obj, filterType, orderType]) Received(ctx context.Context, dbHelper *AutoGqlDB, filter *filterType, order *orderType, first, offset *int) (*gorm.DB, *filterType, *orderType, *int, *int, error) {
	return dbHelper.Db, filter, order, first, offset, nil
}

// Direct before call Database
func (d DefaultQueryHook[obj, filter, order]) BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error) {
	return db, nil
}

// After database call with resultset from database
func (d DefaultQueryHook[obj, filter, order]) AfterCallDb(ctx context.Context, data []*obj) ([]*obj, error) {
	return data, nil
}

// short before return the data
func (d DefaultQueryHook[obj, filter, order]) BeforeReturn(ctx context.Context, data []*obj, db *gorm.DB) ([]*obj, error) {
	return data, nil
}

// Interface description of a add Hook
// Simple you can use DefaultAddHook and only implement the hooks you need:
//
//	type MyAddHook struct {
//	   DefaultAddHook[model.Todo, model.TodoInput, model.AddTodoPayload]
//	}
//	func (m MyAddHook) BeforeCallDb(ctx context.Context, db *gorm.DB, data []model.Todo) (*gorm.DB, []model.Todo, error) {
//	   //do some stuff
//	   return db, data, nil
//	}
type AutoGqlHookAdd[obj AutoGqlHookM, input AutoGqlHookI, res AutoGqlHookAP] interface {
	Received(ctx context.Context, dbHelper *AutoGqlDB, input []*input) (*gorm.DB, []*input, error) // Direct after Resolver is call
	BeforeCallDb(ctx context.Context, db *gorm.DB, data []obj) (*gorm.DB, []obj, error)            // Direct before call Database
	BeforeReturn(ctx context.Context, db *gorm.DB, data []obj, res *res) (*res, error)             // After database call with resultset from database
}

// Default add hook implementation
// Simple you can use DefaultAddHook and only implement the hooks you need:
//
//	type MyAddHook struct {
//	   DefaultAddHook[model.Todo, model.TodoInput, model.AddTodoPayload]
//	}
//	func (m MyAddHook) BeforeCallDb(ctx context.Context, db *gorm.DB, data []model.Todo) (*gorm.DB, []model.Todo, error) {
//	   //do some stuff
//	   return db, data, nil
//	}
type DefaultAddHook[obj AutoGqlHookM, input AutoGqlHookI, res AutoGqlHookAP] struct{}

// Direct after Resolver is call
func (d DefaultAddHook[obj, inputType, resType]) Received(ctx context.Context, dbHelper *AutoGqlDB, input []*inputType) (*gorm.DB, []*inputType, error) {
	return dbHelper.Db, input, nil
}

// Direct before call Database
func (d DefaultAddHook[obj, inputType, resType]) BeforeCallDb(ctx context.Context, db *gorm.DB, data []obj) (*gorm.DB, []obj, error) {
	return db, data, nil
}

// After database call with resultset from database
func (d DefaultAddHook[obj, inputType, resType]) BeforeReturn(ctx context.Context, db *gorm.DB, data []obj, res *resType) (*resType, error) {
	return res, nil
}

// Interface description of a update Hook
// Simple you can use DefaultUpdateHook and only implement the hooks you need:
//
//	type MyUpdateHook struct {
//	   DefaultUpdateHook[model.TodoInput, model.UpdateTodoPayload]
//	}
//	func (m MyUpdateHook) BeforeCallDb(ctx context.Context, db *gorm.DB, data map[string]interface{}) (*gorm.DB,  map[string]interface{}, error) {
//	   //do some stuff
//	   return db, data, nil
//	}
type AutoGqlHookUpdate[input AutoGqlHookU, res AutoGqlHookUP] interface {
	Received(ctx context.Context, dbHelper *AutoGqlDB, input *input) (*gorm.DB, input, error)                             // Direct after Resolver is call
	BeforeCallDb(ctx context.Context, db *gorm.DB, data map[string]interface{}) (*gorm.DB, map[string]interface{}, error) // Direct before call Database
	BeforeReturn(ctx context.Context, db *gorm.DB, res *res) (*res, error)                                                // After database call with resultset from database
}

// Default update hook implementation
// Simple you can use DefaultUpdateHook and only implement the hooks you need:
//
//	type MyUpdateHook struct {
//	   DefaultUpdateHook[model.TodoInput, model.UpdateTodoPayload]
//	}
//	func (m MyUpdateHook) BeforeCallDb(ctx context.Context, db *gorm.DB, data map[string]interface{}) (*gorm.DB,  map[string]interface{}, error) {
//	   //do some stuff
//	   return db, data, nil
//	}
type DefaultUpdateHook[input AutoGqlHookU, res AutoGqlHookUP] struct{}

// Direct after Resolver is call
func (d DefaultUpdateHook[inputType, resType]) Received(ctx context.Context, dbHelper *AutoGqlDB, input *inputType) (*gorm.DB, inputType, error) {
	return dbHelper.Db, *input, nil
}

// Direct before call Database
func (d DefaultUpdateHook[inputType, resType]) BeforeCallDb(ctx context.Context, db *gorm.DB, data map[string]interface{}) (*gorm.DB, map[string]interface{}, error) {
	return db, data, nil
}

// After database call with resultset from database
func (d DefaultUpdateHook[inputType, resType]) BeforeReturn(ctx context.Context, db *gorm.DB, res *resType) (*resType, error) {
	return res, nil
}

// Interface description of a many2many Hook
// Simple you can use DefaultMany2ManyHook and only implement the hooks you need:
//
//	type MyM2mHook struct {
//	   DefaultMany2ManyHook[model.UserRef2TodosInput, model.UpdateTodoPayload]
//	}
//	func (m MyM2mHook) BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error) {
//	   //do some stuff
//	   return db, nil
//	}
type AutoGqlHookMany2ManyAdd[input AutoGqlHookM2M, res AutoGqlHookUP] interface {
	Received(ctx context.Context, dbHelper *AutoGqlDB, input *input) (*gorm.DB, input, error) // Direct after Resolver is call
	BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error)                          // Direct before call Database
	BeforeReturn(ctx context.Context, db *gorm.DB, res *res) (*res, error)                    // After database call with resultset from database
}

// Default many2many hook implementation
// Simple you can use DefaultMany2ManyAddHook and only implement the hooks you need:
//
//	type MyM2mHook struct {
//	   DefaultMany2ManyAddHook[model.UserRef2TodosInput, model.UpdateTodoPayload]
//	}
//	func (m MyM2mAddHook) BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error) {
//	   //do some stuff
//	   return db, nil
//	}
type DefaultMany2ManyAddHook[input AutoGqlHookM2M, res AutoGqlHookUP] struct{}

// Direct after Resolver is call
func (d DefaultMany2ManyAddHook[inputType, resType]) Received(ctx context.Context, dbHelper *AutoGqlDB, input *inputType) (*gorm.DB, inputType, error) {
	return dbHelper.Db, *input, nil
}

// Direct before call Database
func (d DefaultMany2ManyAddHook[inputType, resType]) BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error) {
	return db, nil
}

// After database call with resultset from database
func (d DefaultMany2ManyAddHook[inputType, resType]) BeforeReturn(ctx context.Context, db *gorm.DB, res *resType) (*resType, error) {
	return res, nil
}

// Interface description of a many2many delete Hook
// Simple you can use DefaultMany2ManyHook and only implement the hooks you need:
//
//	type MyM2mDeleteHook struct {
//	   DefaultMany2ManyDeleteHook[model.UserRef2TodosInput, model.UpdateTodoPayload]
//	}
//	func (m MyM2mDeleteHook) BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error) {
//	   //do some stuff
//	   return db, nil
//	}
type AutoGqlHookMany2ManyDelete[input AutoGqlHookM2M, res AutoGqlHookDP] interface {
	Received(ctx context.Context, dbHelper *AutoGqlDB, input *input) (*gorm.DB, input, error) // Direct after Resolver is call
	BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error)                          // Direct before call Database
	BeforeReturn(ctx context.Context, db *gorm.DB, res *res) (*res, error)                    // After database call with resultset from database
}

// Default many2many delete hook implementation
// Simple you can use DefaultMany2ManyDeleteHook and only implement the hooks you need:
//
//	type MyM2mDeleteHook struct {
//	   DefaultMany2ManyDeleteHook[model.UserRef2TodosInput, model.UpdateTodoPayload]
//	}
//	func (m MyM2mDeleteHook) BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error) {
//	   //do some stuff
//	   return db, nil
//	}
type DefaultMany2ManyDeleteHook[input AutoGqlHookM2M, res AutoGqlHookDP] struct{}

// Direct after Resolver is call
func (d DefaultMany2ManyDeleteHook[inputType, resType]) Received(ctx context.Context, dbHelper *AutoGqlDB, input *inputType) (*gorm.DB, inputType, error) {
	return dbHelper.Db, *input, nil
}

// Direct before call Database
func (d DefaultMany2ManyDeleteHook[inputType, resType]) BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error) {
	return db, nil
}

// After database call with resultset from database
func (d DefaultMany2ManyDeleteHook[inputType, resType]) BeforeReturn(ctx context.Context, db *gorm.DB, res *resType) (*resType, error) {
	return res, nil
}

// Interface description of a delete Hook
// Simple you can use DefaultDeleteHook and only implement the hooks you need:
//
//	type MyDeleteHook struct {
//	   DefaultDeleteHook[model.TodoFiltersInput, model.DeleteTodoPayload]
//	}
//	func (m MyDeleteHook) BeforeCallDb(ctx context.Context, db *gorm.DB, input model.TodoFiltersInput) (*gorm.DB, model.TodoFiltersInput, error) {
//	   //do some stuff
//	   return db, input, nil
//	}
type AutoGqlHookDelete[input AutoGqlHookF, res AutoGqlHookDP] interface {
	Received(ctx context.Context, dbHelper *AutoGqlDB, input *input) (*gorm.DB, input, error) // Direct after Resolver is call
	BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error)                          // Direct before call Database
	BeforeReturn(ctx context.Context, db *gorm.DB, res *res) (*res, error)                    // After database call with resultset from database
}

// Default delete hook implementation
// Simple you can use DefaultDeleteHook and only implement the hooks you need:
//
//	type MyM2mHook struct {
//	   DefaultDeleteHook[model.TodoFiltersInput, model.DeleteTodoPayload]
//	}
//	func (m MyM2mHook) BeforeCallDb(ctx context.Context, db *gorm.DB, input model.TodoFiltersInput) (*gorm.DB, model.TodoFiltersInput, error) {
//	   //do some stuff
//	   return db, input, nil
//	}
type DefaultDeleteHook[input AutoGqlHookF, res AutoGqlHookDP] struct{}

// Direct after Resolver is call
func (d DefaultDeleteHook[inputType, resType]) Received(ctx context.Context, dbHelper *AutoGqlDB, input *inputType) (*gorm.DB, inputType, error) {
	return dbHelper.Db, *input, nil
}

// Direct before call Database
func (d DefaultDeleteHook[inputType, resType]) BeforeCallDb(ctx context.Context, db *gorm.DB) (*gorm.DB, error) {
	return db, nil
}

// After database call with resultset from database
func (d DefaultDeleteHook[inputType, resType]) BeforeReturn(ctx context.Context, db *gorm.DB, res *resType) (*resType, error) {
	return res, nil
}
