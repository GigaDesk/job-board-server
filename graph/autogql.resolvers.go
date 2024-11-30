// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graph

import (
	"context"
	"fmt"
	"strings"

	"github.com/GigaDesk/eardrum-server/graph/db"
	"github.com/GigaDesk/eardrum-server/graph/model"
	"github.com/fasibio/autogql/runtimehelper"
	"github.com/huandu/xstrings"
	"gorm.io/gorm/clause"
)

// GetSchool is the resolver for the getSchool field.
func (r *queryResolver) GetSchool(ctx context.Context, id int) (*model.School, error) {
	v, okHook := r.Sql.Hooks[string(db.GetSchool)].(db.AutoGqlHookGet[model.School, int])
	db := r.Sql.Db
	if okHook {
		var err error
		db, err = v.Received(ctx, r.Sql, id)
		if err != nil {
			return nil, err
		}
	}
	db = runtimehelper.GetPreloadSelection(ctx, db, runtimehelper.GetPreloadsMap(ctx, "School"))
	if okHook {
		var err error
		db, err = v.BeforeCallDb(ctx, db)
		if err != nil {
			return nil, err
		}
	}
	var res model.School
	tableName := r.Sql.Db.Config.NamingStrategy.TableName("School")
	db = db.First(&res, tableName+".id = ?", id)
	if okHook {
		r, err := v.AfterCallDb(ctx, &res)
		if err != nil {
			return nil, err
		}
		res = *r
		r, err = v.BeforeReturn(ctx, &res, db)
		if err != nil {
			return nil, err
		}
		res = *r
	}
	return &res, db.Error
}

// QuerySchool is the resolver for the querySchool field.
func (r *queryResolver) QuerySchool(ctx context.Context, filter *model.SchoolFiltersInput, order *model.SchoolOrder, first *int, offset *int, group []model.SchoolGroup) (*model.SchoolQueryResult, error) {
	v, okHook := r.Sql.Hooks[string(db.QuerySchool)].(db.AutoGqlHookQuery[model.School, model.SchoolFiltersInput, model.SchoolOrder])
	db := r.Sql.Db
	if okHook {
		var err error
		db, filter, order, first, offset, err = v.Received(ctx, r.Sql, filter, order, first, offset)
		if err != nil {
			return nil, err
		}
	}
	var res []*model.School
	tableName := r.Sql.Db.Config.NamingStrategy.TableName("School")
	preloadSubTables := runtimehelper.GetPreloadsMap(ctx, "data").SubTables
	if len(preloadSubTables) > 0 {
		db = runtimehelper.GetPreloadSelection(ctx, db, preloadSubTables[0])
	}
	if filter != nil {
		blackList := make(map[string]struct{})
		sql, arguments := runtimehelper.CombineSimpleQuery(filter.ExtendsDatabaseQuery(db, fmt.Sprintf("%[1]s%[2]s%[1]s", runtimehelper.GetQuoteChar(db), tableName), false, blackList), "AND")
		db.Where(sql, arguments...)
	}
	if okHook {
		var err error
		db, err = v.BeforeCallDb(ctx, db)
		if err != nil {
			return nil, err
		}
	}
	if order != nil {
		if order.Asc != nil {
			db = db.Order(fmt.Sprintf("%[1]s%[2]s%[1]s.%[1]s%[3]s%[1]s asc", runtimehelper.GetQuoteChar(db), tableName, order.Asc))
		}
		if order.Desc != nil {
			db = db.Order(fmt.Sprintf("%[1]s%[2]s%[1]s.%[1]s%[3]s%[1]s desc", runtimehelper.GetQuoteChar(db), tableName, order.Desc))
		}
	}
	var total int64
	db.Model(res).Count(&total)
	if first != nil {
		db = db.Limit(*first)
	}
	if offset != nil {
		db = db.Offset(*offset)
	}
	if len(group) > 0 {
		garr := make([]string, len(group))
		for i, g := range group {
			garr[i] = fmt.Sprintf("%s.%s", tableName, xstrings.ToSnakeCase(string(g)))
		}
		db = db.Group(strings.Join(garr, ","))
	}
	db = db.Find(&res)
	if okHook {
		var err error
		res, err = v.AfterCallDb(ctx, res)
		if err != nil {
			return nil, err
		}
		res, err = v.BeforeReturn(ctx, res, db)
		if err != nil {
			return nil, err
		}
	}
	return &model.SchoolQueryResult{
		Data:       res,
		Count:      len(res),
		TotalCount: int(total),
	}, db.Error
}
func (r *Resolver) AddSchoolPayload() AddSchoolPayloadResolver {
	return &schoolPayloadResolver[*model.AddSchoolPayload]{r}
}
func (r *Resolver) DeleteSchoolPayload() DeleteSchoolPayloadResolver {
	return &schoolPayloadResolver[*model.DeleteSchoolPayload]{r}
}
func (r *Resolver) UpdateSchoolPayload() UpdateSchoolPayloadResolver {
	return &schoolPayloadResolver[*model.UpdateSchoolPayload]{r}
}

