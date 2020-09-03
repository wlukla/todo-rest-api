CREATE TABLE todo (
    id bigserial not null primary key,
    title varchar(255) not null,
    description text not null
);