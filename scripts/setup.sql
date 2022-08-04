CREATE TABLE users
(
    id BIGSERIAL PRIMARY KEY,
    
    username VARCHAR UNIQUE NOT NULL,
    email VARCHAR UNIQUE NOT NULL,
    first_name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    photo_url VARCHAR NOT NULL,
    social_links VARCHAR[] NOT NULL
);

CREATE TABLE displays
(
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT REFERENCES users(id),

    title VARCHAR NOT NULL,
    descr VARCHAR NOT NULL, --description
    photo_url VARCHAR NOT NULL
);

CREATE TABLE items
(
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT REFERENCES users(id),
    display_id INTEGER REFERENCES displays(id),

    external_link VARCHAR NOT NULL,
    social_post_link VARCHAR NOT NULL,
    photo_url VARCHAR NOT NULL
);
