-- +goose Up
create table items (
     part_uuid varchar(36) not null,
     created_at timestamp not null default now(),
     updated_at timestamp not null default now()
);

-- +goose Down
drop table order_items;