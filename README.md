# Project Title

## Golang Web Application Starter

## Description

This repository contains a startup build for a web application using Golang. The architecture is well-structured, making it easy to understand and extend. The application utilizes several Golang packages to manage configuration, handle routes, perform template rendering, and manage sessions.

## Overview

Main Packages and Technologies Used

1. `net/http `for handling HTTP requests and responses.
2. `github.com/alexedwards/scs/v2` for session management.
3. `github.com/go-chi/chi`for routing.
4. `html/template` for HTML templating.
5. `github.com/justinas/nosurf` for CSRF protection.

## Application Structure

1. `main.go`: This is the entry point of the application. It sets up the server, session manager, and loads the templates.
2. `middleware.go`: Contains middleware functions like CSRF protection and session management.
3. `routes.go`: Specifies the application's routes and corresponding handler functions.
4. `pkg/config/config.go`: Holds the application's configuration settings.
5. `pkg/handlers/handlers.go`: Contains the handler functions that deal with HTTP requests and responses.
6. `pkg/models/templatedata.go`: Defines the structure of data to be passed into templates.
7. `pkg/render/render.go`: Handles rendering of HTML templates.

## Running the Program

1. Clone this repository: git clone `https://github.com/GiorgiMakharadze/web-application-starter-golang.git`
2. Navigate to the project directory.
3. Install dependencies: `go get ./...`
4. `html/template` for HTML templates
5. To run the application, execute: `go run cmd/web/*.go`
