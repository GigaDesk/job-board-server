package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/AlekSi/pointer"
	"github.com/GigaDesk/eardrum-server/auth"
	"github.com/GigaDesk/eardrum-server/graph/model"
	"github.com/GigaDesk/eardrum-server/shutdown"
	"github.com/rs/zerolog/log"
)

// Job is the resolver for the job field.
func (r *applicationProfileResolver) Job(ctx context.Context, obj *model.ApplicationProfile) (*model.JobProfile, error) {
	// declare an application variable
	var application *model.Application

	// Find the first application that matches the id from the application profile
	if err := r.Sql.Db.Where("id = ?", obj.ID).First(&application).Error; err != nil {
		log.Info().Int("id", obj.ID).Str("object", "ApplicationProfile").Msg("application with id does not exist")
		return nil, errors.New("error finding application with id: " + strconv.Itoa(obj.ID))
	}

	//declare an job variable
	var job *model.Job

	if err := r.Sql.Db.Where("id = ?", application.JobID).First(&job).Error; err != nil {
		log.Info().Int("id", application.JobID).Str("object", "ApplicationProfile").Msg("job with id does not exist")
		return nil, errors.New("error finding job with id: " + strconv.Itoa(application.JobID))
	}

	jobprofile := &model.JobProfile{
		ID:             job.ID,
		CreatedAt:      job.CreatedAt,
		UpdatedAt:      job.UpdatedAt,
		DeletedAt:      job.DeletedAt,
		Title:          job.Title,
		Industry:       job.Industry,
		Description:    job.Description,
		Level:          job.Level,
		Location:       job.Location,
		Deadline:       job.Deadline,
		EducationLevel: job.EducationLevel,
		MinSalary:      job.MinSalary,
		MaxSalary:      job.MaxSalary,
		Experience:     job.Experience,
		JobURL:         job.JobURL,
		Requirements:   strings.Split(pointer.GetString(job.Requirements), "||"),
	}

	return jobprofile, nil
}

// Employee is the resolver for the employee field.
func (r *applicationProfileResolver) Employee(ctx context.Context, obj *model.ApplicationProfile) (*model.EmployeeProfile, error) {
	// declare an application variable
	var application *model.Application

	// Find the first application that matches the id from the application profile
	if err := r.Sql.Db.Where("id = ?", obj.ID).First(&application).Error; err != nil {
		log.Info().Int("id", obj.ID).Str("object", "ApplicationProfile").Msg("application with id does not exist")
		return nil, errors.New("error finding application with id: " + strconv.Itoa(obj.ID))
	}

	//declare an employee variable
	var employee *model.Employee

	if err := r.Sql.Db.Where("id = ?", application.EmployeeID).First(&employee).Error; err != nil {
		log.Info().Int("id", application.EmployeeID).Str("object", "ApplicationProfile").Msg("employee with id does not exist")
		return nil, errors.New("error finding employee with id: " + strconv.Itoa(application.EmployeeID))
	}

	employeeprofile := &model.EmployeeProfile{
		ID:             employee.ID,
		CreatedAt:      employee.CreatedAt,
		UpdatedAt:      employee.UpdatedAt,
		Name:           employee.Name,
		PhoneNumber:    employee.PhoneNumber,
		Profilepicture: employee.Profilepicture,
	}

	return employeeprofile, nil
}

// CreateApplication is the resolver for the createApplication field.
func (r *mutationResolver) CreateApplication(ctx context.Context, input model.NewApplication) (*model.ApplicationProfile, error) {
	//check if system is in shutdown mode
	if *shutdown.IsShutdown {
		return nil, errors.New("System is shut down for maintainance. We are sorry for any incoveniences caused")
	}
	user, err := auth.ForContext(ctx)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("access to create application denied!")
	}
	role := user.GetRole()
	if role != "employee" {
		return nil, errors.New("access to create application denied. Only available for registered and logged in employees. To fix check access token!")
	}

	id, err := user.GetID()

	if err != nil {
		errors.New("could not access employee's id!")
	}

	log.Info().Str("role", role).Int("id", id).Str("path", "CreateApplication").Msg("creating application")

	// create application data

	n := &model.Application{
		JobID:          input.JobID,
		EmployeeID:     id,
		EducationLevel: input.EducationLevel,
		Experience:     input.Experience,
		CoverLetterURL: input.CoverLetterURL,
		ResumeeURL:     input.ResumeeURL,
		Status:         "pending",
	}

	//Create records in postgres
	if err := r.Sql.Db.Create(n).Error; err != nil {
		log.Error().Int("JobId", n.JobID).Str("path", "CreateApplication").Msg(err.Error())
		return nil, errors.New("an unexpected error occurred while creating the application. please try again later or contact support")
	}

	applicationprofile := &model.ApplicationProfile{
		ID:             n.ID,
		CreatedAt:      n.CreatedAt,
		UpdatedAt:      n.UpdatedAt,
		DeletedAt:      n.DeletedAt,
		EducationLevel: n.EducationLevel,
		Experience:     n.Experience,
		CoverLetterURL: n.CoverLetterURL,
		ResumeeURL:     n.ResumeeURL,
		Status:         model.ApplicationStatus(strings.ToUpper(n.Status)),
	}

	return applicationprofile, nil
}

