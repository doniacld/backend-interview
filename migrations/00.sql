CREATE TABLE users (
    id   varchar(264),
    name varchar(264)
);

CREATE TABLE account (
    id      varchar(264),
    user_id varchar(264),
    total double precision
);

CREATE TABLE transaction (
    id         varchar(264),
    amount     double precision,
    account_id varchar(264),
    created_at integer
);
