-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS jobs (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(), -- gen_random_uuid is from cockroachdb
    link VARCHAR(500) UNIQUE,
    description TEXT,
    tech_stack VARCHAR(100)[],
    applied BOOLEAN,
    client VARCHAR(25),
    published_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS jobs CASCADE;
-- +goose StatementEnd
