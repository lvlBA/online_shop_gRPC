-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE carrier
(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name varchar NOT NULL UNIQUE,
    capacity int NOT NULL,
    price float NOT NULL,
    availability bool DEFAULT TRUE,
    created_at timestamp DEFAULT now(),
    changed_at timestamp DEFAULT current_timestamp
);



-- +migrate Down
DROP TABLE carrier;