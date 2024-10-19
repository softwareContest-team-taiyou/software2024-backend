-- +goose Up
-- +goose StatementBegin
ALTER TABLE `users`
ADD COLUMN `is_init` BOOLEAN DEFAULT FALSE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
