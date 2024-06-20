To run the authentication and temperature microservices using Docker Compose, follow these steps:

### Prerequisites
Make sure you have Docker and Docker Compose installed on your machine.

### Steps

1. **Initialize the following variables**
    ```bash
    export DATABASE_USER=*******
    export DATABASE_PASSWORD=*******
    export RABBITMQ_USERNAME=*********
    export RABBITMQ_PASSWORD=************
    ```

2. **Build and Run**

Navigate to the directory where your `docker-compose.yml` file is located and run the following command to build and start the containers:

```bash
docker-compose up -d --build
```

This command will:
- Build Docker images for `authentication`, `temperature`, `frontend` services using the specified Dockerfiles.
- Start containers for MongoDB, RabbitMQ, authentication, temperature, frontend services.
- Expose ports `8080` for the authentication service.
- Expose ports `8081` for the temperature service.
- Expose ports `8082` for the Frontend service

3. **Access the Services**

Once Docker Compose has started the services, you can access them using the following URLs:
- Authentication Service: `http://localhost:8080`
- Temperature Service: `http://localhost:8081`
- Frontend Service: `http://localhost:8082`

4. How to register user
```bash
curl --location 'http://localhost:8080/rbac/api/v1/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email":"haseeb.h002@gmail.com",
    "password":"haseeb",
    "name":"Haseeb Humayun",
    "location": "London"
}'

```
### Stopping the Services

```bash
docker-compose down
```

This command will stop and remove the Docker containers created by Docker Compose.