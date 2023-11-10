-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE ozon ADD COLUMN created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE ozon DROP COLUMN created;
