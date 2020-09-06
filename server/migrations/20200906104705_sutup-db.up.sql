CREATE TABLE todo (
    id bigserial not null primary key,
    title varchar(255) not null,
    description text not null
);

CREATE TABLE users (
    id bigserial not null primary key,
    username varchar(255) not null,
    password varchar(255) not null
)