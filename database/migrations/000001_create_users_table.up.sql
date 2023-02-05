CREATE TABLE IF NOT EXISTS users(
    id uuid PRIMARY KEY default gen_random_uuid(),
    name varchar(255),
    surname varchar(255),
    email varchar(255) unique,
    password_hash varchar unique,
    registered_at timestamp default now(),
    last_visit timestamp
);