type schoolPayload interface {
	*model.AddSchoolPayload | *model.DeleteSchoolPayload | *model.UpdateSchoolPayload
}

type schoolPayloadResolver[T schoolPayload] struct {
	*Resolver
}

func (r *schoolPayloadResolver[T]) School(ctx context.Context, obj T, filter *model.SchoolFiltersInput, order *model.SchoolOrder, first *int, offset *int, group []model.SchoolGroup) (*model.SchoolQueryResult, error) {
	q := r.Query().(*queryResolver)
	return q.QuerySchool(ctx, filter, order, first, offset, group)
}

// AddSchool is the resolver for the addSchool field.
func (r *mutationResolver) AddSchool(ctx context.Context, input []*model.SchoolInput) (*model.AddSchoolPayload, error) {
	v, okHook := r.Sql.Hooks[string(db.AddSchool)].(db.AutoGqlHookAdd[model.School, model.SchoolInput, model.AddSchoolPayload])
	res := &model.AddSchoolPayload{}
	db := r.Sql.Db
	if okHook {
		var err error
		db, input, err = v.Received(ctx, r.Sql, input)
		if err != nil {
			return nil, err
		}
	}
	obj := make([]model.School, len(input))
	for i, v := range input {
		obj[i] = v.MergeToType()
	}
	db = db.Omit(clause.Associations)
	if okHook {
		var err error
		db, obj, err = v.BeforeCallDb(ctx, db, obj)
		if err != nil {
			return nil, err
		}
	}
	db = db.Create(&obj)
	affectedRes := make([]*model.School, len(obj))
	for i, v := range obj {
		tmp := v
		affectedRes[i] = &tmp
	}
	res.Affected = affectedRes
	if okHook {
		var err error
		res, err = v.BeforeReturn(ctx, db, obj, res)
		if err != nil {
			return nil, err
		}
	}
	return res, db.Error
}

// UpdateSchool is the resolver for the updateSchool field.
func (r *mutationResolver) UpdateSchool(ctx context.Context, input model.UpdateSchoolInput) (*model.UpdateSchoolPayload, error) {
	v, okHook := r.Sql.Hooks[string(db.UpdateSchool)].(db.AutoGqlHookUpdate[model.UpdateSchoolInput, model.UpdateSchoolPayload])
	db := r.Sql.Db
	if okHook {
		var err error
		db, input, err = v.Received(ctx, r.Sql, &input)
		if err != nil {
			return nil, err
		}
	}
	tableName := r.Sql.Db.Config.NamingStrategy.TableName("School")
	blackList := make(map[string]struct{})
	queryDb := db.Select(tableName + ".id")
	sql, arguments := runtimehelper.CombineSimpleQuery(input.Filter.ExtendsDatabaseQuery(queryDb, fmt.Sprintf("%[1]s%[2]s%[1]s", runtimehelper.GetQuoteChar(db), tableName), false, blackList), "AND")
	obj := model.School{}
	queryDb = queryDb.Model(&obj).Where(sql, arguments...)
	var toChange []model.School
	queryDb.Find(&toChange)
	update := input.Set.MergeToType()
	if okHook {
		var err error
		db, update, err = v.BeforeCallDb(ctx, db, update)
		if err != nil {
			return nil, err
		}
	}
	ids := make([]int, len(toChange))
	for i, one := range toChange {
		ids[i] = one.ID
	}
	db = db.Model(&obj).Where("id IN ?", ids).Updates(update)
	affectedRes := make([]*model.School, 0)
	subTables := runtimehelper.GetPreloadsMap(ctx, "affected").SubTables
	if len(subTables) > 0 {
		if preloadMap := subTables[0]; len(preloadMap.Fields) > 0 {
			affectedDb := runtimehelper.GetPreloadSelection(ctx, db, preloadMap)
			affectedDb = affectedDb.Model(&obj)
			affectedDb.Find(&affectedRes)
		}
	}

	res := &model.UpdateSchoolPayload{
		Count:    int(db.RowsAffected),
		Affected: affectedRes,
	}
	if okHook {
		var err error
		res, err = v.BeforeReturn(ctx, db, res)
		if err != nil {
			return nil, err
		}
	}
	return res, db.Error
}

