-- +goose Up
-- +goose StatementBegin
CREATE TABLE blogs (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    Author TEXT NOT NULL,
    created_date TIMESTAMP NOT NULL,
    content TEXT NOT NULL,
    userID UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
