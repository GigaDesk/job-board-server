// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"strings"

	"github.com/fasibio/autogql/runtimehelper"
	"gorm.io/gorm"
)

const extendsDatabaseFieldNameFormat string = "%[2]s.%[1]s%[3]s%[1]s"

// PrimaryKeyName return the name of primarykey for Table Admin
func (d *AdminFiltersInput) PrimaryKeyName() string {
	return "id"
}

// ExtendsDatabaseQuery create condition from AdminFiltersInput values
func (d *AdminFiltersInput) ExtendsDatabaseQuery(db *gorm.DB, alias string, deep bool, blackList map[string]struct{}) []runtimehelper.ConditionElement {
	res := make([]runtimehelper.ConditionElement, 0)
	if d.And != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.And {
			tmp = append(tmp, runtimehelper.Complex(runtimehelper.RelationAnd, v.ExtendsDatabaseQuery(db, alias, true, blackList)...))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationAnd, tmp...))
	}

	if d.Or != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.Or {

			tmp = append(tmp, runtimehelper.Complex(runtimehelper.RelationAnd, v.ExtendsDatabaseQuery(db, alias, true, blackList)...))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationOr, tmp...))
	}

	if d.Not != nil {
		res = append(res, runtimehelper.Complex(runtimehelper.RelationNot, d.Not.ExtendsDatabaseQuery(db, alias, true, blackList)...))
	}
	if d.ID != nil {
		res = append(res, d.ID.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "id"), true, blackList)...)
	}
	if d.CreatedAt != nil {
		res = append(res, d.CreatedAt.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "created_at"), true, blackList)...)
	}
	if d.UpdatedAt != nil {
		res = append(res, d.UpdatedAt.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "updated_at"), true, blackList)...)
	}
	if d.Name != nil {
		res = append(res, d.Name.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "name"), true, blackList)...)
	}
	if d.PhoneNumber != nil {
		res = append(res, d.PhoneNumber.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "phone_number"), true, blackList)...)
	}
	if d.Password != nil {
		res = append(res, d.Password.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "password"), true, blackList)...)
	}

	return res
}

// PrimaryKeyName return the name of primarykey for Table Application
func (d *ApplicationFiltersInput) PrimaryKeyName() string {
	return "id"
}

// ExtendsDatabaseQuery create condition from ApplicationFiltersInput values
func (d *ApplicationFiltersInput) ExtendsDatabaseQuery(db *gorm.DB, alias string, deep bool, blackList map[string]struct{}) []runtimehelper.ConditionElement {
	res := make([]runtimehelper.ConditionElement, 0)
	if d.And != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.And {
			tmp = append(tmp, runtimehelper.Complex(runtimehelper.RelationAnd, v.ExtendsDatabaseQuery(db, alias, true, blackList)...))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationAnd, tmp...))
	}

	if d.Or != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.Or {

			tmp = append(tmp, runtimehelper.Complex(runtimehelper.RelationAnd, v.ExtendsDatabaseQuery(db, alias, true, blackList)...))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationOr, tmp...))
	}

	if d.Not != nil {
		res = append(res, runtimehelper.Complex(runtimehelper.RelationNot, d.Not.ExtendsDatabaseQuery(db, alias, true, blackList)...))
	}
	if d.ID != nil {
		res = append(res, d.ID.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "id"), true, blackList)...)
	}
	if d.CreatedAt != nil {
		res = append(res, d.CreatedAt.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "created_at"), true, blackList)...)
	}
	if d.UpdatedAt != nil {
		res = append(res, d.UpdatedAt.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "updated_at"), true, blackList)...)
	}
	if d.JobID != nil {
		res = append(res, d.JobID.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "job_id"), true, blackList)...)
	}
	if d.EmployeeID != nil {
		res = append(res, d.EmployeeID.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "employee_id"), true, blackList)...)
	}
	if d.EducationLevel != nil {
		res = append(res, d.EducationLevel.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "education_level"), true, blackList)...)
	}
	if d.Experience != nil {
		res = append(res, d.Experience.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "experience"), true, blackList)...)
	}
	if d.CoverLetterURL != nil {
		res = append(res, d.CoverLetterURL.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "cover_letter_url"), true, blackList)...)
	}
	if d.ResumeeURL != nil {
		res = append(res, d.ResumeeURL.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "resumee_url"), true, blackList)...)
	}
	if d.Status != nil {
		res = append(res, d.Status.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "status"), true, blackList)...)
	}

	return res
}

