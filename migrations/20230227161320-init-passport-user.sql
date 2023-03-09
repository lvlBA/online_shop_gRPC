-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users
(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    FirstName varchar NOT NULL,
    LastName varchar NOT NULL,
    Age int NOT NULL,
    Sex int,
    Login varchar NOT NULL UNIQUE,
    Password varchar NOT NULL UNIQUE,
    created_at timestamp DEFAULT now(),
    changed_at timestamp DEFAULT current_timestamp
);



-- +migrate Down
DROP TABLE users;
DROP EXTENSION "uuid-ossp";

