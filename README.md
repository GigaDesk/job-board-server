# Eardrum Server    

## Overview  
This project is a school management system that allows multiple schools to host their data, allow restricted access to the data from other parties such as teachers, students and other stakeholders. The sytem also provides various functionalities such as:

* **Creation of a school account**: This project allows schools to register in order to use the system to host their data.

## Table of Contents

* [Technologies](#technologies)
* [Frameworks and Libraries](#frameworks-and-libraries)
* [Third-party Services](#third-party-services)
* [Architecture](#architecture)
* [Security Considerations](#security-considerations)

## Technologies 

- Golang
- GraphQL
- Postgres

## Frameworks and Libraries

- gqlgen
- autogql
- godotenv
- twilio-go
- gorm

## Third-party Services 

- Twilio verify

## Architecture
![GigaDeskERD](https://github.com/user-attachments/assets/fe97da0c-7f0b-4c0d-96e8-047413dd2d42)     

## Security Considerations     
* **School Password Hashing** : plaitext passwords provided by schools is encrypted before being stored in the database
