package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"
	"errors"

	"github.com/GigaDesk/eardrum-graph/neo4jschool"
	"github.com/GigaDesk/eardrum-graph/neo4jstudent"
	"github.com/GigaDesk/eardrum-prefix/prefix"
	"github.com/GigaDesk/eardrum-server/auth"
	"github.com/GigaDesk/eardrum-server/encrypt"
	"github.com/GigaDesk/eardrum-server/graph/model"
	"github.com/GigaDesk/eardrum-server/shutdown"
	"github.com/GigaDesk/eardrum-server/wrappers"
	"github.com/rs/zerolog/log"
)

// AddStudents is the resolver for the AddStudents field.
func (r *mutationResolver) AddStudents(ctx context.Context, students []*model.NewStudent) ([]*model.Student, error) {
	//check if system is in shutdown mode
	if *shutdown.IsShutdown {
		return nil, errors.New("System is shut down for maintainance. We are sorry for any incoveniences caused")
	}
	user := auth.ForContext(ctx)
	if user == nil {
		return nil, errors.New("access to add students denied!")
	}
	role := user.GetRole()
	if role != "school" {
		return nil, errors.New("access to add students denied. Only available for registered and logged in schools. To fix check access token!")
	}

	id, err := user.GetID()

	if err != nil {
		errors.New("could not access school's id!")
	}

	log.Info().Str("role", role).Int("id", id).Str("path", "AddStudents").Msg("adding students")

	var s []*model.Student
	for _, student := range students {
		n := &model.Student{
			RegistrationNumber: prefix.PrefixWithId(student.RegistrationNumber, id),
			Name:               student.Name,
			PhoneNumber:        student.PhoneNumber,
			DateOfAdmission:    student.DateOfAdmission,
			DateOfBirth:        student.DateOfBirth,
			ProfilePicture:     student.ProfilePicture,
		}
		//encrypt student password
		encryptedpassword, err := encrypt.HashPassword(student.Password)
		if err != nil {
			return nil, err
		}
		n.Password = encryptedpassword
		s = append(s, n)
	}
	//Create records in postgres
	if err := r.Sql.Db.Create(s).Error; err != nil {
		log.Error().Str("first_record_name", s[0].Name).Str("path", "AddStudents").Msg(err.Error())
		return nil, errors.New("an unexpected error occurred while creating the school account. please try again later or contact support")
	}
	//Create records in neo4j
	result, err := neo4jschool.CheckSchool(r.Neo4j, id)
	if err != nil {
		go shutdown.InitiateShutdown(err, "AddStudents", id)
		return nil, errors.New("a serious error occurred while adding students. please try again later or contact support")
	}
	if result == false {
		go shutdown.InitiateShutdown(errors.New("school does not exist"), "AddStudents", id)
		return nil, errors.New("a serious synchronization error occurred while adding students. please try again later or contact support")
	}
	for _, student := range s {
		n := wrappers.Neo4jStudentWrapper{
			Student: student,
		}

		if err := neo4jstudent.CreateStudent(r.Neo4j, n, id); err != nil {
			go shutdown.InitiateShutdown(err, "AddStudents", n.ID)
			return nil, errors.New("a serious synchronization error occurred while adding: " + student.Name)
		}
        student.RegistrationNumber = prefix.DePrefixWithId(student.RegistrationNumber, id)
	}
	return s, nil
}
