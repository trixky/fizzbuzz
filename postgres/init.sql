\c fizzbuzz;

CREATE TABLE fizzbuzz_request_statistics (
    int1 INTEGER NOT NULL,
    int2 INTEGER NOT NULL,
    _limit INTEGER NOT NULL,
    str1 VARCHAR(51) NOT NULL,
    str2 VARCHAR(51) NOT NULL,
    request_number INTEGER DEFAULT 1 NOT NULL
);

CREATE TABLE api_users (
    user_id serial PRIMARY KEY,
    pseudo VARCHAR (255) UNIQUE NOT NULL,
    password VARCHAR (65) NOT NULL,
    blocked BOOLEAN DEFAULT 'false' NOT NULL,
    admin BOOLEAN DEFAULT 'false' NOT NULL
);

INSERT INTO api_users (pseudo, password, admin) VALUES ('admin', 'c54b18a947c806a48d7fd825ec6aea73', 'true');