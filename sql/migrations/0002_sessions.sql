-- +goose Up
-- +goose StatementBegin
create table if not exists sessions (
  token uuid primary key default uuid_generate_v4(),
  user_id uuid not null references users(id),
  expires_at timestamp not null default now() + interval '90 days'
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists sessions;
-- +goose StatementEnd
