package model

import (
	"strings"

	"gorm.io/gorm"
)

// returns the appropriate query based on the applied filters
func (j *JobsFilterParameters) Query(query *gorm.DB) *gorm.DB {
	if j.EducationLevel != nil {
		query = query.Where("education_level = ? ", *j.EducationLevel)
	}
	if j.Industry != nil {
		query = query.Where("industry = ? ", *j.Industry)
	}
	if j.Experience != nil {
		query = query.Where("experience <= ? ", *j.Experience)
	}
	return query
}

// returns the appropriate query based on the applied filters
func (e *ApplicationStatus) Query(query *gorm.DB) *gorm.DB {

	if e != nil {
		query = query.Where("status = ? ", strings.ToLower(e.String()))
	}
	return query
}
