create table reelplay.users
(
    id         bigint auto_increment
        primary key,
    first_name longtext    null,
    last_name  longtext    null,
    email      longtext    null,
    password   longtext    null,
    created_at datetime(3) null,
    deleted_at datetime(3) null,
    role_id    bigint      null,
    constraint fk_users_role
        foreign key (role_id) references reelplay.roles (id)
);

create index idx_users_deleted_at
    on reelplay.users (deleted_at);

