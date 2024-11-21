# Eardrum Server    

## Overview  
This project is a school management system that allows multiple schools to host their data, allow restricted access to the data from other parties such as teachers, students and other stakeholders. The sytem also provides various functionalities such as:

* **Creation of a school account**: This project allows schools to register in order to use the system to host their data.

## Table of Contents

* [Technologies](#technologies)
* [Frameworks and Libraries](#frameworks-and-libraries)
* [Third-party Services](#third-party-services)
* [Architecture](#architecture)
* [Deployment Guide](#deployment-guide)
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
## Deployment Guide
* Install [gcloud CLI](https://cloud.google.com/sdk/docs/install)
* `gcloud init` to link up your gcloud CLI with your google cloud project
* Ensure you have installed Docker engine on your local computer. If you haven't install [here](https://docs.docker.com/engine/install/)
* Clone the eardrum-server repo into your local machine
* Build a docker-image off your preffered eardrum-server version or release
* Setup an Artifact Registry Docker repository in Google Cloud. Follow this [guide](https://www.youtube.com/watch?v=b-rg71O3ZwU&t=9s)
* `gcloud auth configure-docker` to configure authentication to Artifact Registry for Docker. For more information [here](https://cloud.google.com/artifact-registry/docs/docker/authentication)
* Tag the local image that you build using this commands
```
docker tag [IMAGE_NAME] \
gcr.io/[PROJECT_ID]/[REPO_NAME]/[IMAGE_NAME]:[TAG]
```
* Next, push the tagged image to Artifact Registry by running:
```
docker push gcr.io/[PROJECT_ID]/[REPO_NAME]/[IMAGE_NAME]:[TAG]
```
If it fails for some reason run `gcloud auth configure-docker` and then push it again
* Now go to Google Cloud Run console > Navigate to "DEPLOY CONTAINER" > pick "Service" option > select where it says "Container image URL" > Select Artifact Registry > Select your repository and image and deploy.
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