// PrimaryKeyName return the name of primarykey for Table Employee
func (d *EmployeeFiltersInput) PrimaryKeyName() string {
	return "id"
}

// ExtendsDatabaseQuery create condition from EmployeeFiltersInput values
func (d *EmployeeFiltersInput) ExtendsDatabaseQuery(db *gorm.DB, alias string, deep bool, blackList map[string]struct{}) []runtimehelper.ConditionElement {
	res := make([]runtimehelper.ConditionElement, 0)
	if d.And != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.And {
			tmp = append(tmp, runtimehelper.Complex(runtimehelper.RelationAnd, v.ExtendsDatabaseQuery(db, alias, true, blackList)...))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationAnd, tmp...))
	}

	if d.Or != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.Or {

			tmp = append(tmp, runtimehelper.Complex(runtimehelper.RelationAnd, v.ExtendsDatabaseQuery(db, alias, true, blackList)...))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationOr, tmp...))
	}

	if d.Not != nil {
		res = append(res, runtimehelper.Complex(runtimehelper.RelationNot, d.Not.ExtendsDatabaseQuery(db, alias, true, blackList)...))
	}
	if d.ID != nil {
		res = append(res, d.ID.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "id"), true, blackList)...)
	}
	if d.CreatedAt != nil {
		res = append(res, d.CreatedAt.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "created_at"), true, blackList)...)
	}
	if d.UpdatedAt != nil {
		res = append(res, d.UpdatedAt.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "updated_at"), true, blackList)...)
	}
	if d.Name != nil {
		res = append(res, d.Name.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "name"), true, blackList)...)
	}
	if d.PhoneNumber != nil {
		res = append(res, d.PhoneNumber.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "phone_number"), true, blackList)...)
	}
	if d.Password != nil {
		res = append(res, d.Password.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "password"), true, blackList)...)
	}
	if d.Profilepicture != nil {
		res = append(res, d.Profilepicture.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "profilepicture"), true, blackList)...)
	}

	return res
}

// PrimaryKeyName return the name of primarykey for Table Employer
func (d *EmployerFiltersInput) PrimaryKeyName() string {
	return "id"
}

// ExtendsDatabaseQuery create condition from EmployerFiltersInput values
func (d *EmployerFiltersInput) ExtendsDatabaseQuery(db *gorm.DB, alias string, deep bool, blackList map[string]struct{}) []runtimehelper.ConditionElement {
	res := make([]runtimehelper.ConditionElement, 0)
	if d.And != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.And {
			tmp = append(tmp, runtimehelper.Complex(runtimehelper.RelationAnd, v.ExtendsDatabaseQuery(db, alias, true, blackList)...))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationAnd, tmp...))
	}

	if d.Or != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.Or {

			tmp = append(tmp, runtimehelper.Complex(runtimehelper.RelationAnd, v.ExtendsDatabaseQuery(db, alias, true, blackList)...))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationOr, tmp...))
	}

	if d.Not != nil {
		res = append(res, runtimehelper.Complex(runtimehelper.RelationNot, d.Not.ExtendsDatabaseQuery(db, alias, true, blackList)...))
	}
	if d.ID != nil {
		res = append(res, d.ID.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "id"), true, blackList)...)
	}
	if d.CreatedAt != nil {
		res = append(res, d.CreatedAt.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "created_at"), true, blackList)...)
	}
	if d.UpdatedAt != nil {
		res = append(res, d.UpdatedAt.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "updated_at"), true, blackList)...)
	}
	if d.Name != nil {
		res = append(res, d.Name.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "name"), true, blackList)...)
	}
	if d.PhoneNumber != nil {
		res = append(res, d.PhoneNumber.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "phone_number"), true, blackList)...)
	}
	if d.Password != nil {
		res = append(res, d.Password.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "password"), true, blackList)...)
	}
	if d.Badge != nil {
		res = append(res, d.Badge.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "badge"), true, blackList)...)
	}
	if d.Website != nil {
		res = append(res, d.Website.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "website"), true, blackList)...)
	}

	return res
}

// PrimaryKeyName return the name of primarykey for Table Job
func (d *JobFiltersInput) PrimaryKeyName() string {
	return "id"
}

