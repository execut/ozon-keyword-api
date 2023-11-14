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
    keyword_id BIGINT    NOT NULL REFERENCES keywords,
    type       smallint not null,
    status     smallint,
    payload    jsonb,
    updated timestamp
);

-- +goose Down
DROP TABLE keyword_events;
DROP TABLE keywords;