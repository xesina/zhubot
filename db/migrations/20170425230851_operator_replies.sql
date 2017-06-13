-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE zhubot.operator_replies (
    id             BIGSERIAL PRIMARY KEY,
    operator_id    BIGINT      NOT NULL REFERENCES zhubot.operators (id) ON DELETE CASCADE ON UPDATE CASCADE,
    operator_reply TEXT,
    message_id     BIGINT      NOT NULL,
    ticket_id      BIGINT      NOT NULL REFERENCES zhubot.tickets (id) ON DELETE CASCADE ON UPDATE CASCADE,
    status         TEXT        NOT NULL DEFAULT 'open',
    created_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at     TIMESTAMPTZ NOT NULL DEFAULT now()

);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE zhubot.operator_replies;