// EditApplication is the resolver for the editApplication field.
func (r *mutationResolver) EditApplication(ctx context.Context, id int, input *model.NewApplication, status *model.ApplicationStatus) (*model.ApplicationProfile, error) {
	//check if system is in shutdown mode
	if *shutdown.IsShutdown {
		return nil, errors.New("System is shut down for maintainance. We are sorry for any incoveniences caused")
	}

	log.Info().Str("path", "EditApplication").Msg("editing application")

	//declare an application variable
	var application *model.Application

	//Find the first application that matches the input id
	if err := r.Sql.Db.Where("id = ?", id).First(&application).Error; err != nil {
		log.Info().Int("id", id).Str("path", "EditApplication").Msg("id does not exist")
		return nil, errors.New("error finding application with id: " + strconv.Itoa(id))
	}

	application.Status = strings.ToLower(status.String())

	// Update application record
	if err := r.Sql.Db.Save(&application).Error; err != nil {
		log.Info().Int("id", id).Str("path", "EditApplication").Msg("failed to update application")
		return nil, errors.New("error updating application with id: " + strconv.Itoa(id))
	}

	// Find the first application that matches the input id again
	if err := r.Sql.Db.Where("id = ?", id).First(&application).Error; err != nil {
		log.Info().Int("id", id).Str("path", "EditApplication").Msg("id does not exist")
		return nil, errors.New("error finding application with id: " + strconv.Itoa(id))
	}

	applicationprofile := &model.ApplicationProfile{
		ID:             application.ID,
		CreatedAt:      application.CreatedAt,
		UpdatedAt:      application.UpdatedAt,
		DeletedAt:      application.DeletedAt,
		EducationLevel: application.EducationLevel,
		Experience:     application.Experience,
		CoverLetterURL: application.CoverLetterURL,
		ResumeeURL:     application.ResumeeURL,
		Status:         model.ApplicationStatus(strings.ToUpper(application.Status)),
	}

	return applicationprofile, nil
}

// FindApplication is the resolver for the findApplication field.
func (r *queryResolver) FindApplication(ctx context.Context, id int) (*model.ApplicationProfile, error) {
	//check if system is in shutdown mode
	if *shutdown.IsShutdown {
		return nil, errors.New("System is shut down for maintainance. We are sorry for any incoveniences caused")
	}

	//declare an application variable
	var application *model.Application

	// Find the first application that matches the input id from the application table
	if err := r.Sql.Db.Where("id = ?", id).First(&application).Error; err != nil {
		log.Info().Int("id", id).Str("path", "FindApplication").Msg("id does not exist")
		return nil, errors.New("error finding application with id: " + strconv.Itoa(id))
	}

	applicationprofile := &model.ApplicationProfile{
		ID:             application.ID,
		CreatedAt:      application.CreatedAt,
		UpdatedAt:      application.UpdatedAt,
		DeletedAt:      application.DeletedAt,
		EducationLevel: application.EducationLevel,
		Experience:     application.Experience,
		CoverLetterURL: application.CoverLetterURL,
		ResumeeURL:     application.ResumeeURL,
		Status:         model.ApplicationStatus(strings.ToUpper(application.Status)),
	}

	return applicationprofile, nil
}

// ApplicationProfile returns ApplicationProfileResolver implementation.
func (r *Resolver) ApplicationProfile() ApplicationProfileResolver {
	return &applicationProfileResolver{r}
}

type applicationProfileResolver struct{ *Resolver }
