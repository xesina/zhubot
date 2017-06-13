-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE zhubot.operators (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    chat_id BIGINT NOT NULL UNIQUE,
    active BOOL DEFAULT FALSE
);

INSERT INTO zhubot.operators(id, name, chat_id, active) VALUES (1, 'Sina', 52128276, TRUE);
INSERT INTO zhubot.operators(id, name, chat_id, active) VALUES (2, 'Test', 360368311, TRUE);
INSERT INTO zhubot.operators(id, name, chat_id) VALUES (3, 'Tina', 215400528);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE zhubot.operators;


