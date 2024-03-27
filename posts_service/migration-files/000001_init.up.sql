CREATE TABLE IF NOT EXISTS posts (
    id serial PRIMARY KEY,
    time_updated timestamp NOT NULL,
    user_id int NOT NULL,
    txt text
);