// ExtendsDatabaseQuery create condition from JobFiltersInput values
func (d *JobFiltersInput) ExtendsDatabaseQuery(db *gorm.DB, alias string, deep bool, blackList map[string]struct{}) []runtimehelper.ConditionElement {
	res := make([]runtimehelper.ConditionElement, 0)
	if d.And != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.And {
			tmp = append(tmp, runtimehelper.Complex(runtimehelper.RelationAnd, v.ExtendsDatabaseQuery(db, alias, true, blackList)...))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationAnd, tmp...))
	}

	if d.Or != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.Or {

			tmp = append(tmp, runtimehelper.Complex(runtimehelper.RelationAnd, v.ExtendsDatabaseQuery(db, alias, true, blackList)...))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationOr, tmp...))
	}

	if d.Not != nil {
		res = append(res, runtimehelper.Complex(runtimehelper.RelationNot, d.Not.ExtendsDatabaseQuery(db, alias, true, blackList)...))
	}
	if d.ID != nil {
		res = append(res, d.ID.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "id"), true, blackList)...)
	}
	if d.CreatedAt != nil {
		res = append(res, d.CreatedAt.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "created_at"), true, blackList)...)
	}
	if d.UpdatedAt != nil {
		res = append(res, d.UpdatedAt.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "updated_at"), true, blackList)...)
	}
	if d.Title != nil {
		res = append(res, d.Title.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "title"), true, blackList)...)
	}
	if d.Industry != nil {
		res = append(res, d.Industry.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "industry"), true, blackList)...)
	}
	if d.Description != nil {
		res = append(res, d.Description.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "description"), true, blackList)...)
	}
	if d.Level != nil {
		res = append(res, d.Level.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "level"), true, blackList)...)
	}
	if d.Location != nil {
		res = append(res, d.Location.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "location"), true, blackList)...)
	}
	if d.Deadline != nil {
		res = append(res, d.Deadline.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "deadline"), true, blackList)...)
	}
	if d.EducationLevel != nil {
		res = append(res, d.EducationLevel.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "education_level"), true, blackList)...)
	}
	if d.Experience != nil {
		res = append(res, d.Experience.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "experience"), true, blackList)...)
	}
	if d.MinSalary != nil {
		res = append(res, d.MinSalary.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "min_salary"), true, blackList)...)
	}
	if d.MaxSalary != nil {
		res = append(res, d.MaxSalary.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "max_salary"), true, blackList)...)
	}
	if d.Requirements != nil {
		res = append(res, d.Requirements.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "requirements"), true, blackList)...)
	}
	if d.EmployerID != nil {
		res = append(res, d.EmployerID.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "employer_id"), true, blackList)...)
	}
	if d.JobURL != nil {
		res = append(res, d.JobURL.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "job_url"), true, blackList)...)
	}

	return res
}

// PrimaryKeyName return the name of primarykey for Table UnapprovedJob
func (d *UnapprovedJobFiltersInput) PrimaryKeyName() string {
	return "id"
}

