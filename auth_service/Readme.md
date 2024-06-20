# Authentication Microservice

## Overview

The Authentication Microservice is responsible for managing user authentication within the system. It supports user registration, login, and JWT-based authentication. The service is implemented using Golang, with Gin for HTTP routing, MongoDB for data storage, and RabbitMQ for messaging.

## Directory Structure

```plaintext
auth_service/
├── Dockerfile
├── Readme.md
├── config.yaml
├── go.mod
├── go.sum
├── internal/
│   ├── di/
│   │   ├── message_queue.go
│   │   ├── repository.go
│   │   └── usecases.go
│   ├── entities/
│   │   └── repository/
│   │       ├── message_queue_mock.go
│   │       ├── message_queue_repo.go
│   │       ├── user_repo.go
│   │       └── user_repo_mock.go
│   ├── handlers/
│   │   ├── user_handler.go
│   │   └── user_handler_test.go
│   ├── infrastrcuture/
│   │   ├── datastore/
│   │   │   ├── mongodb.go
│   │   │   └── user_db.go
│   │   └── message_queue/
│   │       └── message_queue_repo.go
│   ├── middleware/
│   │   └── auth.go
│   └── usecases/
│       ├── user_uc.go
│       └── user_uc_mock.go
├── main.go
└── pkg/
    ├── config/
    │   └── config.go
    └── server/
        ├── server.go
        └── user_routes.go
```

## Packages Description

### `internal/`
Contains the core business logic of the application, organized into several sub-packages.

- **di/**: Dependency Injection configuration.
    - **message_queue.go**: Configures RabbitMQ message queue.
    - **repository.go**: Configures repositories.
    - **usecases.go**: Configures use cases.

- **entities/repository/**: Defines interfaces and mocks for repositories.
    - **message_queue_mock.go**: Mock implementation for message queue repository.
    - **message_queue_repo.go**: Interface for message queue repository.
    - **user_repo.go**: Interface for user repository.
    - **user_repo_mock.go**: Mock implementation for user repository.

- **handlers/**: Contains HTTP handlers for different routes.
    - **user_handler.go**: Implements HTTP handlers for user-related operations like registration and login.
    - **user_handler_test.go**: Unit tests for `user_handler.go`.

- **infrastrcuture/datastore/**: Implements data access layers.
    - **mongodb.go**: MongoDB connection and utility functions.
    - **user_db.go**: MongoDB implementation for user repository.

- **infrastrcuture/message_queue/**: Implements message queue repository.
    - **message_queue_repo.go**: RabbitMQ implementation for message queue repository.

- **middleware/**: Middleware functions.
    - **auth.go**: JWT authentication middleware.

- **usecases/**: Contains business logic and use cases.
    - **user_uc.go**: Implements business logic for user-related operations.
    - **user_uc_mock.go**: Mock implementation for user use cases.

### `pkg/`
Contains shared utility packages.

- **config/**: Manages configuration settings for the application.
    - **config.go**: Loads and provides access to configuration settings using Viper.

- **server/**: Server configuration and route setup.
    - **server.go**: Configures and starts the Gin server.
    - **user_routes.go**: Sets up user-related routes.

### Root Files

- **Dockerfile**: Docker configuration for containerizing the service.
- **Readme.md**: Documentation for the authentication microservice.
- **config.yaml**: Configuration file for the service.
- **go.mod**: Go module dependencies.
- **go.sum**: Go module dependency checksums.
- **main.go**: Entry point for the application.

## Configuration

The service uses Viper to manage configuration. Ensure you have a `config.yaml` file with the following structure:

```yaml
port: 8080
database:
  host: ""
  port: 27017
  user: ""
  password: ""
  db_name: ""
  min_pool: 1
  max_pool: 10
message_queue:
  rabbitmq_uri: ""
  exchange_name: ""
  routing_key: ""
```

## Running the Service

To run the authentication service, execute the following command:

```bash
go run main.go
```

Alternatively, you can build and run the Docker container:

```bash
docker build -t auth-service .
docker run -p 8080:8080 auth-service
```

## Endpoints

The service exposes the following endpoints:

- `POST /register`: Registers a new user.
- `POST /login`: Authenticates a user and returns a JWT token.

## Testing

The service includes unit tests for all use cases and handlers. To run the tests, use the following command:

```bash
go test ./internal/...
```