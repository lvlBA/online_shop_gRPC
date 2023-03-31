-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE storage
(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    warehouse_id     uuid NOT NULL REFERENCES warehouse (id),
    warehouse_name   uuid NOT NULL REFERENCES warehouse (name),
    goods_id uuid NOT NULL REFERENCES goods (id),
    max_capacity_weight int NOT NULL,
    current_capacity_weight int NOT NULL,
    created_at timestamp DEFAULT now(),
    changed_at timestamp DEFAULT current_timestamp
);



-- +migrate Down
DROP TABLE storage;