// ExtendsDatabaseQuery create condition from UnapprovedJobFiltersInput values
func (d *UnapprovedJobFiltersInput) ExtendsDatabaseQuery(db *gorm.DB, alias string, deep bool, blackList map[string]struct{}) []runtimehelper.ConditionElement {
	res := make([]runtimehelper.ConditionElement, 0)
	if d.And != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.And {
			tmp = append(tmp, runtimehelper.Complex(runtimehelper.RelationAnd, v.ExtendsDatabaseQuery(db, alias, true, blackList)...))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationAnd, tmp...))
	}

	if d.Or != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.Or {

			tmp = append(tmp, runtimehelper.Complex(runtimehelper.RelationAnd, v.ExtendsDatabaseQuery(db, alias, true, blackList)...))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationOr, tmp...))
	}

	if d.Not != nil {
		res = append(res, runtimehelper.Complex(runtimehelper.RelationNot, d.Not.ExtendsDatabaseQuery(db, alias, true, blackList)...))
	}
	if d.ID != nil {
		res = append(res, d.ID.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "id"), true, blackList)...)
	}
	if d.CreatedAt != nil {
		res = append(res, d.CreatedAt.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "created_at"), true, blackList)...)
	}
	if d.UpdatedAt != nil {
		res = append(res, d.UpdatedAt.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "updated_at"), true, blackList)...)
	}
	if d.Title != nil {
		res = append(res, d.Title.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "title"), true, blackList)...)
	}
	if d.Industry != nil {
		res = append(res, d.Industry.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "industry"), true, blackList)...)
	}
	if d.Description != nil {
		res = append(res, d.Description.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "description"), true, blackList)...)
	}
	if d.Level != nil {
		res = append(res, d.Level.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "level"), true, blackList)...)
	}
	if d.Location != nil {
		res = append(res, d.Location.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "location"), true, blackList)...)
	}
	if d.Deadline != nil {
		res = append(res, d.Deadline.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "deadline"), true, blackList)...)
	}
	if d.EducationLevel != nil {
		res = append(res, d.EducationLevel.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "education_level"), true, blackList)...)
	}
	if d.Experience != nil {
		res = append(res, d.Experience.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "experience"), true, blackList)...)
	}
	if d.MinSalary != nil {
		res = append(res, d.MinSalary.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "min_salary"), true, blackList)...)
	}
	if d.MaxSalary != nil {
		res = append(res, d.MaxSalary.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "max_salary"), true, blackList)...)
	}
	if d.Requirements != nil {
		res = append(res, d.Requirements.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "requirements"), true, blackList)...)
	}
	if d.EmployerID != nil {
		res = append(res, d.EmployerID.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "employer_id"), true, blackList)...)
	}
	if d.JobURL != nil {
		res = append(res, d.JobURL.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "job_url"), true, blackList)...)
	}

	return res
}

// PrimaryKeyName return the name of primarykey for Table UnverifiedAdmin
func (d *UnverifiedAdminFiltersInput) PrimaryKeyName() string {
	return "id"
}

// ExtendsDatabaseQuery create condition from UnverifiedAdminFiltersInput values
func (d *UnverifiedAdminFiltersInput) ExtendsDatabaseQuery(db *gorm.DB, alias string, deep bool, blackList map[string]struct{}) []runtimehelper.ConditionElement {
	res := make([]runtimehelper.ConditionElement, 0)
	if d.And != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.And {
			tmp = append(tmp, runtimehelper.Complex(runtimehelper.RelationAnd, v.ExtendsDatabaseQuery(db, alias, true, blackList)...))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationAnd, tmp...))
	}

	if d.Or != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.Or {

			tmp = append(tmp, runtimehelper.Complex(runtimehelper.RelationAnd, v.ExtendsDatabaseQuery(db, alias, true, blackList)...))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationOr, tmp...))
	}

	if d.Not != nil {
		res = append(res, runtimehelper.Complex(runtimehelper.RelationNot, d.Not.ExtendsDatabaseQuery(db, alias, true, blackList)...))
	}
	if d.ID != nil {
		res = append(res, d.ID.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "id"), true, blackList)...)
	}
	if d.CreatedAt != nil {
		res = append(res, d.CreatedAt.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "created_at"), true, blackList)...)
	}
	if d.UpdatedAt != nil {
		res = append(res, d.UpdatedAt.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "updated_at"), true, blackList)...)
	}
	if d.Name != nil {
		res = append(res, d.Name.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "name"), true, blackList)...)
	}
	if d.PhoneNumber != nil {
		res = append(res, d.PhoneNumber.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "phone_number"), true, blackList)...)
	}
	if d.Password != nil {
		res = append(res, d.Password.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "password"), true, blackList)...)
	}

	return res
}

// PrimaryKeyName return the name of primarykey for Table UnverifiedEmployee
func (d *UnverifiedEmployeeFiltersInput) PrimaryKeyName() string {
	return "id"
}

