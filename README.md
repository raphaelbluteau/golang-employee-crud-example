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

# Stopping the Containers

To stop the containers, run:

```bash
make down
```
