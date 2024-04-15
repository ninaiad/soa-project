CREATE TABLE IF NOT EXISTS posts (
    id serial PRIMARY KEY,
    timeUpdated timestamp NOT NULL,
    user_id int NOT NULL,
    txt text
);
