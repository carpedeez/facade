CREATE TABLE users
(
    username VARCHAR PRIMARY KEY,
    fname VARCHAR NOT NULL,
    lname VARCHAR NOT NULL,
    photourl VARCHAR NOT NULL,
    sociallinks VARCHAR[] NOT NULL
);

CREATE TABLE displays
(
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR REFERENCES users(username),
    title VARCHAR NOT NULL,
    descr VARCHAR NOT NULL, --description
    photourl VARCHAR NOT NULL
);

CREATE TABLE items
(
    id BIGSERIAL PRIMARY KEY,
    externallink VARCHAR NOT NULL,
    socialpostlink VARCHAR NOT NULL,
    photourl VARCHAR NOT NULL,
    username VARCHAR NOT NULL,
    displayid INTEGER REFERENCES displays(id)
);
