-- +goose Up
-- +goose StatementBegin
create or replace function delete_old_otp() returns trigger as $$
begin
  -- delete expired magic links for the user
  delete from magic_links 
  where email = new.email 
    and (id != new.id);
  
  return new;
end;
$$ language plpgsql;

create trigger delete_old_otp_trigger
before insert on magic_links
for each row
execute function delete_old_otp();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop trigger if exists delete_old_otp_trigger on magic_links;
drop function if exists delete_old_otp;
-- +goose StatementEnd