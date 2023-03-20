-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE auth
(
    id         uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id    uuid  not null REFERENCES users (id),
    token      bytea NOT NULL,
    created_at timestamp        DEFAULT now(),
    changed_at timestamp        DEFAULT current_timestamp
);

CREATE TABLE access
(
    id          uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id     uuid not null REFERENCES users (id),
    resource_id uuid not null REFERENCES resources (id),
    created_at  timestamp        DEFAULT now(),
    changed_at  timestamp        DEFAULT current_timestamp
);


-- +migrate Down
DROP TABLE auth;
DROP TABLE access;

