# Golang Employee CRUD Example

This is a CRUD API example for employee management written in Golang, utilizing the Echo framework and a PostgreSQL database.

This project and the accompanying loadtest were both created using the GPT-4 language model from OpenAI. The loadtest uses Grafana k6, which is included in the project's Docker configuration.

## Prerequisites

To run this project, you will need to have Docker and Docker Compose installed on your machine.

## Setup

In order to run the application, you will need to create a .env file in the root of the project with the following content:

```text
POSTGRES_USER=your_username
POSTGRES_PASSWORD=your_password
POSTGRES_DB=employee
POSTGRES_HOST=postgres
POSTGRES_PORT=5432
FLYWAY_POSTGRES_HOST=postgres
FLYWAY_URL=jdbc:postgresql://${FLYWAY_POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}
FLYWAY_USER=${POSTGRES_USER}
FLYWAY_PASSWORD=${POSTGRES_PASSWORD}
```

Make sure to replace your_username and your_password with your own values. These environment variables will be used to set up the database connection.

After creating the .env file, you can run the application using the following command:

```bash
make run
```

This will start the PostgreSQL database and run the migration scripts to set up the database schema. Once the database is set up, the application will be started and will listen on port 1323.

## Running load tests

To run the load test, you can use the following command:

```bash
make loadtest
```

This will start the application, wait for it to become available, and then run the load test using Grafana k6. Please note that you do not need to have Grafana k6 installed on your machine, as it will be run in a Docker container.
