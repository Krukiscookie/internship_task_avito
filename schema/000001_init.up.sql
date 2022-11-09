CREATE TABLE accounts
(
    id serial not null unique,
    balance numeric,
    reserve numeric DEFAULT 0
    CONSTRAINT not_neg_balance CHECK (balance>=0 and reserve>=0)
);

CREATE TABLE transaction
(
    id serial not null unique,
    id_from int references accounts(id),
    id_to int references accounts(id),
    amount numeric,
    status varchar(255) not null,
    create_time timestamp
);

CREATE TABLE services
(
    id serial PRIMARY KEY,
    account_id int not null references accounts(id),
    amount numeric,
    service_id int not null,
    order_id int not null,
    status varchar(255) not null,
    created_at timestamp,
    updated_at timestamp,
    UNIQUE (service_id, order_id)
);