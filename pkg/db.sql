create table if not exists snippets (
    id serial not null primary key,
    title varchar(100) not null,
    content text not null,
    created date not null,
    expires date not null
);