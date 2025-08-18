create table cities (
    id serial primary key,
    name text
);

insert into cities
    (name)
values
    ('Москва'),
    ('Санкт-Петербург'),
    ('Краснодар');


create table users (
    id serial primary key,
    name text not null ,
    city_id int not null references cities (id)
);


insert into users
(name, city_id)
values
    ('Иван', 1),
    ('Анна', 1),
    ('Олег', 1);
update users set city_id = 2 where name = 'Олег';

select * from cities;
select * from users;