// ExtendsDatabaseQuery create condition from UnverifiedEmployeeFiltersInput values
func (d *UnverifiedEmployeeFiltersInput) ExtendsDatabaseQuery(db *gorm.DB, alias string, deep bool, blackList map[string]struct{}) []runtimehelper.ConditionElement {
	res := make([]runtimehelper.ConditionElement, 0)
	if d.And != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.And {
			tmp = append(tmp, runtimehelper.Complex(runtimehelper.RelationAnd, v.ExtendsDatabaseQuery(db, alias, true, blackList)...))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationAnd, tmp...))
	}

	if d.Or != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.Or {

			tmp = append(tmp, runtimehelper.Complex(runtimehelper.RelationAnd, v.ExtendsDatabaseQuery(db, alias, true, blackList)...))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationOr, tmp...))
	}

	if d.Not != nil {
		res = append(res, runtimehelper.Complex(runtimehelper.RelationNot, d.Not.ExtendsDatabaseQuery(db, alias, true, blackList)...))
	}
	if d.ID != nil {
		res = append(res, d.ID.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "id"), true, blackList)...)
	}
	if d.CreatedAt != nil {
		res = append(res, d.CreatedAt.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "created_at"), true, blackList)...)
	}
	if d.UpdatedAt != nil {
		res = append(res, d.UpdatedAt.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "updated_at"), true, blackList)...)
	}
	if d.Name != nil {
		res = append(res, d.Name.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "name"), true, blackList)...)
	}
	if d.PhoneNumber != nil {
		res = append(res, d.PhoneNumber.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "phone_number"), true, blackList)...)
	}
	if d.Password != nil {
		res = append(res, d.Password.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "password"), true, blackList)...)
	}
	if d.Profilepicture != nil {
		res = append(res, d.Profilepicture.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "profilepicture"), true, blackList)...)
	}

	return res
}

// PrimaryKeyName return the name of primarykey for Table UnverifiedEmployer
func (d *UnverifiedEmployerFiltersInput) PrimaryKeyName() string {
	return "id"
}

// ExtendsDatabaseQuery create condition from UnverifiedEmployerFiltersInput values
func (d *UnverifiedEmployerFiltersInput) ExtendsDatabaseQuery(db *gorm.DB, alias string, deep bool, blackList map[string]struct{}) []runtimehelper.ConditionElement {
	res := make([]runtimehelper.ConditionElement, 0)
	if d.And != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.And {
			tmp = append(tmp, runtimehelper.Complex(runtimehelper.RelationAnd, v.ExtendsDatabaseQuery(db, alias, true, blackList)...))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationAnd, tmp...))
	}

	if d.Or != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.Or {

			tmp = append(tmp, runtimehelper.Complex(runtimehelper.RelationAnd, v.ExtendsDatabaseQuery(db, alias, true, blackList)...))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationOr, tmp...))
	}

	if d.Not != nil {
		res = append(res, runtimehelper.Complex(runtimehelper.RelationNot, d.Not.ExtendsDatabaseQuery(db, alias, true, blackList)...))
	}
	if d.ID != nil {
		res = append(res, d.ID.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "id"), true, blackList)...)
	}
	if d.CreatedAt != nil {
		res = append(res, d.CreatedAt.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "created_at"), true, blackList)...)
	}
	if d.UpdatedAt != nil {
		res = append(res, d.UpdatedAt.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "updated_at"), true, blackList)...)
	}
	if d.Name != nil {
		res = append(res, d.Name.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "name"), true, blackList)...)
	}
	if d.PhoneNumber != nil {
		res = append(res, d.PhoneNumber.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "phone_number"), true, blackList)...)
	}
	if d.Password != nil {
		res = append(res, d.Password.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "password"), true, blackList)...)
	}
	if d.Badge != nil {
		res = append(res, d.Badge.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "badge"), true, blackList)...)
	}
	if d.Website != nil {
		res = append(res, d.Website.ExtendsDatabaseQuery(db, fmt.Sprintf(extendsDatabaseFieldNameFormat, runtimehelper.GetQuoteChar(db), alias, "website"), true, blackList)...)
	}

	return res
}

