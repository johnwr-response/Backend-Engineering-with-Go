# Backend-Engineering-with-Go

## Introduction
### Project Overview
### Why Go With Go?
### Preface for Udemy Students
### Course Resources
- Feel free to use the completed project source code below as a reference during your journey. However, I recommend avoiding direct copy-pasting of the code if you want to
  [retain and better understand what you’ve learned](https://www.cpduk.co.uk/news/importance-of-repetition-in-learning) (except for configuration files to avoid errors there).
  The changes in the module code are somewhat separated by commit. Use the commits to guide you. Good luck & Happy learning 💪
- Resources for this lecture: [GitHub Source Code](https://github.com/sikozonpc/GopherSocial)
### Getting Your Tools Ready
- Starting on Module 3. (Scaffolding our API Server) we'll start using Go specific concepts such as:
  - [context](https://pkg.go.dev/context)
  - [interfaces](https://go.dev/tour/methods/9)
  - [error handling](https://go.dev/blog/error-handling-and-go)
  - [pointers](https://go.dev/tour/moretypes/1)
  - [Goroutines](https://go.dev/tour/concurrency/1)
  - [Channels & Maps](https://go.dev/tour/concurrency/1)
- In this course, I won’t be diving into the basics of Go (that’s coming in the future). However, I’ve created separate
  videos on these topics, and I’d recommend going through them to learn the fundamental concepts before moving forward.
- You can also refer back to these resources and explore them as you come across relevant topics during the course.
  That’s what I tend to do — I prefer learning things as I need them.
- Resources:
  - Context: https://youtu.be/Q0BdETrs1Ok
  - Error Handling: https://youtu.be/dKUiCF3abHc
  - Interfaces: https://youtu.be/4OVJ-ir9hL8?si=nZcSoQrTXrYh69y4
  - Maps: https://youtu.be/999h-iyp4Hw?si=fPLtWRs7DWIVBIk-
  - Pointers: https://youtu.be/DVNOP1LE3Mg?si=KXaKeHeIipjLg1HZ
  - Goroutines & Channels: https://youtu.be/3QESpVGiiB8?si=kqpETtKp73Abyiyw

## Project Architecture
### Design Principles for a REST API
- [The Twelve-Factor App](https://12factor.net)
  - I. Codebase : One codebase tracked in revision control, many deploys
  - II. Dependencies : Explicitly declare and isolate dependencies
  - III. Config : Store config in environment
  - IV. Backing services : Treat backing services as attached resources
  - V. Build, release, run : Strictly separate build and run stages
  - VI. Processes : Execute the app as one or more stateless processes
  - VII. Port binding : Export services via port binding
  - VIII. Concurrency : Scale out via the process model
  - IX. Disposability : Maximize robustness with fast startup and graceful shutdown
  - X. Dev/prod parity : Keep development, staging, and production as similar as possible
  - XI. Logs : Treat logs as event streams
  - XII. Admin processes : Run admin/management tasks as one-off processes
- [Roy Fielding REST Dissertation](https://ics.uci.edu/~fielding/pubs/dissertation/fielding_dissertation.pdf)
- [Richardson Maturity Model](https://martinfowler.com/articles/richardsonMaturityModel.html)

## Mini Course: Advanced Go
### Introduction
**What is this module?**
- This is an **optional mini course** inside the Backend Engineering with Go full course.
- It's an introductory material where we **cover the hard parts of Go** before diving into the world of building production web APIs.
- Although Go is a pretty straight forward language to learn, it has some advanced features that are very **important to learn and get them right** from the start.
- And so, in this mini course module we'll get hands-on experience with a project-based approach **by building a Truck Distribution Center simulation**
  which we'll gradually increment with advanced concepts.
- **Disclaimer**: Go basic syntax is not covered on this section, but you can quick and easily learn it in the Go's official docs. (I can record one by popular demand as well)
### Effective Error Handling
### Interfaces
### Testing
### Pointers
- References memory locations
- Modify data in place
- The ampersand (``&``) returns a reference to a memory location
- The star (``*``) dereferences the variable to return the value
### Goroutines
### Context and Timeouts
- Context is immutable! You need to make a new context from a previous context if you want to change it
- It is a convention to send the context as the first argument
### Concurrency with Channels
- NB! Main code not working due to insufficient showing of used methods by instructor
### Maps
- Maps are not concurrency safe
- The concurrency file ran ok every time, but should supposedly create a race condition sometimes
### Capstone Project (Exercise)
### Capstone Project Solution
### Map Concurrency & Mutexes
### Final word
Now that you've got the essential parts of Go covered, **you're ready to start building real awesome projects.** From
the next section forwards we'll start diving deeper into the world of building web APIs with Go. My goal with this
material and the course is for you to learn how to navigate the complex world of building production grade software by
your own. Good luck and any question feel free to drop it in!

## Building a Server from TCP to HTTP
### TCP Server - net package
- Go documentation: [net](https://pkg.go.dev/net)
- TCP is a reliable transmission protocol that runs on top of an unreliable protocol: IP
- TCP offers two guaranties:
  1. Reliable delivery - All packages are acknowledged, and retried if not.
  2. Ordered delivery - All packages are ordered, and retried if missing.
- TCP is a two-way connection protocol. The server and the client.
### Understanding Routing
### HTTP Server - The net/http package
- Go documentation: [http](https://pkg.go.dev/net/http)
### Encoding & Decoding JSON Requests
- Note! Windows Powershell uses an alias to Invoke-WebRequest. If so, switch to newer powershell or use the following:  
``curl -method POST -v http://localhost:8080/users -Headers @{'Accept' = 'application/json'; 'Content-Type' = 'application/json'} -Body '{"first_name":"Tiago", "last_name": "User_123"}'``
- Powershell users can use curl as normal:  
``curl -i 'POST' http://localhost:8080/users -H 'Accept: application/json' -H 'Content-Type: application/json' -d '{"first_name":"Tiago", "last_name": "User_123"}'``
- Using GET without headers is the same for both:  
``curl http://localhost:8080/users``

## Scaffolding our API Server
### Setting up your Development Environment
- [Docker Desktop](https://www.docker.com/products/docker-desktop/)
- [Go API Mini Course for beginners Video](https://www.youtube.com/watch?v=7VLmLOiQ3ck&t=3647s)
- Tech used in this course
  - Go 1.22 or later
  - Docker
  - Postgres running on Docker
  - Swagger for docs
  - Golang migrate for migrations
```
go mod init github.com/johnwr-response/Backend-Engineering-with-Go/social
md social/bin,social/cmd/api,social/cmd/migrate/migrations,social/internal/env,social/docs,social/scripts,social/web
``` 
### Clean Layered Architecture
[Clean Architecture: A Craftsman's Guide to Software Structure and Design (Robert C. Martin Series) 1st Edition](https://www.amazon.com/Clean-Architecture-Craftsmans-Software-Structure/dp/0134494164)
- **Separation of concerns** : Each level in your program should be separated by a clear barrier, the transport layer, the service layer, the storage layer, etc.
- **Dependency Inversion Principle (DIP)**: You're injecting the dependencies in your layers. Tou don't directly call them! Why? It promotes loose coupling and makes it easier to test your programs.
- **Adaptability to Change** : By organizing your code in a modular and flexible way, you can more easily introduce more features, refactor existing code, and respond to evolving business requirements.  
  Your systems should be easy to change, if you have to change a lot of existing code to add a new feature you're doing it wrong.
- **Focus on business value** : Focus on delivering value to your users, they are the ones who will be paying your bills at the end of the month. So focus on the business value.
### Setting up the HTTP server and API
  ``` powershell
  ni social/cmd/api/main.go -type file -Value "package main`n`nfunc main() {`n`n}`n"
  ni social/cmd/api/api.go -type file -Value "package main`n`n"
  ni social/cmd/api/health.go -type file -Value "package main`n`n"
  ```
- Chi - A lightweight, idiomatic and composable router for building Go HTTP services.
  [link](https://github.com/go-chi/chi)
    - Adds no transient dependencies
  ```shell
  cd social
  go get -u github.com/go-chi/chi/v5
  go get -u github.com/go-chi/chi/v5/middleware
  cd ..
  ```
### Hot Reloading in Go
Note! This is and should be optional as it uses and downloads a lot of dependencies.  
-[Air](https://github.com/air-verse/air) - Live reload for Go apps
  ```shell
  cd social
  go install github.com/air-verse/air@latest
go: downloading github.com/air-verse/air v1.63.0
go: downloading github.com/gohugoio/hugo v0.149.1
go: downloading dario.cat/mergo v1.0.2
go: downloading github.com/fatih/color v1.18.0
go: downloading github.com/fsnotify/fsnotify v1.9.0
go: downloading github.com/pelletier/go-toml v1.9.5
go: downloading golang.org/x/sys v0.35.0
go: downloading github.com/mattn/go-colorable v0.1.14
go: downloading golang.org/x/text v0.28.0
go: downloading github.com/bep/godartsass/v2 v2.5.0
go: downloading github.com/spf13/afero v1.14.0
go: downloading github.com/spf13/cast v1.9.2
go: downloading github.com/tdewolff/parse/v2 v2.8.3
go: downloading github.com/pelletier/go-toml/v2 v2.2.4
go: downloading github.com/gobwas/glob v0.2.3
go: downloading github.com/bep/golibsass v1.2.0
go: downloading google.golang.org/protobuf v1.36.8
  cd ..
  ```
- You can initialize the .air.toml configuration file to the current directory with the default settings running the following command:
  ```shell
  air init
  ```
- To use air, simply execute air in the folder 
  ```shell
  cd social
  air
  cd ..
  ```
### Environment Variables
 ``` powershell
  md social/internal/env"
  ni social/internal/env/env.go -type file -Value "package env`n`n"
  ni social/.env -type file -Value "`n"
  ```
- [The Twelve-Factor App Config](https://12factor.net/config)
- [DirEnv](https://direnv.net) : An extension for your shell, augmenting it with a new feature that can load and unload environment variables depending on the current directory.  
  Note: Not yet for windows, even though Powershell is supported
- GoDotEnv - A Go port of Ruby's dotenv library (Loads environment variables from .env files).
  [link](https://github.com/joho/godotenv)
    - Adds no transient dependencies
  ```shell
  cd social
  go get github.com/joho/godotenv
  cd ..
  ```

## Databases
### The Repository Pattern
[The Repository pattern in Go: a painless way to simplify your service logic](https://threedots.tech/post/repository-pattern-in-go/)
### Implementing the Repository Pattern
 ``` powershell
  md social/internal/store"
  ni social/internal/store/storage.go -type file -Value "package store`n`n"
  ni social/internal/store/posts.go -type file -Value "package store`n`n"
  ni social/internal/store/users.go -type file -Value "package store`n`n"
  ```
### Persisting data with SQL
- [sqlx](https://github.com/jmoiron/sqlx) : general purpose extensions to golang's database/sql
    - Adds unchecked number of transient dependencies
- [SQLBoiler](https://github.com/aarondl/sqlboiler) : Generate a Go ORM tailored to your database schema.
    - Adds unchecked number of transient dependencies
- [gorm](https://github.com/go-gorm/gorm) : The fantastic ORM library for Golang, aims to be developer friendly
    - Adds unchecked number of transient dependencies
- [pg](https://github.com/go-pg/pg/v11) : PostgresSQL driver and toolkit for Go
    - Adds the following transient dependencies:
      - `github.com/go-pg/zerochecker`
      - `github.com/jinzhu/inflection`
      - `github.com/tmthrgd/go-hex`
      - `github.com/vmihailenco/msgpack/v5`
      - `github.com/vmihailenco/tagparser`
      - `go.opentelemetry.io/otel`
      - `go.opentelemetry.io/otel/metric`
      - `go.opentelemetry.io/otel/trace`
      - `golang.org/x/sys`
      - `mellium.im/sasl`
- ```shell
  cd social 
  go get github.com/go-pg/pg/v11
  cd ..
  ```
- [pgx](https://github.com/jackc/pgx) : PostgresSQL driver and toolkit for Go
    - Adds the following transient dependencies:
      - `golang.org/x/text`
      - `golang.org/x/crypto`
      - `github.com/pkg/errors`
  ```shell
  cd social 
  go get github.com/jackc/pgx/v5
  cd ..
  ```
### Configuring the DB Connection Pool
 ``` powershell
  md social/internal/db
  ni social/internal/db/db.go -type file -Value "package db`n`n"
  ni social/scripts/db-init.sql -type file -Value "CREATE DATABASE social_network;`n"
  ni social/docker-compose.yaml -type file
  ```
- Run database:
  ```
  cd social
  docker compose up --build
  cd ..
  ```
### SQL Migrations
- [migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) : Database migrations. CLI and Golang library.
- [goose](https://github.com/pressly/goose) : A database migration tool. Supports SQL migrations and Go functions.
- [PostgresSQL CItext](https://www.postgresql.org/docs/current/citext.html) : A case-insensitive character string type
- [GNU Make](https://gnuwin32.sourceforge.net/packages/make.htm)
  - Install: ```winget install -e --id GnuWin32.Make```
- [Create migration]
  ```migrate create -seq -ext sql -dir .\cmd\migrate\migrations\ create_users```
- [Run migration]
  ```migrate -verbose -database "postgres://admin:adminPassword@localhost/social_network?sslmode=disable" -source file://./cmd/migrate/migrations up```
  ```migrate -verbose -database "postgres://admin:adminPassword@localhost/social_network?sslmode=disable" -source file://./cmd/migrate/migrations down```
- [Create makefile]
  ```ni social/Makefile -type file```
- [Run migrations via make]
  ````
  make migrate-down
  make migrate-up
  make migration posts-create
  make migrate-up
  make migration alter-post-table
  make migrate-up
  ````


## Posts CRUD

## User Feed

## Filtering, Sorting, and Pagination

## Documentation

## Structured Logging

## User Creation

## Sending Emails

## Authentication

## Authorization

## Redis Caching

## Testing

## Graceful Shutdown

## Rate Limiting

## Handling CORS

## Server Metrics

## Automation (CI/CD)

## Production Deployment
