-- +goose Up 
CREATE TABLE IF NOT EXISTS users (
    id INT PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    name TEXT UNIQUE
);

-- +goose Down
DROP TABLE users;