func (d *StringFilterInput) ExtendsDatabaseQuery(db *gorm.DB, fieldName string, deep bool, blackList map[string]struct{}) []runtimehelper.ConditionElement {
	res := make([]runtimehelper.ConditionElement, 0)
	if d.And != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.And {
			tmp = append(tmp, runtimehelper.Equal(fieldName, *v))
		}
		res = append(res, tmp...)
	}
	if d.Contains != nil {
		res = append(res, runtimehelper.Like(fieldName, fmt.Sprintf("%%%s%%", *d.Contains)))
	}

	if d.Containsi != nil {
		res = append(res, runtimehelper.Like(fmt.Sprintf("lower(%s)", fieldName), fmt.Sprintf("%%%s%%", strings.ToLower(*d.Containsi))))
	}

	if d.EndsWith != nil {
		res = append(res, runtimehelper.Like(fieldName, fmt.Sprintf("%%%s", *d.EndsWith)))
	}

	if d.Eq != nil {
		res = append(res, runtimehelper.Equal(fieldName, *d.Eq))
	}

	if d.Eqi != nil {
		res = append(res, runtimehelper.Equal(fmt.Sprintf("lower(%s)", fieldName), strings.ToLower(*d.Eqi)))
	}

	if d.In != nil {
		res = append(res, runtimehelper.In(fieldName, d.In))
	}

	if d.Ne != nil {
		res = append(res, runtimehelper.NotEqual(fieldName, *d.Ne))
	}

	if d.Not != nil {
		res = append(res, runtimehelper.Complex(runtimehelper.RelationNot, d.Not.ExtendsDatabaseQuery(db, fieldName, true, blackList)...))
	}

	if d.NotContains != nil {
		res = append(res, runtimehelper.NotLike(fieldName, fmt.Sprintf("%%%s%%", *d.NotContains)))
	}

	if d.NotContainsi != nil {
		res = append(res, runtimehelper.NotLike(fmt.Sprintf("lower(%s)", fieldName), fmt.Sprintf("%%%s%%", strings.ToLower(*d.NotContainsi))))
	}

	if d.NotIn != nil {
		res = append(res, runtimehelper.NotIn(fieldName, d.NotIn))
	}

	if d.NotNull != nil {
		res = append(res, runtimehelper.NotNull(fieldName))
	}

	if d.Null != nil {
		res = append(res, runtimehelper.Null(fieldName))
	}

	if d.Or != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.Or {
			tmp = append(tmp, runtimehelper.Equal(fieldName, *v))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationOr, tmp...))
	}

	if d.StartsWith != nil {
		res = append(res, runtimehelper.Like(fieldName, fmt.Sprintf("%s%%", *d.StartsWith)))
	}

	return res
}

// ExtendsDatabaseQuery create condition from values
func (d *IntFilterInput) ExtendsDatabaseQuery(db *gorm.DB, fieldName string, deep bool, blackList map[string]struct{}) []runtimehelper.ConditionElement {

	res := make([]runtimehelper.ConditionElement, 0)

	if d.And != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.And {
			tmp = append(tmp, runtimehelper.Equal(fieldName, *v))
		}
		res = append(res, tmp...)
	}

	if d.Between != nil {
		res = append(res, runtimehelper.Between(fieldName, d.Between.Start, d.Between.End))
	}

	if d.Eq != nil {
		res = append(res, runtimehelper.Equal(fieldName, *d.Eq))
	}
	if d.Gt != nil {
		res = append(res, runtimehelper.More(fieldName, *d.Gt))
	}

	if d.Gte != nil {
		res = append(res, runtimehelper.MoreOrEqual(fieldName, *d.Gte))
	}

	if d.In != nil {
		res = append(res, runtimehelper.In(fieldName, d.In))
	}

	if d.Lt != nil {
		res = append(res, runtimehelper.Less(fieldName, *d.Lt))
	}

	if d.Lte != nil {
		res = append(res, runtimehelper.LessOrEqual(fieldName, *d.Lte))
	}

	if d.Ne != nil {
		res = append(res, runtimehelper.NotEqual(fieldName, *d.Ne))
	}
	if d.Not != nil {
		res = append(res, runtimehelper.Complex(runtimehelper.RelationNot, d.Not.ExtendsDatabaseQuery(db, fieldName, true, blackList)...))
	}

	if d.NotIn != nil {
		res = append(res, runtimehelper.NotIn(fieldName, d.NotIn))
	}

	if d.NotNull != nil && *d.NotNull {
		res = append(res, runtimehelper.NotNull(fieldName))
	}

	if d.Null != nil && *d.Null {
		res = append(res, runtimehelper.Null(fieldName))
	}

	if d.Or != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.Or {
			tmp = append(tmp, runtimehelper.Equal(fieldName, *v))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationOr, tmp...))
	}

	return res
}