// DeleteSchool is the resolver for the deleteSchool field.
func (r *mutationResolver) DeleteSchool(ctx context.Context, filter model.SchoolFiltersInput) (*model.DeleteSchoolPayload, error) {
	v, okHook := r.Sql.Hooks[string(db.DeleteSchool)].(db.AutoGqlHookDelete[model.SchoolFiltersInput, model.DeleteSchoolPayload])
	db := r.Sql.Db
	if okHook {
		var err error
		db, filter, err = v.Received(ctx, r.Sql, &filter)
		if err != nil {
			return nil, err
		}
	}
	tableName := r.Sql.Db.Config.NamingStrategy.TableName("School")
	blackList := make(map[string]struct{})
	queryDb := db.Select(tableName + ".id")
	sql, arguments := runtimehelper.CombineSimpleQuery(filter.ExtendsDatabaseQuery(queryDb, fmt.Sprintf("%[1]s%[2]s%[1]s", runtimehelper.GetQuoteChar(db), tableName), false, blackList), "AND")
	obj := model.School{}
	queryDb = queryDb.Model(&obj).Where(sql, arguments...)
	var toChange []model.School
	queryDb.Find(&toChange)
	if okHook {
		var err error
		db, err = v.BeforeCallDb(ctx, db)
		if err != nil {
			return nil, err
		}
	}
	ids := make([]int, len(toChange))
	for i, one := range toChange {
		ids[i] = one.ID
	}
	db = db.Model(&obj).Where("id IN ?", ids).Delete(&obj)
	msg := fmt.Sprintf("%d rows deleted", db.RowsAffected)
	res := &model.DeleteSchoolPayload{
		Count: int(db.RowsAffected),
		Msg:   &msg,
	}
	if okHook {
		var err error
		res, err = v.BeforeReturn(ctx, db, res)
		if err != nil {
			return nil, err
		}
	}
	return res, db.Error
}

// GetStudent is the resolver for the getStudent field.
func (r *queryResolver) GetStudent(ctx context.Context, id int) (*model.Student, error) {
	v, okHook := r.Sql.Hooks[string(db.GetStudent)].(db.AutoGqlHookGet[model.Student, int])
	db := r.Sql.Db
	if okHook {
		var err error
		db, err = v.Received(ctx, r.Sql, id)
		if err != nil {
			return nil, err
		}
	}
	db = runtimehelper.GetPreloadSelection(ctx, db, runtimehelper.GetPreloadsMap(ctx, "Student"))
	if okHook {
		var err error
		db, err = v.BeforeCallDb(ctx, db)
		if err != nil {
			return nil, err
		}
	}
	var res model.Student
	tableName := r.Sql.Db.Config.NamingStrategy.TableName("Student")
	db = db.First(&res, tableName+".id = ?", id)
	if okHook {
		r, err := v.AfterCallDb(ctx, &res)
		if err != nil {
			return nil, err
		}
		res = *r
		r, err = v.BeforeReturn(ctx, &res, db)
		if err != nil {
			return nil, err
		}
		res = *r
	}
	return &res, db.Error
}

