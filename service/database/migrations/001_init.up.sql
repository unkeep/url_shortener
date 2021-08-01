create table "urls"
(
    "id" bigserial primary key,
    "url" text,
    "created_at" timestamp not null default current_timestamp
);

create unique index "urls_uidx" on "urls" ("url");