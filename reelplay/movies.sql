create table reelplay.movies
(
    id             bigint auto_increment
        primary key,
    name           longtext         null,
    year           bigint           null,
    num_episodes   bigint           null,
    description    longtext         null,
    language       longtext         null,
    country_id     bigint           null,
    time_for_ep    bigint           null,
    thumbnail      longtext         null,
    trailer        longtext         null,
    rate           float            null,
    predict_rate   float            null,
    is_recommended tinyint(1)       null,
    created_at     datetime(3)      null,
    updated_at     datetime(3)      null,
    deleted_at     datetime(3)      null,
    view           bigint default 0 null,
    constraint fk_movies_country
        foreign key (country_id) references reelplay.countries (id)
);

create index idx_movies_deleted_at
    on reelplay.movies (deleted_at);

