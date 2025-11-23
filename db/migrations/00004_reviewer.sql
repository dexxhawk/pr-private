-- +goose Up
-- +goose StatementBegin
CREATE TABLE reviewer (
    pr_id TEXT REFERENCES pr(id) ON DELETE CASCADE,
    user_id TEXT REFERENCES "user"(id) ON DELETE CASCADE,
    PRIMARY KEY (pr_id, user_id)
);

COMMENT ON TABLE reviewer IS 'Связь many-to-many между PR и пользователями';
COMMENT ON COLUMN reviewer.pr_id IS 'ID PR';
COMMENT ON COLUMN reviewer.user_id IS 'ID пользователя';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS reviewer;
-- +goose StatementEnd
