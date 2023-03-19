# Employee API

This project provides a CRUD API for managing employee data. It uses Golang, PostgreSQL, Echo framework, and Flyway for database migrations.

## Prerequisites

- Docker and Docker Compose
- Make

## Setup

1. Clone the repository:

```bash
git clone https://github.com/your_username/employee-api.git
cd employee-api
```

2. Create a .env file in the project root to store your PostgreSQL and Flyway credentials:

```bash
cp .env.example .env
```

Edit the .env file and replace the placeholders with your own values:

```text
POSTGRES_USER=your_username
POSTGRES_PASSWORD=your_password
POSTGRES_DB=your_database
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
FLYWAY_POSTGRES_HOST=postgres
FLYWAY_URL=jdbc:postgresql://${FLYWAY_POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}
FLYWAY_USER=${POSTGRES_USER}
FLYWAY_PASSWORD=${POSTGRES_PASSWORD}
```

## Running the Containers

1. Start the PostgreSQL container:

```bash
docker-compose up -d postgres
```

2. Start the Employee API:

```bash
make run
```

The Employee API will be running on http://localhost:1323.

If you want to run Flyway migrations at any time:

```bash
make migrate
```

## Stopping the Containers

To stop the containers, run:

```bash
make down
```

## Running the Load Test

This project includes a load test using Grafana k6, a modern load testing tool. The load test script is located in the loadtest folder.

To run the load test, make sure you have Docker installed and then execute the following command from the root folder of the project:

```bash
make loadtest
```

This command will start the required services using Docker Compose, including the Grafana k6 container, which will run the load test.

Once the load test has finished, the results will be displayed in the terminal. You can also access the load test results in a web interface by opening the URL http://localhost:3000 in your browser. The default username and password are both admin.

For more information about Grafana k6, please visit their official website: https://k6.io/docs/getting-started/what-is-k6.