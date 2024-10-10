CREATE TABLE IF NOT EXISTS songs(
    group varchar(255) not null,
    name varchar(255) not null,
    releaseDate varchar(255) not null,
    text text not null, 
    link varchar(255) not null
);