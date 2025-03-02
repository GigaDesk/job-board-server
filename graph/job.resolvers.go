package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"
	"errors"
	"strings"

	"github.com/AlekSi/pointer"
	"github.com/GigaDesk/eardrum-server/auth"
	"github.com/GigaDesk/eardrum-server/graph/model"
	"github.com/GigaDesk/eardrum-server/shutdown"
	"github.com/rs/zerolog/log"
)

// CreateJob is the resolver for the CreateJob field.
func (r *mutationResolver) CreateJob(ctx context.Context, input model.NewJob) (*model.JobProfile, error) {
	//check if system is in shutdown mode
	if *shutdown.IsShutdown {
		return nil, errors.New("System is shut down for maintainance. We are sorry for any incoveniences caused")
	}
	user, err := auth.ForContext(ctx)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("access to create job denied!")
	}
	role := user.GetRole()
	if role != "admin" {
		return nil, errors.New("access to create job denied. Only available for registered and logged in admins. To fix check access token!")
	}

	id, err := user.GetID()

	if err != nil {
		errors.New("could not access admin's id!")
	}

	log.Info().Str("role", role).Int("id", id).Str("path", "AddStudents").Msg("adding students")

	//Join input.Requiremets with "||" separator

	requirements := strings.Join(input.Requirements, "||")

	n := &model.Job{
		Title:          input.Title,
		Industry:       input.Industry,
		Description:    input.Description,
		Level:          input.Level,
		Location:       input.Location,
		Deadline:       input.Deadline,
		EducationLevel: input.EducationLevel,
		MinSalary:      input.MinSalary,
		MaxSalary:      input.MaxSalary,
		Experience:     input.Experience,
		Requirements:   &requirements,
	}

	//Create records in postgres
	if err := r.Sql.Db.Create(n).Error; err != nil {
		log.Error().Str("Job Title", n.Title).Str("path", "CreateJob").Msg(err.Error())
		return nil, errors.New("an unexpected error occurred while creating the job. please try again later or contact support")
	}

	jobprofile := &model.JobProfile{
		ID:             n.ID,
		CreatedAt:      n.CreatedAt,
		UpdatedAt:      n.UpdatedAt,
		DeletedAt:      n.DeletedAt,
		Title:          n.Title,
		Industry:       n.Industry,
		Description:    n.Description,
		Level:          n.Level,
		Location:       n.Location,
		Deadline:       n.Deadline,
		EducationLevel: n.EducationLevel,
		MinSalary:      n.MinSalary,
		MaxSalary:      n.MaxSalary,
		Experience:     n.Experience,
		Requirements:   strings.Split(pointer.GetString(n.Requirements), "||"),
	}

	return jobprofile, nil
}

// GetJobs is the resolver for the getJobs field.
func (r *queryResolver) GetJobs(ctx context.Context) ([]*model.JobProfile, error) {
	//check if system is in shutdown mode
	if *shutdown.IsShutdown {
		return nil, errors.New("System is shut down for maintainance. We are sorry for any incoveniences caused")
	}

	var jobs []*model.Job
	if err := r.Sql.Db.Find(&jobs).Error; err != nil {
		log.Error().Str("path", "GetJobs").Msg(err.Error())
		return nil, errors.New("could not access jobs!")
	}
	log.Info().Str("path", "GetJobs").Msg("getting jobs")

	var jobprofiles []*model.JobProfile

	for _, job := range jobs {
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
			Requirements:   strings.Split(pointer.GetString(job.Requirements), "||"),
		}
		jobprofiles = append(jobprofiles, jobprofile)
	}

	return jobprofiles, nil
}
