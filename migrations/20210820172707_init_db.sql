-- +goose Up
CREATE TABLE ozon (
  id BIGSERIAL PRIMARY KEY,
  foo BIGINT NOT NULL
);

-- +goose Down
DROP TABLE ozon;
