create table reelplay.bookmarks
(
    user_id    bigint      not null,
    movie_id   bigint      not null,
    created_at datetime(3) null,
    primary key (user_id, movie_id),
    constraint fk_bookmarks_movie
        foreign key (movie_id) references reelplay.movies (id)
            on update cascade on delete cascade,
    constraint fk_bookmarks_user
        foreign key (user_id) references reelplay.users (id)
            on update cascade on delete cascade
);

