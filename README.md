# Eardrum Server    

## Overview  
This project is a school management system that allows multiple schools to host their data, allow restricted access to the data from other parties such as teachers, students and other stakeholders. The sytem also provides various functionalities such as:

* **Creation of a school account**: This project allows schools to register in order to use the system to host their data.

## Table of Contents

* [Technologies](#technologies)
* [Frameworks and Libraries](#frameworks-and-libraries)
* [Third-party Services](#third-party-services)
* [Architecture](#architecture)
* [Database Design](#database-design)
* [Testing and Quality Assurance](#testing-and-quality-assurance)
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

## Database Design     
The database is designed to support the eardum server application managing data for verified and unverified schools.      
### Data Model
A relational database model is used to structure the data, leveraging tables to store information.    
### Table Schema      
#### Schools     
**Purpose**: Stores data for officially verified school accounts.     
**Schema**:
- id (BIGINT, PRIMARY KEY, AUTO_INCREMENT, NOT NULL)
- created_at (DATETIME, NOT NULL)
- updated_at (DATETIME, NOT NULL)
- deleted_at (DATETIME)
- name (TEXT, NOT NULL)
- phone_number (TEXT, NOT NULL, UNIQUE)
- password (TEXT, NOT NULL)
- badge (TEXT)
- website (TEXT)
#### Unverified Schools    
**Purpose**: Stores data for registered but yet to be verified schools
**Schema**:
- id (BIGINT, PRIMARY KEY, AUTO_INCREMENT, NOT NULL)
- created_at (DATETIME, NOT NULL)
- updated_at (DATETIME, NOT NULL)
- deleted_at (DATETIME)
- name (TEXT, NOT NULL)
- phone_number (TEXT, NOT NULL, UNIQUE)
- password (TEXT, NOT NULL)
- badge (TEXT)
- website (TEXT)
### Data Flow     
![DatabaseDataFlow](https://github.com/user-attachments/assets/2be496b2-632f-4250-a892-b7ca4c828c8b)        
## Testing and Quality Assurance   
- **GraphQL Playground**: An in-browser GraphQL IDE that allows you to explore and test your GraphQL API. It includes features like syntax highlighting, auto-complete, 
   query history, and more. GraphQL Playground was used to test query and mutation types. Responses were evaluated if they met the expected results.
## Security Considerations     
* **School Password Hashing** : plaitext passwords provided by schools is encrypted before being stored in the database
* **JSON Web Tokens**: The application leverages JSON Web Tokens to identify users that are logged into the system. Particularly, their id and role. For example: id: "1", 
   role: "school".