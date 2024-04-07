-- Create new database
CREATE DATABASE user_db;

-- DDL
CREATE TABLE users(
    id BIGSERIAL,
    user_name VARCHAR,
    email VARCHAR,
    password VARCHAR,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

-- DML
UPDATE users
SET user_name='John Doe', 
    email='john.doe@mail.com', 
    password='$2a$12$W3/W6PWvPnWVjTCtxFiQn.7yRLiwn.ds1MPc.O6dD0b9gUV3puZ8S';