// ExtendsDatabaseQuery create condition from values
func (d *FloatFilterInput) ExtendsDatabaseQuery(db *gorm.DB, fieldName string, deep bool, blackList map[string]struct{}) []runtimehelper.ConditionElement {

	res := make([]runtimehelper.ConditionElement, 0)

	if d.And != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.And {
			tmp = append(tmp, runtimehelper.Equal(fieldName, *v))
		}
		res = append(res, tmp...)
	}

	if d.Between != nil {
		res = append(res, runtimehelper.Between(fieldName, d.Between.Start, d.Between.End))
	}

	if d.Eq != nil {
		res = append(res, runtimehelper.Equal(fieldName, *d.Eq))
	}
	if d.Gt != nil {
		res = append(res, runtimehelper.More(fieldName, *d.Gt))
	}

	if d.Gte != nil {
		res = append(res, runtimehelper.MoreOrEqual(fieldName, *d.Gte))
	}

	if d.In != nil {
		res = append(res, runtimehelper.In(fieldName, d.In))
	}

	if d.Lt != nil {
		res = append(res, runtimehelper.Less(fieldName, *d.Lt))
	}

	if d.Lte != nil {
		res = append(res, runtimehelper.LessOrEqual(fieldName, *d.Lte))
	}

	if d.Ne != nil {
		res = append(res, runtimehelper.NotEqual(fieldName, *d.Ne))
	}
	if d.Not != nil {
		res = append(res, runtimehelper.Complex(runtimehelper.RelationNot, d.Not.ExtendsDatabaseQuery(db, fieldName, true, blackList)...))
	}

	if d.NotIn != nil {
		res = append(res, runtimehelper.NotIn(fieldName, d.NotIn))
	}

	if d.NotNull != nil && *d.NotNull {
		res = append(res, runtimehelper.NotNull(fieldName))
	}

	if d.Null != nil && *d.Null {
		res = append(res, runtimehelper.Null(fieldName))
	}

	if d.Or != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.Or {
			tmp = append(tmp, runtimehelper.Equal(fieldName, *v))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationOr, tmp...))
	}

	return res
}

// ExtendsDatabaseQuery create condition from values
func (d *BooleanFilterInput) ExtendsDatabaseQuery(db *gorm.DB, fieldName string, deep bool, blackList map[string]struct{}) []runtimehelper.ConditionElement {
	res := make([]runtimehelper.ConditionElement, 0)

	if d.And != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.And {
			tmp = append(tmp, runtimehelper.Equal(fieldName, *v))
		}
		res = append(res, tmp...)
	}

	if d.Is != nil {
		res = append(res, runtimehelper.Equal(fieldName, *d.Is))
	}

	if d.Not != nil {
		res = append(res, runtimehelper.Complex(runtimehelper.RelationNot, d.Not.ExtendsDatabaseQuery(db, fieldName, true, blackList)...))
	}

	if d.NotNull != nil && *d.NotNull {
		res = append(res, runtimehelper.NotNull(fieldName))
	}

	if d.Null != nil && *d.Null {
		res = append(res, runtimehelper.Null(fieldName))
	}

	if d.Or != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.Or {
			tmp = append(tmp, runtimehelper.Equal(fieldName, *v))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationOr, tmp...))
	}

	return res
}

// ExtendsDatabaseQuery create condition from values
func (d *TimeFilterInput) ExtendsDatabaseQuery(db *gorm.DB, fieldName string, deep bool, blackList map[string]struct{}) []runtimehelper.ConditionElement {

	res := make([]runtimehelper.ConditionElement, 0)

	if d.And != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.And {
			tmp = append(tmp, runtimehelper.Equal(fieldName, *v))
		}
		res = append(res, tmp...)
	}

	if d.Between != nil {
		res = append(res, runtimehelper.Between(fieldName, d.Between.Start, d.Between.End))
	}

	if d.Eq != nil {
		res = append(res, runtimehelper.Equal(fieldName, *d.Eq))
	}
	if d.Gt != nil {
		res = append(res, runtimehelper.More(fieldName, *d.Gt))
	}

	if d.Gte != nil {
		res = append(res, runtimehelper.MoreOrEqual(fieldName, *d.Gte))
	}

	if d.In != nil {
		res = append(res, runtimehelper.In(fieldName, d.In))
	}

	if d.Lt != nil {
		res = append(res, runtimehelper.Less(fieldName, *d.Lt))
	}

	if d.Lte != nil {
		res = append(res, runtimehelper.LessOrEqual(fieldName, *d.Lte))
	}

	if d.Ne != nil {
		res = append(res, runtimehelper.NotEqual(fieldName, *d.Ne))
	}
	if d.Not != nil {
		res = append(res, runtimehelper.Complex(runtimehelper.RelationNot, d.Not.ExtendsDatabaseQuery(db, fieldName, true, blackList)...))
	}

	if d.NotIn != nil {
		res = append(res, runtimehelper.NotIn(fieldName, d.NotIn))

	}

	if d.NotNull != nil && *d.NotNull {
		res = append(res, runtimehelper.NotNull(fieldName))
	}

	if d.Null != nil && *d.Null {
		res = append(res, runtimehelper.Null(fieldName))
	}

	if d.Or != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.Or {
			tmp = append(tmp, runtimehelper.Equal(fieldName, *v))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationOr, tmp...))
	}

	return res
}

