CREATE TABLE employees
(
    id         UUID PRIMARY KEY,
    first_name VARCHAR(255)        NOT NULL,
    last_name  VARCHAR(255)        NOT NULL,
    email      VARCHAR(255) UNIQUE NOT NULL
);