// QueryStudent is the resolver for the queryStudent field.
func (r *queryResolver) QueryStudent(ctx context.Context, filter *model.StudentFiltersInput, order *model.StudentOrder, first *int, offset *int, group []model.StudentGroup) (*model.StudentQueryResult, error) {
	v, okHook := r.Sql.Hooks[string(db.QueryStudent)].(db.AutoGqlHookQuery[model.Student, model.StudentFiltersInput, model.StudentOrder])
	db := r.Sql.Db
	if okHook {
		var err error
		db, filter, order, first, offset, err = v.Received(ctx, r.Sql, filter, order, first, offset)
		if err != nil {
			return nil, err
		}
	}
	var res []*model.Student
	tableName := r.Sql.Db.Config.NamingStrategy.TableName("Student")
	preloadSubTables := runtimehelper.GetPreloadsMap(ctx, "data").SubTables
	if len(preloadSubTables) > 0 {
		db = runtimehelper.GetPreloadSelection(ctx, db, preloadSubTables[0])
	}
	if filter != nil {
		blackList := make(map[string]struct{})
		sql, arguments := runtimehelper.CombineSimpleQuery(filter.ExtendsDatabaseQuery(db, fmt.Sprintf("%[1]s%[2]s%[1]s", runtimehelper.GetQuoteChar(db), tableName), false, blackList), "AND")
		db.Where(sql, arguments...)
	}
	if okHook {
		var err error
		db, err = v.BeforeCallDb(ctx, db)
		if err != nil {
			return nil, err
		}
	}
	if order != nil {
		if order.Asc != nil {
			db = db.Order(fmt.Sprintf("%[1]s%[2]s%[1]s.%[1]s%[3]s%[1]s asc", runtimehelper.GetQuoteChar(db), tableName, order.Asc))
		}
		if order.Desc != nil {
			db = db.Order(fmt.Sprintf("%[1]s%[2]s%[1]s.%[1]s%[3]s%[1]s desc", runtimehelper.GetQuoteChar(db), tableName, order.Desc))
		}
	}
	var total int64
	db.Model(res).Count(&total)
	if first != nil {
		db = db.Limit(*first)
	}
	if offset != nil {
		db = db.Offset(*offset)
	}
	if len(group) > 0 {
		garr := make([]string, len(group))
		for i, g := range group {
			garr[i] = fmt.Sprintf("%s.%s", tableName, xstrings.ToSnakeCase(string(g)))
		}
		db = db.Group(strings.Join(garr, ","))
	}
	db = db.Find(&res)
	if okHook {
		var err error
		res, err = v.AfterCallDb(ctx, res)
		if err != nil {
			return nil, err
		}
		res, err = v.BeforeReturn(ctx, res, db)
		if err != nil {
			return nil, err
		}
	}
	return &model.StudentQueryResult{
		Data:       res,
		Count:      len(res),
		TotalCount: int(total),
	}, db.Error
}
func (r *Resolver) AddStudentPayload() AddStudentPayloadResolver {
	return &studentPayloadResolver[*model.AddStudentPayload]{r}
}
func (r *Resolver) DeleteStudentPayload() DeleteStudentPayloadResolver {
	return &studentPayloadResolver[*model.DeleteStudentPayload]{r}
}
func (r *Resolver) UpdateStudentPayload() UpdateStudentPayloadResolver {
	return &studentPayloadResolver[*model.UpdateStudentPayload]{r}
}

type studentPayload interface {
	*model.AddStudentPayload | *model.DeleteStudentPayload | *model.UpdateStudentPayload
}

type studentPayloadResolver[T studentPayload] struct {
	*Resolver
}

func (r *studentPayloadResolver[T]) Student(ctx context.Context, obj T, filter *model.StudentFiltersInput, order *model.StudentOrder, first *int, offset *int, group []model.StudentGroup) (*model.StudentQueryResult, error) {
	q := r.Query().(*queryResolver)
	return q.QueryStudent(ctx, filter, order, first, offset, group)
}

