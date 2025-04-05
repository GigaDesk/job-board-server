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
	"gorm.io/gorm"
)

// Employer is the resolver for the employer field.
func (r *jobProfileResolver) Employer(ctx context.Context, obj *model.JobProfile) (*model.EmployerProfile, error) {
	// declare a job variable
	var job *model.Job

	// Find the first job that matches the id from the job profile
	if err := r.Sql.Db.Where("id = ?", obj.ID).First(&job).Error; err != nil {
		log.Info().Int("id", obj.ID).Str("object", "JobProfile").Msg("job with id does not exist")
		return nil, errors.New("error finding job with id: " + strconv.Itoa(obj.ID))
	}

	if job.EmployerID != nil {
		//declare an employer variable
		var employer *model.Employer

		if err := r.Sql.Db.Where("id = ?", job.EmployerID).First(&employer).Error; err != nil {
			log.Info().Int("id", *job.EmployerID).Str("object", "JobProfile").Msg("employer with id does not exist")
			return nil, errors.New("error finding employer with id: " + strconv.Itoa(*job.EmployerID))
		}

		employerprofile := &model.EmployerProfile{
			ID:          employer.ID,
			CreatedAt:   employer.CreatedAt,
			UpdatedAt:   employer.UpdatedAt,
			Name:        employer.Name,
			PhoneNumber: employer.PhoneNumber,
			Badge:       employer.Badge,
			Website:     employer.Website,
		}

		return employerprofile, nil
	}

	return nil, nil
}

// Applications is the resolver for the applications field.
func (r *jobProfileResolver) Applications(ctx context.Context, obj *model.JobProfile, status *model.ApplicationStatus) ([]*model.ApplicationProfile, error) {
	var applications []model.Application

	query := status.Query(r.Sql.Db.Model(&model.Application{}))

	if err := query.Where("job_id = ?", obj.ID).Find(&applications).Error; err != nil {
		log.Error().Str("object", "JobProfile").Msg(err.Error())
		return nil, errors.New("could not access applications!")
	}
	var applicationprofiles []*model.ApplicationProfile

	for _, application := range applications {
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
		applicationprofiles = append(applicationprofiles, applicationprofile)
	}

	return applicationprofiles, nil
}

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

	log.Info().Str("role", role).Int("id", id).Str("path", "CreateJob").Msg("creating job")

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
		JobURL:         input.JobURL,
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
		JobURL:         n.JobURL,
		Requirements:   strings.Split(pointer.GetString(n.Requirements), "||"),
	}

	return jobprofile, nil
}

// CreateUnapprovedJob is the resolver for the createUnapprovedJob field.
func (r *mutationResolver) CreateUnapprovedJob(ctx context.Context, input model.NewJob) (*model.UnapprovedJobProfile, error) {
	//check if system is in shutdown mode
	if *shutdown.IsShutdown {
		return nil, errors.New("System is shut down for maintainance. We are sorry for any incoveniences caused")
	}

	var employerid *int

	user, _ := auth.ForContext(ctx)

	role := user.GetRole()
	if role == "employer" {
		id, err := user.GetID()

		employerid = &id

		if err != nil {
			return nil, errors.New("could not access employer's id!")
		}

		log.Info().Str("role", role).Int("id", id).Str("path", "CreateUnapprovedJob").Msg("creating unapproved job")
	}

	//Join input.Requiremets with "||" separator

	requirements := strings.Join(input.Requirements, "||")

	n := &model.UnapprovedJob{
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
		EmployerID:     employerid,
		JobURL:         input.JobURL,
	}

	//Create records in postgres
	if err := r.Sql.Db.Create(n).Error; err != nil {
		log.Error().Str("Job Title", n.Title).Str("path", "CreateUnapprovedJob").Msg(err.Error())
		return nil, errors.New("an unexpected error occurred while creating the unapproved job. please try again later or contact support")
	}

	unapprovedjobprofile := &model.UnapprovedJobProfile{
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
		JobURL:         n.JobURL,
		Requirements:   strings.Split(pointer.GetString(n.Requirements), "||"),
	}

	return unapprovedjobprofile, nil
}

