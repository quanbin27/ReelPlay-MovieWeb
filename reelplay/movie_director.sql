create table reelplay.movie_director
(
    movie_id    bigint not null,
    director_id bigint not null,
    primary key (movie_id, director_id),
    constraint fk_movie_director_director
        foreign key (director_id) references reelplay.directors (id),
    constraint fk_movie_director_movie
        foreign key (movie_id) references reelplay.movies (id)
);

