-- +goose Up
-- +goose StatementBegin
ALTER TABLE todos
ADD COLUMN user_id CHAR(36),
ADD CONSTRAINT fk_user
  FOREIGN KEY (user_id)
  REFERENCES users(id)
  ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