// ApproveJob is the resolver for the approveJob field.
func (r *mutationResolver) ApproveJob(ctx context.Context, id int) (*model.JobProfile, error) {
	//check if system is in shutdown mode
	if *shutdown.IsShutdown {
		return nil, errors.New("System is shut down for maintainance. We are sorry for any incoveniences caused")
	}

	user, err := auth.ForContext(ctx)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("access to approve job denied!")
	}
	role := user.GetRole()
	if role != "admin" {
		return nil, errors.New("access to approve job denied. Only available for registered and logged in admins. To fix check access token!")
	}

	Id, err := user.GetID()

	if err != nil {
		errors.New("could not access admin's id!")
	}

	log.Info().Str("role", role).Int("id", Id).Str("path", "ApproveJob").Msg("approving job")

	//declare an unapprovedjob variable
	var unapprovedjob *model.UnapprovedJob

	// Find the first unapproved job that matches the input id from the unapproved job table
	if err := r.Sql.Db.Where("id = ?", id).First(&unapprovedjob).Error; err != nil {
		log.Info().Int("id", id).Str("path", "ApproveJob").Msg("id does not exist")
		return nil, errors.New("error finding unappproved job with id: " + strconv.Itoa(id))
	}

	// transform the unapproved job model into job model and copy it
	job := &model.Job{
		Title:          unapprovedjob.Title,
		Industry:       unapprovedjob.Industry,
		Description:    unapprovedjob.Description,
		Level:          unapprovedjob.Level,
		Location:       unapprovedjob.Location,
		Deadline:       unapprovedjob.Deadline,
		EducationLevel: unapprovedjob.EducationLevel,
		MinSalary:      unapprovedjob.MinSalary,
		MaxSalary:      unapprovedjob.MaxSalary,
		Experience:     unapprovedjob.Experience,
		JobURL:         unapprovedjob.JobURL,
		Requirements:   unapprovedjob.Requirements,
		EmployerID:     unapprovedjob.EmployerID,
	}

	// take the newly transformed and copied job data and transfer it into the official verified job table
	if err := r.Sql.Db.Create(job).Error; err != nil {
		log.Info().Str("Title", job.Title).Str("path", "ApproveJob").Msg(err.Error())
		return nil, errors.New("Failed to approve job. please try again later or contact support")
	}

	// delete the unapproved job from the unapproved job table
	if err := r.Sql.Db.Delete(unapprovedjob).Error; err != nil {
		log.Error().Str("path", "ApprovedJob").Int("record_id", unapprovedjob.ID).Msg(err.Error())
		return nil, errors.New("Failed to complete job approval. please try again later or contact support")
	}

	log.Info().Int("initial_record_id", unapprovedjob.ID).Int("final_record_id", unapprovedjob.ID).Str("path", "ApproveJob").Msg("completed unapproved job to job data transaction")

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

// EditJob is the resolver for the editJob field.
func (r *mutationResolver) EditJob(ctx context.Context, id int, input model.NewJob) (*model.JobProfile, error) {
	//check if system is in shutdown mode
	if *shutdown.IsShutdown {
		return nil, errors.New("System is shut down for maintainance. We are sorry for any incoveniences caused")
	}

	//declare a job variable
	var job *model.Job

	//Find the first job that matches the input id
	if err := r.Sql.Db.Where("id = ?", id).First(&job).Error; err != nil {
		log.Info().Int("id", id).Str("path", "EditJob").Msg("id does not exist")
		return nil, errors.New("error finding job with id: " + strconv.Itoa(id))
	}

	//Join input.Requiremets with "||" separator
	requirements := strings.Join(input.Requirements, "||")

	job.Title = input.Title
	job.Industry = input.Industry
	job.Description = input.Description
	job.Level = input.Level
	job.Location = input.Location
	job.Deadline = input.Deadline
	job.EducationLevel = input.EducationLevel
	job.MinSalary = input.MinSalary
	job.MaxSalary = input.MaxSalary
	job.Experience = input.Experience
	job.JobURL = input.JobURL
	job.Requirements = &requirements

	// Update job record
	if err := r.Sql.Db.Save(&job).Error; err != nil {
		log.Info().Int("id", id).Str("path", "EditJob").Msg("failed to update job")
		return nil, errors.New("error updating job with id: " + strconv.Itoa(id))
	}

	// Find the first job that matches the input id again
	if err := r.Sql.Db.Where("id = ?", id).First(&job).Error; err != nil {
		log.Info().Int("id", id).Str("path", "EditJob").Msg("id does not exist")
		return nil, errors.New("error finding job with id: " + strconv.Itoa(id))
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
		JobURL:         job.JobURL,
		Experience:     job.Experience,
		Requirements:   strings.Split(pointer.GetString(job.Requirements), "||"),
	}

	return jobprofile, nil
}

