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
md social/bin,social/cmd/api,social/cmd/migrate/migrations,social/internal,social/docs,social/scripts,social/web

``` 



## Databases

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
