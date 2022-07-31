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
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR REFERENCES users(username),
    title VARCHAR NOT NULL,
    descr VARCHAR NOT NULL, --description
    photoURL VARCHAR NOT NULL
);

CREATE TABLE items
(
    id BIGSERIAL PRIMARY KEY,
    externalLink VARCHAR NOT NULL,
    socialPostLink VARCHAR NOT NULL,
    photoURL VARCHAR NOT NULL,
    username VARCHAR NOT NULL,
    displayID INTEGER REFERENCES displays(id)
);
