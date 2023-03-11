-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE service
(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    urn text NOT NULL,
    created_at timestamp DEFAULT now(),
    changed_at timestamp
);



-- +migrate Down
DROP TABLE service;


