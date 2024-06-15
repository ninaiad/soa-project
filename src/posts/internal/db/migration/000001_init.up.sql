CREATE TABLE IF NOT EXISTS posts (
    id bigserial PRIMARY KEY,
    time_updated timestamp NOT NULL,
    user_id bigint NOT NULL,
    txt text
);
