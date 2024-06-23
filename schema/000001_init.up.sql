CREATE TABLE authors
(
    id serial not null unique,
    name varchar(255) not null,
    surname varchar(255),
    biography varchar(8000),
    born_date date
);

CREATE TABLE books
(
    id serial not null unique,
    name varchar(255) not null,
    author_id int references authors(id) on DELETE cascade not null,
    publishing_year int,
    isbn varchar(255) not null
);