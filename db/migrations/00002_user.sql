-- +goose Up
-- +goose StatementBegin
CREATE TABLE "user" (
    id TEXT PRIMARY KEY,
    "name" TEXT NOT NULL,
    is_active BOOLEAN NOT NULL,
    team_name TEXT NOT NULL REFERENCES team("name")
);

COMMENT ON TABLE "user" IS 'Таблица пользователей';
COMMENT ON COLUMN "user".id IS 'Уникальный идентификатор пользователя';
COMMENT ON COLUMN "user".name IS 'Имя пользователя';
COMMENT ON COLUMN "user".is_active IS 'Активен ли пользователь';
COMMENT ON COLUMN "user".team_name IS 'Команда пользователя';

CREATE INDEX user_team_active_idx ON "user" (team_name, is_active);

COMMENT ON INDEX user_team_active_idx IS 'Индекс для фильтрации активных пользователей по команде';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS user_team_active_idx;
DROP TABLE IF EXISTS "user";
-- +goose StatementEnd