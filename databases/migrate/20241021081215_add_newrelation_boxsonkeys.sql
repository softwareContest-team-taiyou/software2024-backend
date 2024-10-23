-- +goose Up
-- +goose StatementBegin
ALTER TABLE `boxs_on_keys`
ADD PRIMARY KEY (`boxs_id`, `keys_id`);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
