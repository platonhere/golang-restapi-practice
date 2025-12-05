CREATE TABLE users (
    id bigserial PRIMARY KEY,
    email varchar NOT NULL UNIQUE,
    encrypted_password varchar NOT NULL
);