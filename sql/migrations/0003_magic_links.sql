-- +goose Up
-- +goose StatementBegin
create table if not exists magic_links (
  id serial primary key,
  email varchar(255) not null unique,
  otp varchar(6) not null,
  is_used boolean not null default false,
  created_at timestamp not null default now(),
  expires_at timestamp not null default now() + interval '10 minutes'
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists magic_links;
-- +goose StatementEnd