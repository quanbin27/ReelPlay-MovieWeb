create table reelplay.episodes
(
    id             bigint auto_increment
        primary key,
    episode_number bigint      not null,
    movie_id       bigint      not null,
    source         longtext    null,
    duration       bigint      null,
    created_at     datetime(3) null,
    updated_at     datetime(3) null,
    deleted_at     datetime(3) null,
    constraint idx_episode_movie
        unique (episode_number, movie_id),
    constraint fk_episodes_movie
        foreign key (movie_id) references reelplay.movies (id)
);

create index idx_episodes_deleted_at
    on reelplay.episodes (deleted_at);

