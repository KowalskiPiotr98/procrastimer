create table migrations (
    id integer constraint pk_migrations primary key,
    applied_at date not null default now()
);

create function get_current_unix()
    returns bigint as
    $$
    begin
        return extract(epoch from now());
    end
    $$ language plpgsql;

create table users (
    id uuid primary key default gen_random_uuid(),
    email varchar(500) not null constraint ix_users_email unique,
    disabled boolean not null default false
);

create table games (
    id uuid primary key default gen_random_uuid(),
    name varchar(500) not null,
    normalised_name varchar(500) generated always as ( upper(name) ) stored,
    added_on bigint not null default get_current_unix(),
    user_id uuid not null references users(id),

    constraint games_name_user unique (normalised_name, user_id)
);

-- this table doesn't have user permission checks, since it's not really meant to be edited at runtime
create table platforms (
    id uuid primary key default gen_random_uuid(),
    name varchar(200) not null,
    normalised_name varchar(500) constraint platforms_name unique generated always as ( upper(name) ) stored,
    short_name varchar(5) not null constraint platforms_short_name unique
);

insert into platforms (name, short_name) VALUES
    ('PC', 'PC'),
    ('PlayStation 5', 'PS5'),
    ('Xbox Series S/X', 'XBOS'),
    ('Nintendo Switch', 'SWTCH');

create table playthroughs (
    id uuid primary key default gen_random_uuid(),
    game uuid not null references games(id),
    player uuid not null references users(id),
    platform uuid not null references platforms(id),
    start_time bigint not null,
    end_time bigint null check ( end_time is null or end_time > start_time ),
    completed bool null -- true if reached end, false if not, null if not applicable
);