// RemoveJob is the resolver for the removeJob field.
func (r *mutationResolver) RemoveJob(ctx context.Context, id int) (*model.JobProfile, error) {
	//check if system is in shutdown mode
	if *shutdown.IsShutdown {
		return nil, errors.New("System is shut down for maintainance. We are sorry for any incoveniences caused")
	}

	//check if system is in shutdown mode
	if *shutdown.IsShutdown {
		return nil, errors.New("System is shut down for maintainance. We are sorry for any incoveniences caused")
	}

	//declare a job variable
	var job *model.Job

	// Find the first job that matches the input id from the job table
	if err := r.Sql.Db.Where("id = ?", id).First(&job).Error; err != nil {
		log.Info().Int("id", id).Str("path", "RemoveJob").Msg("id does not exist")
		return nil, errors.New("error finding job with id: " + strconv.Itoa(id))
	}

	// delete the job from the job table
	if err := r.Sql.Db.Delete(job).Error; err != nil {
		log.Error().Str("path", "RemoveJob").Int("record_id", job.ID).Msg(err.Error())
		return nil, errors.New("Failed to complete job removal. please try again later or contact support")
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

// RemoveUnapprovedJob is the resolver for the removeUnapprovedJob field.
func (r *mutationResolver) RemoveUnapprovedJob(ctx context.Context, id int) (*model.UnapprovedJobProfile, error) {
	//check if system is in shutdown mode
	if *shutdown.IsShutdown {
		return nil, errors.New("System is shut down for maintainance. We are sorry for any incoveniences caused")
	}

	//declare an unapproved job variable
	var unapprovedjob *model.UnapprovedJob

	// Find the first unapproved job that matches the input id from the unapproved job table
	if err := r.Sql.Db.Where("id = ?", id).First(&unapprovedjob).Error; err != nil {
		log.Info().Int("id", id).Str("path", "RemoveUnapprovedJob").Msg("id does not exist")
		return nil, errors.New("error finding unapproved job with id: " + strconv.Itoa(id))
	}

	// delete the unapproved job from the unapproved job table
	if err := r.Sql.Db.Delete(unapprovedjob).Error; err != nil {
		log.Error().Str("path", "RemoveUnapprovedJob").Int("record_id", unapprovedjob.ID).Msg(err.Error())
		return nil, errors.New("Failed to complete unapproved job removal. please try again later or contact support")
	}

	unapprovedjobprofile := &model.UnapprovedJobProfile{
		ID:             unapprovedjob.ID,
		CreatedAt:      unapprovedjob.CreatedAt,
		UpdatedAt:      unapprovedjob.UpdatedAt,
		DeletedAt:      unapprovedjob.DeletedAt,
		Title:          unapprovedjob.Title,
		Industry:       unapprovedjob.Industry,
		Description:    unapprovedjob.Description,
		Level:          unapprovedjob.Level,
		Location:       unapprovedjob.Location,
		Deadline:       unapprovedjob.Deadline,
		EducationLevel: unapprovedjob.EducationLevel,
		MinSalary:      unapprovedjob.MinSalary,
		MaxSalary:      unapprovedjob.MaxSalary,
		Experience:     unapprovedjob.Experience,
		JobURL:         unapprovedjob.JobURL,
		Requirements:   strings.Split(pointer.GetString(unapprovedjob.Requirements), "||"),
	}

	return unapprovedjobprofile, nil
}

// EditUnapprovedJob is the resolver for the editUnapprovedJob field.
func (r *mutationResolver) EditUnapprovedJob(ctx context.Context, id int, input *model.NewJob) (*model.UnapprovedJobProfile, error) {
	//check if system is in shutdown mode
	if *shutdown.IsShutdown {
		return nil, errors.New("System is shut down for maintainance. We are sorry for any incoveniences caused")
	}

	//declare an unapproved job variable
	var unapprovedjob *model.UnapprovedJob

	//Find the first unapproved job that matches the input id
	if err := r.Sql.Db.Where("id = ?", id).First(&unapprovedjob).Error; err != nil {
		log.Info().Int("id", id).Str("path", "EditUnapprovedJob").Msg("id does not exist")
		return nil, errors.New("error finding unapproved job with id: " + strconv.Itoa(id))
	}

	//Join input.Requiremets with "||" separator
	requirements := strings.Join(input.Requirements, "||")

	unapprovedjob.Title = input.Title
	unapprovedjob.Industry = input.Industry
	unapprovedjob.Description = input.Description
	unapprovedjob.Level = input.Level
	unapprovedjob.Location = input.Location
	unapprovedjob.Deadline = input.Deadline
	unapprovedjob.EducationLevel = input.EducationLevel
	unapprovedjob.MinSalary = input.MinSalary
	unapprovedjob.MaxSalary = input.MaxSalary
	unapprovedjob.Experience = input.Experience
	unapprovedjob.JobURL = input.JobURL
	unapprovedjob.Requirements = &requirements

	// Update unapproved job record
	if err := r.Sql.Db.Save(&unapprovedjob).Error; err != nil {
		log.Info().Int("id", id).Str("path", "EditUnapprovedJob").Msg("failed to update job")
		return nil, errors.New("error updating unapproved job with id: " + strconv.Itoa(id))
	}

	// Find the first unapprovedjob that matches the input id again
	if err := r.Sql.Db.Where("id = ?", id).First(&unapprovedjob).Error; err != nil {
		log.Info().Int("id", id).Str("path", "EditJob").Msg("id does not exist")
		return nil, errors.New("error finding job with id: " + strconv.Itoa(id))
	}

	jobprofile := &model.UnapprovedJobProfile{
		ID:             unapprovedjob.ID,
		CreatedAt:      unapprovedjob.CreatedAt,
		UpdatedAt:      unapprovedjob.UpdatedAt,
		DeletedAt:      unapprovedjob.DeletedAt,
		Title:          unapprovedjob.Title,
		Industry:       unapprovedjob.Industry,
		Description:    unapprovedjob.Description,
		Level:          unapprovedjob.Level,
		Location:       unapprovedjob.Location,
		Deadline:       unapprovedjob.Deadline,
		EducationLevel: unapprovedjob.EducationLevel,
		MinSalary:      unapprovedjob.MinSalary,
		MaxSalary:      unapprovedjob.MaxSalary,
		JobURL:         unapprovedjob.JobURL,
		Experience:     unapprovedjob.Experience,
		Requirements:   strings.Split(pointer.GetString(unapprovedjob.Requirements), "||"),
	}

	return jobprofile, nil
}

// GetJobs is the resolver for the getJobs field.
func (r *queryResolver) GetJobs(ctx context.Context, filterparameters *model.JobsFilterParameters) ([]*model.JobProfile, error) {
	//check if system is in shutdown mode
	if *shutdown.IsShutdown {
		return nil, errors.New("System is shut down for maintainance. We are sorry for any incoveniences caused")
	}

	var query *gorm.DB

	if filterparameters != nil {
		// extract queries from the filter parameters
		query = filterparameters.Query(r.Sql.Db.Model(&model.Job{}))
	} else {
		query = r.Sql.Db.Model(&model.Job{})
	}

	var jobs []*model.Job

	if err := query.Find(&jobs).Error; err != nil {
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
			JobURL:         job.JobURL,
			Requirements:   strings.Split(pointer.GetString(job.Requirements), "||"),
		}
		jobprofiles = append(jobprofiles, jobprofile)
	}

	return jobprofiles, nil
}

// FindJob is the resolver for the findJob field.
func (r *queryResolver) FindJob(ctx context.Context, id int) (*model.JobProfile, error) {
	//check if system is in shutdown mode
	if *shutdown.IsShutdown {
		return nil, errors.New("System is shut down for maintainance. We are sorry for any incoveniences caused")
	}

	//declare a job variable
	var job *model.Job

	// Find the first job that matches the input id from the job table
	if err := r.Sql.Db.Where("id = ?", id).First(&job).Error; err != nil {
		log.Info().Int("id", id).Str("path", "FindJob").Msg("id does not exist")
		return nil, errors.New("error finding job with id: " + strconv.Itoa(id))
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

// GetUnapprovedJobs is the resolver for the getUnapprovedJobs field.
func (r *queryResolver) GetUnapprovedJobs(ctx context.Context, filterparameters *model.JobsFilterParameters) ([]*model.UnapprovedJobProfile, error) {
	//check if system is in shutdown mode
	if *shutdown.IsShutdown {
		return nil, errors.New("System is shut down for maintainance. We are sorry for any incoveniences caused")
	}

	var query *gorm.DB

	if filterparameters != nil {
		// extract queries from the filter parameters
		query = filterparameters.Query(r.Sql.Db.Model(&model.UnapprovedJob{}))
	} else {
		query = r.Sql.Db.Model(&model.UnapprovedJob{})
	}

	var jobs []*model.UnapprovedJob
	if err := query.Find(&jobs).Error; err != nil {
		log.Error().Str("path", "GetUnapprovedJobs").Msg(err.Error())
		return nil, errors.New("could not access unapproved jobs!")
	}
	log.Info().Str("path", "GetUnapprovedJobs").Msg("getting unapproved jobs")

	var unapprovedjobprofiles []*model.UnapprovedJobProfile

	for _, job := range jobs {
		unapprovedjobprofile := &model.UnapprovedJobProfile{
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
		unapprovedjobprofiles = append(unapprovedjobprofiles, unapprovedjobprofile)
	}

	return unapprovedjobprofiles, nil
}

// FindUnapprovedJob is the resolver for the findUnapprovedJob field.
func (r *queryResolver) FindUnapprovedJob(ctx context.Context, id int) (*model.UnapprovedJobProfile, error) {
	//check if system is in shutdown mode
	if *shutdown.IsShutdown {
		return nil, errors.New("System is shut down for maintainance. We are sorry for any incoveniences caused")
	}

	//declare an unapproved job variable
	var unapprovedjob *model.UnapprovedJob

	// Find the first unapproved job that matches the input id from the unapproved job table
	if err := r.Sql.Db.Where("id = ?", id).First(&unapprovedjob).Error; err != nil {
		log.Info().Int("id", id).Str("path", "FindUnapprovedJob").Msg("id does not exist")
		return nil, errors.New("error finding unapproved job with id: " + strconv.Itoa(id))
	}

	unapprovedjobprofile := &model.UnapprovedJobProfile{
		ID:             unapprovedjob.ID,
		CreatedAt:      unapprovedjob.CreatedAt,
		UpdatedAt:      unapprovedjob.UpdatedAt,
		DeletedAt:      unapprovedjob.DeletedAt,
		Title:          unapprovedjob.Title,
		Industry:       unapprovedjob.Industry,
		Description:    unapprovedjob.Description,
		Level:          unapprovedjob.Level,
		Location:       unapprovedjob.Location,
		Deadline:       unapprovedjob.Deadline,
		EducationLevel: unapprovedjob.EducationLevel,
		MinSalary:      unapprovedjob.MinSalary,
		MaxSalary:      unapprovedjob.MaxSalary,
		Experience:     unapprovedjob.Experience,
		JobURL:         unapprovedjob.JobURL,
		Requirements:   strings.Split(pointer.GetString(unapprovedjob.Requirements), "||"),
	}

	return unapprovedjobprofile, nil
}

// Employer is the resolver for the employer field.
func (r *unapprovedJobProfileResolver) Employer(ctx context.Context, obj *model.UnapprovedJobProfile) (*model.EmployerProfile, error) {
	// declare an unapproved job variable
	var unapprovedjob *model.UnapprovedJob

	// Find the first unapproved job that matches the id from the unapproved job profile
	if err := r.Sql.Db.Where("id = ?", obj.ID).First(&unapprovedjob).Error; err != nil {
		log.Info().Int("id", obj.ID).Str("object", "UnapprovedJobProfile").Msg("unapproved job with id does not exist")
		return nil, errors.New("error finding unapproved job with id: " + strconv.Itoa(obj.ID))
	}

	if unapprovedjob.EmployerID != nil {
		//declare an employer variable
		var employer *model.Employer

		if err := r.Sql.Db.Where("id = ?", unapprovedjob.EmployerID).First(&employer).Error; err != nil {
			log.Info().Int("id", *unapprovedjob.EmployerID).Str("object", "UnapprovedJobProfile").Msg("employer with id does not exist")
			return nil, errors.New("error finding employer with id: " + strconv.Itoa(*unapprovedjob.EmployerID))
		}

		employerprofile := &model.EmployerProfile{
			ID:          employer.ID,
			CreatedAt:   employer.CreatedAt,
			UpdatedAt:   employer.UpdatedAt,
			Name:        employer.Name,
			PhoneNumber: employer.PhoneNumber,
			Badge:       employer.Badge,
			Website:     employer.Website,
		}

		return employerprofile, nil
	}

	return nil, nil
}

// JobProfile returns JobProfileResolver implementation.
func (r *Resolver) JobProfile() JobProfileResolver { return &jobProfileResolver{r} }

// UnapprovedJobProfile returns UnapprovedJobProfileResolver implementation.
func (r *Resolver) UnapprovedJobProfile() UnapprovedJobProfileResolver {
	return &unapprovedJobProfileResolver{r}
}

type jobProfileResolver struct{ *Resolver }
type unapprovedJobProfileResolver struct{ *Resolver }
