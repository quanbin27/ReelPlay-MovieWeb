create table reelplay.user_watcheds
(
    id            bigint auto_increment
        primary key,
    user_id       bigint           not null,
    episode_id    bigint           not null,
    last_position bigint default 0 not null,
    updated_at    datetime(3)      null,
    constraint fk_user_watcheds_episode
        foreign key (episode_id) references reelplay.episodes (id),
    constraint fk_user_watcheds_user
        foreign key (user_id) references reelplay.users (id)
);

