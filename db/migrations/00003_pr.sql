-- +goose Up
-- +goose StatementBegin
CREATE TABLE pr (
    id TEXT PRIMARY KEY,
    "name" TEXT NOT NULL,
    author_id TEXT NOT NULL REFERENCES "user"(id),
    status SMALLINT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    merged_at TIMESTAMPTZ
);

COMMENT ON TABLE pr IS 'Таблица pull request''ов';
COMMENT ON COLUMN pr.id IS 'Уникальный идентификатор PR';
COMMENT ON COLUMN pr.name IS 'Название PR';
COMMENT ON COLUMN pr.author_id IS 'Автор PR';
COMMENT ON COLUMN pr.status IS 'Статус PR';
COMMENT ON COLUMN pr.created_at IS 'Время создания PR';
COMMENT ON COLUMN pr.merged_at IS 'Время мерджа PR';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS pr;
-- +goose StatementEnd