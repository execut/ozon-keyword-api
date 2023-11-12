-- +goose Up
CREATE TABLE keywords
(
    id      BIGSERIAL PRIMARY KEY,
    name    varchar(1000) NOT NULL,
    removed boolean       NOT NULL NOT NULL,
    created timestamp     NOT NULL,
    updated timestamp
);

CREATE TABLE keyword_events
(
    id         BIGSERIAL PRIMARY KEY,
    keyword_id BIGSERIAL    NOT NULL REFERENCES keywords,
    type       varchar(255) not null,
    status     varchar(255),
    payload    jsonb,
    updated timestamp
);

-- +goose Down
DROP TABLE keyword_events;
DROP TABLE keywords;