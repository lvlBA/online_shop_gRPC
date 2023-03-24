-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE region
(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name varchar NOT NULL UNIQUE,
    site_id     uuid default Null REFERENCES sites (id),
    created_at timestamp DEFAULT now(),
    changed_at timestamp DEFAULT current_timestamp
);



-- +migrate Down
DROP TABLE region;
