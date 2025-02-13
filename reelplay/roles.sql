create table reelplay.roles
(
    id   bigint auto_increment
        primary key,
    name varchar(191) null,
    constraint uni_roles_name
        unique (name)
);

