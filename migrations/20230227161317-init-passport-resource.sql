-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE resources
(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    urn text NOT NULL UNIQUE,
    created_at timestamp DEFAULT now(),
    changed_at timestamp DEFAULT current_timestamp
);



-- +migrate Down
DROP TABLE resources;


