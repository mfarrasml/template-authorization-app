-- Create new database
CREATE DATABASE authorization_test_db;

-- DDL
CREATE TABLE users(
    id BIGSERIAL,
    user_name VARCHAR,
    email VARCHAR UNIQUE,
    password VARCHAR,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

-- DML
INSERT INTO users(user_name, email, password)
VALUES ('John Doe', 'john.doe@mail.com', '$2a$12$ahXEFxFGDIeO3QC5CfB/DuO0EQ8W60KsLGIYkzgX3Bt3luiE0rdUy');