// AddStudent is the resolver for the addStudent field.
func (r *mutationResolver) AddStudent(ctx context.Context, input []*model.StudentInput) (*model.AddStudentPayload, error) {
	v, okHook := r.Sql.Hooks[string(db.AddStudent)].(db.AutoGqlHookAdd[model.Student, model.StudentInput, model.AddStudentPayload])
	res := &model.AddStudentPayload{}
	db := r.Sql.Db
	if okHook {
		var err error
		db, input, err = v.Received(ctx, r.Sql, input)
		if err != nil {
			return nil, err
		}
	}
	obj := make([]model.Student, len(input))
	for i, v := range input {
		obj[i] = v.MergeToType()
	}
	db = db.Omit(clause.Associations)
	if okHook {
		var err error
		db, obj, err = v.BeforeCallDb(ctx, db, obj)
		if err != nil {
			return nil, err
		}
	}
	db = db.Create(&obj)
	affectedRes := make([]*model.Student, len(obj))
	for i, v := range obj {
		tmp := v
		affectedRes[i] = &tmp
	}
	res.Affected = affectedRes
	if okHook {
		var err error
		res, err = v.BeforeReturn(ctx, db, obj, res)
		if err != nil {
			return nil, err
		}
	}
	return res, db.Error
}

// UpdateStudent is the resolver for the updateStudent field.
func (r *mutationResolver) UpdateStudent(ctx context.Context, input model.UpdateStudentInput) (*model.UpdateStudentPayload, error) {
	v, okHook := r.Sql.Hooks[string(db.UpdateStudent)].(db.AutoGqlHookUpdate[model.UpdateStudentInput, model.UpdateStudentPayload])
	db := r.Sql.Db
	if okHook {
		var err error
		db, input, err = v.Received(ctx, r.Sql, &input)
		if err != nil {
			return nil, err
		}
	}
	tableName := r.Sql.Db.Config.NamingStrategy.TableName("Student")
	blackList := make(map[string]struct{})
	queryDb := db.Select(tableName + ".id")
	sql, arguments := runtimehelper.CombineSimpleQuery(input.Filter.ExtendsDatabaseQuery(queryDb, fmt.Sprintf("%[1]s%[2]s%[1]s", runtimehelper.GetQuoteChar(db), tableName), false, blackList), "AND")
	obj := model.Student{}
	queryDb = queryDb.Model(&obj).Where(sql, arguments...)
	var toChange []model.Student
	queryDb.Find(&toChange)
	update := input.Set.MergeToType()
	if okHook {
		var err error
		db, update, err = v.BeforeCallDb(ctx, db, update)
		if err != nil {
			return nil, err
		}
	}
	ids := make([]int, len(toChange))
	for i, one := range toChange {
		ids[i] = one.ID
	}
	db = db.Model(&obj).Where("id IN ?", ids).Updates(update)
	affectedRes := make([]*model.Student, 0)
	subTables := runtimehelper.GetPreloadsMap(ctx, "affected").SubTables
	if len(subTables) > 0 {
		if preloadMap := subTables[0]; len(preloadMap.Fields) > 0 {
			affectedDb := runtimehelper.GetPreloadSelection(ctx, db, preloadMap)
			affectedDb = affectedDb.Model(&obj)
			affectedDb.Find(&affectedRes)
		}
	}

	res := &model.UpdateStudentPayload{
		Count:    int(db.RowsAffected),
		Affected: affectedRes,
	}
	if okHook {
		var err error
		res, err = v.BeforeReturn(ctx, db, res)
		if err != nil {
			return nil, err
		}
	}
	return res, db.Error
}

// DeleteStudent is the resolver for the deleteStudent field.
func (r *mutationResolver) DeleteStudent(ctx context.Context, filter model.StudentFiltersInput) (*model.DeleteStudentPayload, error) {
	v, okHook := r.Sql.Hooks[string(db.DeleteStudent)].(db.AutoGqlHookDelete[model.StudentFiltersInput, model.DeleteStudentPayload])
	db := r.Sql.Db
	if okHook {
		var err error
		db, filter, err = v.Received(ctx, r.Sql, &filter)
		if err != nil {
			return nil, err
		}
	}
	tableName := r.Sql.Db.Config.NamingStrategy.TableName("Student")
	blackList := make(map[string]struct{})
	queryDb := db.Select(tableName + ".id")
	sql, arguments := runtimehelper.CombineSimpleQuery(filter.ExtendsDatabaseQuery(queryDb, fmt.Sprintf("%[1]s%[2]s%[1]s", runtimehelper.GetQuoteChar(db), tableName), false, blackList), "AND")
	obj := model.Student{}
	queryDb = queryDb.Model(&obj).Where(sql, arguments...)
	var toChange []model.Student
	queryDb.Find(&toChange)
	if okHook {
		var err error
		db, err = v.BeforeCallDb(ctx, db)
		if err != nil {
			return nil, err
		}
	}
	ids := make([]int, len(toChange))
	for i, one := range toChange {
		ids[i] = one.ID
	}
	db = db.Model(&obj).Where("id IN ?", ids).Delete(&obj)
	msg := fmt.Sprintf("%d rows deleted", db.RowsAffected)
	res := &model.DeleteStudentPayload{
		Count: int(db.RowsAffected),
		Msg:   &msg,
	}
	if okHook {
		var err error
		res, err = v.BeforeReturn(ctx, db, res)
		if err != nil {
			return nil, err
		}
	}
	return res, db.Error
}

