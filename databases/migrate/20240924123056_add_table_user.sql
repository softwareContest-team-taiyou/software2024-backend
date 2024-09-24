-- +goose Up
-- +goose StatementBegin
-- `users` テーブルを作成
CREATE TABLE users (
    id CHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- `todos` テーブルに `user_id` カラムを追加し、`users` とのリレーションを作成

-- +goose StatementEnd
