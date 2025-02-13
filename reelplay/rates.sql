create table reelplay.rates
(
    user_id    bigint      not null,
    movie_id   bigint      not null,
    rate       bigint      not null,
    created_at datetime(3) null,
    updated_at datetime(3) null,
    primary key (user_id, movie_id),
    constraint fk_rates_movie
        foreign key (movie_id) references reelplay.movies (id)
            on update cascade on delete cascade,
    constraint fk_rates_user
        foreign key (user_id) references reelplay.users (id)
            on update cascade on delete cascade
);

