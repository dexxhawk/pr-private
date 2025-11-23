-- +goose Up
-- +goose StatementBegin
CREATE TABLE team (
    "name" TEXT PRIMARY KEY
);

COMMENT ON TABLE team IS 'Таблица команд';
COMMENT ON COLUMN team.name IS 'Название команды';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS team;
-- +goose StatementEnd
