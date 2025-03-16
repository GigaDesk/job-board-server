// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"time"

	"github.com/mitchellh/mapstructure"
)

// GetInputStruct returns struct filled from map obj defined by name
// Example useage struct validation with github.com/go-playground/validator by directive:
//
//	func ValidateDirective(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
//	  field := graphql.GetPathContext(ctx)
//	  if data, ok := obj.(map[string]interface{}); ok {
//	    for _, v := range field.ParentField.Field.Arguments {
//	      name := v.Value.ExpectedType.Name()
//	      model, err := model.GetInputStruct(name, data)
//	      if err != nil {
//	        //handle not found error
//	      }
//	      if err := validate.Struct(model); err != nil {
//	      //handle error
//	      }
//	    }
//	  }
//	  return next(ctx)
//	}
func GetInputStruct(name string, obj map[string]interface{}) (interface{}, error) {
	switch name {
	case "AdminInput":
		return AdminInputFromMap(obj)
	case "JobInput":
		return JobInputFromMap(obj)
	case "SchoolInput":
		return SchoolInputFromMap(obj)
	case "StudentInput":
		return StudentInputFromMap(obj)
	case "UnapprovedJobInput":
		return UnapprovedJobInputFromMap(obj)
	case "UnverifiedAdminInput":
		return UnverifiedAdminInputFromMap(obj)
	case "UnverifiedSchoolInput":
		return UnverifiedSchoolInputFromMap(obj)
	}
	return nil, fmt.Errorf("%s not found", name)
}

// AdminInputFromMap return a AdminInput from data map
// use github.com/mitchellh/mapstructure with reflaction
func AdminInputFromMap(data map[string]interface{}) (AdminInput, error) {
	model := AdminInput{}
	err := mapstructure.Decode(data, &model)
	return model, err
}

// MergeToType returns a map with all values set to AdminPatch
func (d *AdminPatch) MergeToType() map[string]interface{} {
	res := make(map[string]interface{})
	if d.Name != nil {
		res["name"] = *d.Name
	}
	if d.PhoneNumber != nil {
		res["phone_number"] = *d.PhoneNumber
	}
	if d.Password != nil {
		res["password"] = *d.Password
	}
	return res
}

// MergeToType retuns a Admin filled from AdminInput
func (d *AdminInput) MergeToType() Admin {

	tmpName := d.Name

	tmpPhoneNumber := d.PhoneNumber

	tmpPassword := d.Password
	return Admin{
		Name:        tmpName,
		PhoneNumber: tmpPhoneNumber,
		Password:    tmpPassword,
	}
}

// JobInputFromMap return a JobInput from data map
// use github.com/mitchellh/mapstructure with reflaction
func JobInputFromMap(data map[string]interface{}) (JobInput, error) {
	model := JobInput{}
	err := mapstructure.Decode(data, &model)
	return model, err
}

// MergeToType returns a map with all values set to JobPatch
func (d *JobPatch) MergeToType() map[string]interface{} {
	res := make(map[string]interface{})
	if d.Title != nil {
		res["title"] = *d.Title
	}
	if d.Industry != nil {
		res["industry"] = d.Industry
	}
	if d.Description != nil {
		res["description"] = *d.Description
	}
	if d.Level != nil {
		res["level"] = d.Level
	}
	if d.Location != nil {
		res["location"] = d.Location
	}
	if d.Deadline != nil {
		res["deadline"] = d.Deadline
	}
	if d.EducationLevel != nil {
		res["education_level"] = d.EducationLevel
	}
	if d.Experience != nil {
		res["experience"] = d.Experience
	}
	if d.MinSalary != nil {
		res["min_salary"] = d.MinSalary
	}
	if d.MaxSalary != nil {
		res["max_salary"] = d.MaxSalary
	}
	if d.Requirements != nil {
		res["requirements"] = d.Requirements
	}
	return res
}

// MergeToType retuns a Job filled from JobInput
func (d *JobInput) MergeToType() Job {

	tmpTitle := d.Title

	var tmpIndustry *string
	if d.Industry != nil {
		tmpIndustry = d.Industry
	}

	tmpDescription := d.Description

	var tmpLevel *string
	if d.Level != nil {
		tmpLevel = d.Level
	}

	var tmpLocation *string
	if d.Location != nil {
		tmpLocation = d.Location
	}

	var tmpDeadline *time.Time
	if d.Deadline != nil {
		tmpDeadline = d.Deadline
	}

	var tmpEducationLevel *string
	if d.EducationLevel != nil {
		tmpEducationLevel = d.EducationLevel
	}

	var tmpExperience *int
	if d.Experience != nil {
		tmpExperience = d.Experience
	}

	var tmpMinSalary *int
	if d.MinSalary != nil {
		tmpMinSalary = d.MinSalary
	}

	var tmpMaxSalary *int
	if d.MaxSalary != nil {
		tmpMaxSalary = d.MaxSalary
	}

	var tmpRequirements *string
	if d.Requirements != nil {
		tmpRequirements = d.Requirements
	}
	return Job{
		Title:          tmpTitle,
		Industry:       tmpIndustry,
		Description:    tmpDescription,
		Level:          tmpLevel,
		Location:       tmpLocation,
		Deadline:       tmpDeadline,
		EducationLevel: tmpEducationLevel,
		Experience:     tmpExperience,
		MinSalary:      tmpMinSalary,
		MaxSalary:      tmpMaxSalary,
		Requirements:   tmpRequirements,
	}
}