// ExtendsDatabaseQuery create condition from values
func (d *SoftDeleteFilterInput) ExtendsDatabaseQuery(db *gorm.DB, fieldName string, deep bool, blackList map[string]struct{}) []runtimehelper.ConditionElement {

	res := make([]runtimehelper.ConditionElement, 0)

	if d.And != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.And {
			tmp = append(tmp, runtimehelper.Equal(fieldName, *v))
		}
		res = append(res, tmp...)
	}

	if d.Between != nil {
		res = append(res, runtimehelper.Between(fieldName, d.Between.Start, d.Between.End))
	}

	if d.Eq != nil {
		res = append(res, runtimehelper.Equal(fieldName, *d.Eq))
	}
	if d.Gt != nil {
		res = append(res, runtimehelper.More(fieldName, *d.Gt))
	}

	if d.Gte != nil {
		res = append(res, runtimehelper.MoreOrEqual(fieldName, *d.Gte))
	}

	if d.In != nil {
		res = append(res, runtimehelper.In(fieldName, d.In))
	}

	if d.Lt != nil {
		res = append(res, runtimehelper.Less(fieldName, *d.Lt))
	}

	if d.Lte != nil {
		res = append(res, runtimehelper.LessOrEqual(fieldName, *d.Lte))
	}

	if d.Ne != nil {
		res = append(res, runtimehelper.NotEqual(fieldName, *d.Ne))
	}
	if d.Not != nil {
		res = append(res, runtimehelper.Complex(runtimehelper.RelationNot, d.Not.ExtendsDatabaseQuery(db, fieldName, true, blackList)...))
	}

	if d.NotIn != nil {
		res = append(res, runtimehelper.NotIn(fieldName, d.NotIn))

	}

	if d.NotNull != nil && *d.NotNull {
		res = append(res, runtimehelper.NotNull(fieldName))
	}

	if d.Null != nil && *d.Null {
		res = append(res, runtimehelper.Null(fieldName))
	}

	if d.Or != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.Or {
			tmp = append(tmp, runtimehelper.Equal(fieldName, *v))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationOr, tmp...))
	}

	return res
}

// ExtendsDatabaseQuery create condition from values
func (d *IDFilterInput) ExtendsDatabaseQuery(db *gorm.DB, fieldName string, deep bool, blackList map[string]struct{}) []runtimehelper.ConditionElement {

	res := make([]runtimehelper.ConditionElement, 0)

	if d.And != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.And {
			tmp = append(tmp, runtimehelper.Equal(fieldName, *v))
		}
		res = append(res, tmp...)
	}

	if d.Eq != nil {
		res = append(res, runtimehelper.Equal(fieldName, *d.Eq))
	}

	if d.In != nil {
		res = append(res, runtimehelper.In(fieldName, d.In))
	}

	if d.Ne != nil {
		res = append(res, runtimehelper.NotEqual(fieldName, *d.Ne))
	}
	if d.Not != nil {
		res = append(res, runtimehelper.Complex(runtimehelper.RelationNot, d.Not.ExtendsDatabaseQuery(db, fieldName, true, blackList)...))
	}

	if d.NotNull != nil && *d.NotNull {
		res = append(res, runtimehelper.NotNull(fieldName))
	}

	if d.Null != nil && *d.Null {
		res = append(res, runtimehelper.Null(fieldName))
	}

	if d.Or != nil {
		tmp := make([]runtimehelper.ConditionElement, 0)
		for _, v := range d.Or {
			tmp = append(tmp, runtimehelper.Equal(fieldName, *v))
		}
		res = append(res, runtimehelper.Complex(runtimehelper.RelationOr, tmp...))
	}

	return res
}