// GetUnverifiedSchool is the resolver for the getUnverifiedSchool field.
func (r *queryResolver) GetUnverifiedSchool(ctx context.Context, id int) (*model.UnverifiedSchool, error) {
	v, okHook := r.Sql.Hooks[string(db.GetUnverifiedSchool)].(db.AutoGqlHookGet[model.UnverifiedSchool, int])
	db := r.Sql.Db
	if okHook {
		var err error
		db, err = v.Received(ctx, r.Sql, id)
		if err != nil {
			return nil, err
		}
	}
	db = runtimehelper.GetPreloadSelection(ctx, db, runtimehelper.GetPreloadsMap(ctx, "UnverifiedSchool"))
	if okHook {
		var err error
		db, err = v.BeforeCallDb(ctx, db)
		if err != nil {
			return nil, err
		}
	}
	var res model.UnverifiedSchool
	tableName := r.Sql.Db.Config.NamingStrategy.TableName("UnverifiedSchool")
	db = db.First(&res, tableName+".id = ?", id)
	if okHook {
		r, err := v.AfterCallDb(ctx, &res)
		if err != nil {
			return nil, err
		}
		res = *r
		r, err = v.BeforeReturn(ctx, &res, db)
		if err != nil {
			return nil, err
		}
		res = *r
	}
	return &res, db.Error
}

// QueryUnverifiedSchool is the resolver for the queryUnverifiedSchool field.
func (r *queryResolver) QueryUnverifiedSchool(ctx context.Context, filter *model.UnverifiedSchoolFiltersInput, order *model.UnverifiedSchoolOrder, first *int, offset *int, group []model.UnverifiedSchoolGroup) (*model.UnverifiedSchoolQueryResult, error) {
	v, okHook := r.Sql.Hooks[string(db.QueryUnverifiedSchool)].(db.AutoGqlHookQuery[model.UnverifiedSchool, model.UnverifiedSchoolFiltersInput, model.UnverifiedSchoolOrder])
	db := r.Sql.Db
	if okHook {
		var err error
		db, filter, order, first, offset, err = v.Received(ctx, r.Sql, filter, order, first, offset)
		if err != nil {
			return nil, err
		}
	}
	var res []*model.UnverifiedSchool
	tableName := r.Sql.Db.Config.NamingStrategy.TableName("UnverifiedSchool")
	preloadSubTables := runtimehelper.GetPreloadsMap(ctx, "data").SubTables
	if len(preloadSubTables) > 0 {
		db = runtimehelper.GetPreloadSelection(ctx, db, preloadSubTables[0])
	}
	if filter != nil {
		blackList := make(map[string]struct{})
		sql, arguments := runtimehelper.CombineSimpleQuery(filter.ExtendsDatabaseQuery(db, fmt.Sprintf("%[1]s%[2]s%[1]s", runtimehelper.GetQuoteChar(db), tableName), false, blackList), "AND")
		db.Where(sql, arguments...)
	}
	if okHook {
		var err error
		db, err = v.BeforeCallDb(ctx, db)
		if err != nil {
			return nil, err
		}
	}
	if order != nil {
		if order.Asc != nil {
			db = db.Order(fmt.Sprintf("%[1]s%[2]s%[1]s.%[1]s%[3]s%[1]s asc", runtimehelper.GetQuoteChar(db), tableName, order.Asc))
		}
		if order.Desc != nil {
			db = db.Order(fmt.Sprintf("%[1]s%[2]s%[1]s.%[1]s%[3]s%[1]s desc", runtimehelper.GetQuoteChar(db), tableName, order.Desc))
		}
	}
	var total int64
	db.Model(res).Count(&total)
	if first != nil {
		db = db.Limit(*first)
	}
	if offset != nil {
		db = db.Offset(*offset)
	}
	if len(group) > 0 {
		garr := make([]string, len(group))
		for i, g := range group {
			garr[i] = fmt.Sprintf("%s.%s", tableName, xstrings.ToSnakeCase(string(g)))
		}
		db = db.Group(strings.Join(garr, ","))
	}
	db = db.Find(&res)
	if okHook {
		var err error
		res, err = v.AfterCallDb(ctx, res)
		if err != nil {
			return nil, err
		}
		res, err = v.BeforeReturn(ctx, res, db)
		if err != nil {
			return nil, err
		}
	}
	return &model.UnverifiedSchoolQueryResult{
		Data:       res,
		Count:      len(res),
		TotalCount: int(total),
	}, db.Error
}
func (r *Resolver) AddUnverifiedSchoolPayload() AddUnverifiedSchoolPayloadResolver {
	return &unverifiedSchoolPayloadResolver[*model.AddUnverifiedSchoolPayload]{r}
}
func (r *Resolver) DeleteUnverifiedSchoolPayload() DeleteUnverifiedSchoolPayloadResolver {
	return &unverifiedSchoolPayloadResolver[*model.DeleteUnverifiedSchoolPayload]{r}
}
func (r *Resolver) UpdateUnverifiedSchoolPayload() UpdateUnverifiedSchoolPayloadResolver {
	return &unverifiedSchoolPayloadResolver[*model.UpdateUnverifiedSchoolPayload]{r}
}

