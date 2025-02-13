create table reelplay.movie_actor
(
    movie_id bigint not null,
    actor_id bigint not null,
    primary key (movie_id, actor_id),
    constraint fk_movie_actor_actor
        foreign key (actor_id) references reelplay.actors (id),
    constraint fk_movie_actor_movie
        foreign key (movie_id) references reelplay.movies (id)
);

