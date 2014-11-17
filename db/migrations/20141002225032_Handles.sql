
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE products ADD COLUMN handle varchar(200) DEFAULT '';
ALTER TABLE categories ADD COLUMN handle varchar(200) DEFAULT '';

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE categories DROP COLUMN handle;
ALTER TABLE products DROP COLUMN handle;
