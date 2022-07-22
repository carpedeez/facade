CREATE TABLE users
(
    username VARCHAR PRIMARY KEY,
    fname VARCHAR NOT NULL,
    lname VARCHAR NOT NULL,
    photoURL VARCHAR NOT NULL,
    socialLinks VARCHAR[] NOT NULL
);

CREATE TABLE displays
(
    id SERIAL PRIMARY KEY,
    username VARCHAR REFERENCES users(username),
    descr VARCHAR NOT NULL, --description
    photoURL VARCHAR NOT NULL
);

CREATE TABLE items
(
    id SERIAL PRIMARY KEY,
    externalLink VARCHAR NOT NULL,
    socialPostLink VARCHAR NOT NULL,
    photoURL VARCHAR NOT NULL,
    username VARCHAR NOT NULL,
    displayID INTEGER REFERENCES displays(id)
);