-- +goose Up
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS applied_idx ON jobs (applied);
CREATE INDEX IF NOT EXISTS published_at_idx ON jobs (published_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS jobs@applied_idx;
DROP INDEX IF EXISTS jobs@published_at_idx;
-- +goose StatementEnd
