TO DO 

I want to add another relation for users to workouts (one to many)
I want to add testing (adding a workout/user, making sure theres a return)
I want to improve postgresql dump file to pre populate users w/ workouts
I want to add the Dockerfile for the application itself


Packages that need to be installed with go get:

"gorm.io/driver/postgres"
"gorm.io/gorm"

cmd/main contains the file that you actually want to run 
so like :
go run main.go

Right now you can post and get and update and delete a workout 
POSTing to : localhost:3007/workout/
with a body of 

{
    "Name":"Push Ups",
    "Muscle":"Tricep and Chest",
    "Reps":"10"
}

will create a workout 

and a get on localhost:3007/workout/

if you go inside of my pkg folder you can see the structure I was going for 
see the list of routes in pkg/routes/workout-routes.go

the controllers do most of the heavy lifting 





# Goals for Backend

----------------------------------------

## 1. A REST API that: 
### a. Is written in a commonly used language, e.g., Python, Node, Golang, etc. b. 
### b. Has at least one endpoint which: 
i. receives requests as JSON 
ii. records the requests in a database 
### c. Has at least one endpoint which: 
i. Receives requests as query parameters 
ii. Retrieves data from a database 
## 2. A database that: 
a. Is a PostgreSQL Database
b. Is pre-seeded with data 


## QA Reguirements 
80%+ test coverage 

## Devops 
Containerized Applications
---------------------------------------

The Data Model will be as such :

Users have a one to many relationship with Workouts

    Users               Workouts            
  -----------      /   -----------       
    Name          /      Name              
    Email       -----    Muscle
    Title         \      Reps
    Workouts[]       \     


Postgres Running with Docker 
-------------------------------------------

docker run --name postgres-http-db -e POSTGRES_PASSWORD=secretpassword -p 5432:5432 -d postgres

Host: localhost
Port: 5432
User: postgres
Password: secretpassword

inside "/postgres" directory, build docker file with an SQL Dump

docker build -t postgres-db ./

docker run -d --name postgresdb-container -p 5432:5432 postgres-db

sudo docker exec -it <postgresdb-container> psql -U postgres

If you have something conflicting with your port use this 

Find pid of whatever running on 5432 
sudo lsof -i :5432
kill -9 "-PID-"

\l               Will list databases

\c dbname        Will let you use a specific db

\dt              Will list all tables


---------------------------------------------

I have a source.env with the following variables

export DIALECT="postgres"
export HOST="localhost"
export DBPORT="5432"
export USER="postgres"
export NAME="workoutdb"
export PASSWORD="secretpassword"