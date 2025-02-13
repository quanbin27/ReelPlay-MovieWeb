create table reelplay.category_fits
(
    user_id     bigint      not null,
    category_id bigint      not null,
    fit_rate    float       not null,
    created_at  datetime(3) null,
    updated_at  datetime(3) null,
    primary key (user_id, category_id),
    constraint fk_category_fits_category
        foreign key (category_id) references reelplay.categories (id)
            on update cascade on delete cascade,
    constraint fk_category_fits_user
        foreign key (user_id) references reelplay.users (id)
            on update cascade on delete cascade
);