type unverifiedSchoolPayload interface {
	*model.AddUnverifiedSchoolPayload | *model.DeleteUnverifiedSchoolPayload | *model.UpdateUnverifiedSchoolPayload
}

type unverifiedSchoolPayloadResolver[T unverifiedSchoolPayload] struct {
	*Resolver
}

func (r *unverifiedSchoolPayloadResolver[T]) UnverifiedSchool(ctx context.Context, obj T, filter *model.UnverifiedSchoolFiltersInput, order *model.UnverifiedSchoolOrder, first *int, offset *int, group []model.UnverifiedSchoolGroup) (*model.UnverifiedSchoolQueryResult, error) {
	q := r.Query().(*queryResolver)
	return q.QueryUnverifiedSchool(ctx, filter, order, first, offset, group)
}

// AddUnverifiedSchool is the resolver for the addUnverifiedSchool field.
func (r *mutationResolver) AddUnverifiedSchool(ctx context.Context, input []*model.UnverifiedSchoolInput) (*model.AddUnverifiedSchoolPayload, error) {
	v, okHook := r.Sql.Hooks[string(db.AddUnverifiedSchool)].(db.AutoGqlHookAdd[model.UnverifiedSchool, model.UnverifiedSchoolInput, model.AddUnverifiedSchoolPayload])
	res := &model.AddUnverifiedSchoolPayload{}
	db := r.Sql.Db
	if okHook {
		var err error
		db, input, err = v.Received(ctx, r.Sql, input)
		if err != nil {
			return nil, err
		}
	}
	obj := make([]model.UnverifiedSchool, len(input))
	for i, v := range input {
		obj[i] = v.MergeToType()
	}
	db = db.Omit(clause.Associations)
	if okHook {
		var err error
		db, obj, err = v.BeforeCallDb(ctx, db, obj)
		if err != nil {
			return nil, err
		}
	}
	db = db.Create(&obj)
	affectedRes := make([]*model.UnverifiedSchool, len(obj))
	for i, v := range obj {
		tmp := v
		affectedRes[i] = &tmp
	}
	res.Affected = affectedRes
	if okHook {
		var err error
		res, err = v.BeforeReturn(ctx, db, obj, res)
		if err != nil {
			return nil, err
		}
	}
	return res, db.Error
}

