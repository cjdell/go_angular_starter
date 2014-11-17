
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO users (name, email, hash) VALUES ('Administrator', 'admin@example.com', '5f4dcc3b5aa765d61d8327deb882cf99');


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM users WHERE email = 'admin@example.com';
