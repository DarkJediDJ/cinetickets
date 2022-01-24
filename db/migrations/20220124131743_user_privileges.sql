-- +goose Up
CREATE TABLE IF NOT EXISTS public.user_privileges
(
    user_id integer NOT NULL,
    privilege_id integer NOT NULL,
    CONSTRAINT "FK_user_privileges_to_privileges" FOREIGN KEY (privilege_id)
        REFERENCES public.privileges (privilege_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT "FK_user_privileges_to_users" FOREIGN KEY (user_id)
        REFERENCES public.users (user_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

-- +goose Down
DROP TABLE public.user_privileges;
