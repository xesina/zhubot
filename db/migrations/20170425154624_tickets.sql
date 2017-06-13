-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE zhubot.tickets (
    id             BIGSERIAL PRIMARY KEY,
    chat_id        TEXT        NOT NULL,
    message_id     BIGINT      NOT NULL,
    text           TEXT        NOT NULL,
    status         TEXT        NOT NULL DEFAULT 'open',
    created_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at     TIMESTAMPTZ NOT NULL DEFAULT now()

);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE zhubot.tickets;
