create table if not exists users (
    userId UUID primary key DEFAULT gen_random_uuid(),
    fname varchar(64) not null,
    lname varchar (64) not null,
    age int not null,
    email varchar(64) not null,
    passwordHash varchar (64) not null
    );