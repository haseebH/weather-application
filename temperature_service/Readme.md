# Temperature Microservice

## Overview

The Temperature Microservice retrieves and manages air temperature data for a user's hometown over specified durations (past month, year, or 3 years). This service is built using Golang, utilizing Gin for HTTP routing, MongoDB for data storage, RabbitMQ for messaging, and external weather APIs for fetching weather information.

## Directory Structure

```plaintext
temperature_service/
├── Dockerfile
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
│   │       ├── message_queue_repo.go
│   │       ├── temperature_repo.go
│   │       └── user_repo.go
│   ├── handlers/
│   │   └── temperature_handler.go
│   ├── infrastrcuture/
│   │   ├── datastore/
│   │   │   ├── mongodb.go
│   │   │   └── temperature_db.go
│   │   └── message_queue/
│   │       └── message_queue_repo.go
│   ├── middleware/
│   │   └── auth.go
│   ├── usecases/
│   │   ├── temperature_uc.go
│   │   └── weather_uc.go
│   └── utils/
│       └── fetch_weather_data.go
├── main.go
└── pkg/
    ├── config/
    │   └── config.go
    └── server/
        ├── server.go
        └── temperature_routes.go
```

## Packages Description

### `internal/`
Contains the core business logic of the application, organized into several sub-packages.

- **di/**: Dependency Injection configuration.
    - **message_queue.go**: Configures RabbitMQ message queue.
    - **repository.go**: Configures repositories.
    - **usecases.go**: Configures use cases.

- **entities/repository/**: Defines interfaces for repositories.
    - **message_queue_repo.go**: Interface for message queue repository.
    - **temperature_repo.go**: Interface for temperature repository.
    - **user_repo.go**: Interface for user repository.

- **handlers/**: Contains HTTP handlers for different routes.
    - **temperature_handler.go**: Implements HTTP handlers for temperature-related operations.

- **infrastrcuture/datastore/**: Implements data access layers.
    - **mongodb.go**: MongoDB connection and utility functions.
    - **temperature_db.go**: MongoDB implementation for temperature repository.

- **infrastrcuture/message_queue/**: Implements message queue repository.
    - **message_queue_repo.go**: RabbitMQ implementation for message queue repository.

- **middleware/**: Middleware functions.
    - **auth.go**: JWT authentication middleware.

- **usecases/**: Contains business logic and use cases.
    - **temperature_uc.go**: Implements business logic for retrieving temperature data.
    - **weather_uc.go**: Implements logic for fetching weather data from external APIs.

- **utils/**: Utility functions.
    - **fetch_weather_data.go**: Functionality to fetch weather data from external APIs.

### `pkg/`
Contains shared utility packages.

- **config/**: Manages configuration settings for the application.
    - **config.go**: Loads and provides access to configuration settings using Viper.

- **server/**: Server configuration and route setup.
    - **server.go**: Configures and starts the Gin server.
    - **temperature_routes.go**: Sets up temperature-related routes.

### Root Files

- **Dockerfile**: Docker configuration for containerizing the service.
- **config.yaml**: Configuration file for the service.
- **go.mod**: Go module dependencies.
- **go.sum**: Go module dependency checksums.
- **main.go**: Entry point for the application.

## Configuration

Ensure you have a `config.yaml` file with the following structure to configure the service:

```yaml
port: 8081
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
  queue: ""

weather_apikey: ""
weather_api_baseurl: "https://api.open-meteo.com/v1/forecast"
```

## Running the Service

To run the temperature service, execute the following command:

```bash
go run main.go
```

Alternatively, build and run the Docker container:

```bash
docker build -t temperature-service .
docker run -p 8080:8080 temperature-service
```

## Endpoints

The service exposes the following endpoints:

- `GET /temperature/:location/:period`: Retrieves air temperature data for the past month/ past year/ past 3 years.

## Testing

The service includes unit tests for all use cases and handlers. To run the tests, use the following command:

```bash
go test ./internal/...
```

