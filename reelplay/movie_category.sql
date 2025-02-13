create table reelplay.movie_category
(
    movie_id    bigint not null,
    category_id bigint not null,
    primary key (movie_id, category_id),
    constraint fk_movie_category_category
        foreign key (category_id) references reelplay.categories (id),
    constraint fk_movie_category_movie
        foreign key (movie_id) references reelplay.movies (id)
);