// UpdateUnverifiedSchool is the resolver for the updateUnverifiedSchool field.
func (r *mutationResolver) UpdateUnverifiedSchool(ctx context.Context, input model.UpdateUnverifiedSchoolInput) (*model.UpdateUnverifiedSchoolPayload, error) {
	v, okHook := r.Sql.Hooks[string(db.UpdateUnverifiedSchool)].(db.AutoGqlHookUpdate[model.UpdateUnverifiedSchoolInput, model.UpdateUnverifiedSchoolPayload])
	db := r.Sql.Db
	if okHook {
		var err error
		db, input, err = v.Received(ctx, r.Sql, &input)
		if err != nil {
			return nil, err
		}
	}
	tableName := r.Sql.Db.Config.NamingStrategy.TableName("UnverifiedSchool")
	blackList := make(map[string]struct{})
	queryDb := db.Select(tableName + ".id")
	sql, arguments := runtimehelper.CombineSimpleQuery(input.Filter.ExtendsDatabaseQuery(queryDb, fmt.Sprintf("%[1]s%[2]s%[1]s", runtimehelper.GetQuoteChar(db), tableName), false, blackList), "AND")
	obj := model.UnverifiedSchool{}
	queryDb = queryDb.Model(&obj).Where(sql, arguments...)
	var toChange []model.UnverifiedSchool
	queryDb.Find(&toChange)
	update := input.Set.MergeToType()
	if okHook {
		var err error
		db, update, err = v.BeforeCallDb(ctx, db, update)
		if err != nil {
			return nil, err
		}
	}
	ids := make([]int, len(toChange))
	for i, one := range toChange {
		ids[i] = one.ID
	}
	db = db.Model(&obj).Where("id IN ?", ids).Updates(update)
	affectedRes := make([]*model.UnverifiedSchool, 0)
	subTables := runtimehelper.GetPreloadsMap(ctx, "affected").SubTables
	if len(subTables) > 0 {
		if preloadMap := subTables[0]; len(preloadMap.Fields) > 0 {
			affectedDb := runtimehelper.GetPreloadSelection(ctx, db, preloadMap)
			affectedDb = affectedDb.Model(&obj)
			affectedDb.Find(&affectedRes)
		}
	}

	res := &model.UpdateUnverifiedSchoolPayload{
		Count:    int(db.RowsAffected),
		Affected: affectedRes,
	}
	if okHook {
		var err error
		res, err = v.BeforeReturn(ctx, db, res)
		if err != nil {
			return nil, err
		}
	}
	return res, db.Error
}

// DeleteUnverifiedSchool is the resolver for the deleteUnverifiedSchool field.
func (r *mutationResolver) DeleteUnverifiedSchool(ctx context.Context, filter model.UnverifiedSchoolFiltersInput) (*model.DeleteUnverifiedSchoolPayload, error) {
	v, okHook := r.Sql.Hooks[string(db.DeleteUnverifiedSchool)].(db.AutoGqlHookDelete[model.UnverifiedSchoolFiltersInput, model.DeleteUnverifiedSchoolPayload])
	db := r.Sql.Db
	if okHook {
		var err error
		db, filter, err = v.Received(ctx, r.Sql, &filter)
		if err != nil {
			return nil, err
		}
	}
	tableName := r.Sql.Db.Config.NamingStrategy.TableName("UnverifiedSchool")
	blackList := make(map[string]struct{})
	queryDb := db.Select(tableName + ".id")
	sql, arguments := runtimehelper.CombineSimpleQuery(filter.ExtendsDatabaseQuery(queryDb, fmt.Sprintf("%[1]s%[2]s%[1]s", runtimehelper.GetQuoteChar(db), tableName), false, blackList), "AND")
	obj := model.UnverifiedSchool{}
	queryDb = queryDb.Model(&obj).Where(sql, arguments...)
	var toChange []model.UnverifiedSchool
	queryDb.Find(&toChange)
	if okHook {
		var err error
		db, err = v.BeforeCallDb(ctx, db)
		if err != nil {
			return nil, err
		}
	}
	ids := make([]int, len(toChange))
	for i, one := range toChange {
		ids[i] = one.ID
	}
	db = db.Model(&obj).Where("id IN ?", ids).Delete(&obj)
	msg := fmt.Sprintf("%d rows deleted", db.RowsAffected)
	res := &model.DeleteUnverifiedSchoolPayload{
		Count: int(db.RowsAffected),
		Msg:   &msg,
	}
	if okHook {
		var err error
		res, err = v.BeforeReturn(ctx, db, res)
		if err != nil {
			return nil, err
		}
	}
	return res, db.Error
}
