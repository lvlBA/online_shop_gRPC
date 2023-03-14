-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users
(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    first_name text NOT NULL,
    last_name text NOT NULL,
    age int NOT NULL,
    sex text NOT NULL,
    login text NOT NULL UNIQUE,
    hash_password text NOT NULL UNIQUE,
    created_at timestamp DEFAULT now(),
    changed_at timestamp DEFAULT current_timestamp
);

-- +migrate Down
DROP TABLE users;
