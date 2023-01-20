-- +migrate Up
CREATE TABLE orders(
    id text,
    src_chain text,
    account char(42) NOT NULL,
    sell_tokens char(42) NOT NULL,
    buy_tokens char(42) NOT NULL,
    sell_amount text NOT NULL,
    buy_amount text NOT NULL,
    dest_chain text NOT NULL,

    state smallint NOT NULL,
    executed_by text NOT NULL,
    match_swapica text NOT NULL,
    PRIMARY KEY (id, src_chain)
);

CREATE TABLE last_blocks(
    id text NOT NULL,
    src_chain text PRIMARY KEY
);

-- +migrate Down
DROP TABLE orders;
DROP TABLE last_blocks;
