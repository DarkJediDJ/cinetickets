-- +goose Up
CREATE TABLE IF NOT EXISTS public.users
(
    email text NOT NULL,
    password text NOT NULL,
    id SERIAL,
    CONSTRAINT users_pkey PRIMARY KEY (id)
);


-- +goose Down
DROP TABLE public.users;
