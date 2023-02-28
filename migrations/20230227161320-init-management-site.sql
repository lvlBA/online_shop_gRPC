-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE sites
(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name varchar NOT NULL UNIQUE,
    created_at timestamp DEFAULT now(),
    changed_at timestamp DEFAULT current_timestamp
);

-- +migrate StatementBegin

CREATE OR REPLACE FUNCTION update_modified_column()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.changed_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';

-- +migrate StatementEnd


CREATE TRIGGER update_sites_modtime BEFORE UPDATE ON sites FOR EACH ROW EXECUTE PROCEDURE  update_modified_column();

-- +migrate Down
DROP TABLE sites;
DROP FUNCTION update_modified_column;
DROP TRIGGER update_sites_modtime ON sites;
DROP EXTENSION "uuid-ossp";