// SchoolInputFromMap return a SchoolInput from data map
// use github.com/mitchellh/mapstructure with reflaction
func SchoolInputFromMap(data map[string]interface{}) (SchoolInput, error) {
	model := SchoolInput{}
	err := mapstructure.Decode(data, &model)
	return model, err
}

// MergeToType returns a map with all values set to SchoolPatch
func (d *SchoolPatch) MergeToType() map[string]interface{} {
	res := make(map[string]interface{})
	if d.Name != nil {
		res["name"] = *d.Name
	}
	if d.PhoneNumber != nil {
		res["phone_number"] = *d.PhoneNumber
	}
	if d.Password != nil {
		res["password"] = *d.Password
	}
	if d.Badge != nil {
		res["badge"] = d.Badge
	}
	if d.Website != nil {
		res["website"] = d.Website
	}
	return res
}

// MergeToType retuns a School filled from SchoolInput
func (d *SchoolInput) MergeToType() School {

	tmpName := d.Name

	tmpPhoneNumber := d.PhoneNumber

	tmpPassword := d.Password

	var tmpBadge *string
	if d.Badge != nil {
		tmpBadge = d.Badge
	}

	var tmpWebsite *string
	if d.Website != nil {
		tmpWebsite = d.Website
	}
	return School{
		Name:        tmpName,
		PhoneNumber: tmpPhoneNumber,
		Password:    tmpPassword,
		Badge:       tmpBadge,
		Website:     tmpWebsite,
	}
}

// StudentInputFromMap return a StudentInput from data map
// use github.com/mitchellh/mapstructure with reflaction
func StudentInputFromMap(data map[string]interface{}) (StudentInput, error) {
	model := StudentInput{}
	err := mapstructure.Decode(data, &model)
	return model, err
}

// MergeToType returns a map with all values set to StudentPatch
func (d *StudentPatch) MergeToType() map[string]interface{} {
	res := make(map[string]interface{})
	if d.RegistrationNumber != nil {
		res["registration_number"] = *d.RegistrationNumber
	}
	if d.Name != nil {
		res["name"] = *d.Name
	}
	if d.PhoneNumber != nil {
		res["phone_number"] = *d.PhoneNumber
	}
	if d.Password != nil {
		res["password"] = *d.Password
	}
	if d.DateOfAdmission != nil {
		res["date_of_admission"] = d.DateOfAdmission
	}
	if d.DateOfBirth != nil {
		res["date_of_birth"] = d.DateOfBirth
	}
	if d.ProfilePicture != nil {
		res["profile_picture"] = d.ProfilePicture
	}
	return res
}

// MergeToType retuns a Student filled from StudentInput
func (d *StudentInput) MergeToType() Student {

	tmpRegistrationNumber := d.RegistrationNumber

	tmpName := d.Name

	tmpPhoneNumber := d.PhoneNumber

	tmpPassword := d.Password

	var tmpDateOfAdmission *time.Time
	if d.DateOfAdmission != nil {
		tmpDateOfAdmission = d.DateOfAdmission
	}

	var tmpDateOfBirth *time.Time
	if d.DateOfBirth != nil {
		tmpDateOfBirth = d.DateOfBirth
	}

	var tmpProfilePicture *string
	if d.ProfilePicture != nil {
		tmpProfilePicture = d.ProfilePicture
	}
	return Student{
		RegistrationNumber: tmpRegistrationNumber,
		Name:               tmpName,
		PhoneNumber:        tmpPhoneNumber,
		Password:           tmpPassword,
		DateOfAdmission:    tmpDateOfAdmission,
		DateOfBirth:        tmpDateOfBirth,
		ProfilePicture:     tmpProfilePicture,
	}
}

// UnapprovedJobInputFromMap return a UnapprovedJobInput from data map
// use github.com/mitchellh/mapstructure with reflaction
func UnapprovedJobInputFromMap(data map[string]interface{}) (UnapprovedJobInput, error) {
	model := UnapprovedJobInput{}
	err := mapstructure.Decode(data, &model)
	return model, err
}

// MergeToType returns a map with all values set to UnapprovedJobPatch
func (d *UnapprovedJobPatch) MergeToType() map[string]interface{} {
	res := make(map[string]interface{})
	if d.Title != nil {
		res["title"] = *d.Title
	}
	if d.Industry != nil {
		res["industry"] = d.Industry
	}
	if d.Description != nil {
		res["description"] = *d.Description
	}
	if d.Level != nil {
		res["level"] = d.Level
	}
	if d.Location != nil {
		res["location"] = d.Location
	}
	if d.Deadline != nil {
		res["deadline"] = d.Deadline
	}
	if d.EducationLevel != nil {
		res["education_level"] = d.EducationLevel
	}
	if d.Experience != nil {
		res["experience"] = d.Experience
	}
	if d.MinSalary != nil {
		res["min_salary"] = d.MinSalary
	}
	if d.MaxSalary != nil {
		res["max_salary"] = d.MaxSalary
	}
	if d.Requirements != nil {
		res["requirements"] = d.Requirements
	}
	return res
}

