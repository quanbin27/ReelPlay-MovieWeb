create table reelplay.comments
(
    id         bigint auto_increment
        primary key,
    content    longtext    not null,
    user_id    bigint      not null,
    movie_id   bigint      not null,
    created_at datetime(3) null,
    deleted_at datetime(3) null,
    constraint fk_comments_movie
        foreign key (movie_id) references reelplay.movies (id),
    constraint fk_comments_user
        foreign key (user_id) references reelplay.users (id)
);

create index idx_comments_deleted_at
    on reelplay.comments (deleted_at);