// MergeToType retuns a UnapprovedJob filled from UnapprovedJobInput
func (d *UnapprovedJobInput) MergeToType() UnapprovedJob {

	tmpTitle := d.Title

	var tmpIndustry *string
	if d.Industry != nil {
		tmpIndustry = d.Industry
	}

	tmpDescription := d.Description

	var tmpLevel *string
	if d.Level != nil {
		tmpLevel = d.Level
	}

	var tmpLocation *string
	if d.Location != nil {
		tmpLocation = d.Location
	}

	var tmpDeadline *time.Time
	if d.Deadline != nil {
		tmpDeadline = d.Deadline
	}

	var tmpEducationLevel *string
	if d.EducationLevel != nil {
		tmpEducationLevel = d.EducationLevel
	}

	var tmpExperience *int
	if d.Experience != nil {
		tmpExperience = d.Experience
	}

	var tmpMinSalary *int
	if d.MinSalary != nil {
		tmpMinSalary = d.MinSalary
	}

	var tmpMaxSalary *int
	if d.MaxSalary != nil {
		tmpMaxSalary = d.MaxSalary
	}

	var tmpRequirements *string
	if d.Requirements != nil {
		tmpRequirements = d.Requirements
	}
	return UnapprovedJob{
		Title:          tmpTitle,
		Industry:       tmpIndustry,
		Description:    tmpDescription,
		Level:          tmpLevel,
		Location:       tmpLocation,
		Deadline:       tmpDeadline,
		EducationLevel: tmpEducationLevel,
		Experience:     tmpExperience,
		MinSalary:      tmpMinSalary,
		MaxSalary:      tmpMaxSalary,
		Requirements:   tmpRequirements,
	}
}

// UnverifiedAdminInputFromMap return a UnverifiedAdminInput from data map
// use github.com/mitchellh/mapstructure with reflaction
func UnverifiedAdminInputFromMap(data map[string]interface{}) (UnverifiedAdminInput, error) {
	model := UnverifiedAdminInput{}
	err := mapstructure.Decode(data, &model)
	return model, err
}

// MergeToType returns a map with all values set to UnverifiedAdminPatch
func (d *UnverifiedAdminPatch) MergeToType() map[string]interface{} {
	res := make(map[string]interface{})
	if d.Name != nil {
		res["name"] = *d.Name
	}
	if d.PhoneNumber != nil {
		res["phone_number"] = *d.PhoneNumber
	}
	if d.Password != nil {
		res["password"] = *d.Password
	}
	return res
}

// MergeToType retuns a UnverifiedAdmin filled from UnverifiedAdminInput
func (d *UnverifiedAdminInput) MergeToType() UnverifiedAdmin {

	tmpName := d.Name

	tmpPhoneNumber := d.PhoneNumber

	tmpPassword := d.Password
	return UnverifiedAdmin{
		Name:        tmpName,
		PhoneNumber: tmpPhoneNumber,
		Password:    tmpPassword,
	}
}

// UnverifiedSchoolInputFromMap return a UnverifiedSchoolInput from data map
// use github.com/mitchellh/mapstructure with reflaction
func UnverifiedSchoolInputFromMap(data map[string]interface{}) (UnverifiedSchoolInput, error) {
	model := UnverifiedSchoolInput{}
	err := mapstructure.Decode(data, &model)
	return model, err
}

// MergeToType returns a map with all values set to UnverifiedSchoolPatch
func (d *UnverifiedSchoolPatch) MergeToType() map[string]interface{} {
	res := make(map[string]interface{})
	if d.Name != nil {
		res["name"] = *d.Name
	}
	if d.PhoneNumber != nil {
		res["phone_number"] = *d.PhoneNumber
	}
	if d.Password != nil {
		res["password"] = *d.Password
	}
	if d.Badge != nil {
		res["badge"] = d.Badge
	}
	if d.Website != nil {
		res["website"] = d.Website
	}
	return res
}

// MergeToType retuns a UnverifiedSchool filled from UnverifiedSchoolInput
func (d *UnverifiedSchoolInput) MergeToType() UnverifiedSchool {

	tmpName := d.Name

	tmpPhoneNumber := d.PhoneNumber

	tmpPassword := d.Password

	var tmpBadge *string
	if d.Badge != nil {
		tmpBadge = d.Badge
	}

	var tmpWebsite *string
	if d.Website != nil {
		tmpWebsite = d.Website
	}
	return UnverifiedSchool{
		Name:        tmpName,
		PhoneNumber: tmpPhoneNumber,
		Password:    tmpPassword,
		Badge:       tmpBadge,
		Website:     tmpWebsite,
	}
}
