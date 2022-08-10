CREATE TABLE displays
(
    id BIGSERIAL PRIMARY KEY,
    user_id VARCHAR NOT NULL,

    title VARCHAR NOT NULL,
    descr VARCHAR NOT NULL, --description
    photo_url VARCHAR NOT NULL
);

CREATE TABLE items
(
    id BIGSERIAL PRIMARY KEY,
    user_id VARCHAR NOT NULL,
    display_id BIGINT REFERENCES displays(id) NOT NULL,

    external_link VARCHAR NOT NULL,
    social_post_link VARCHAR NOT NULL,
    photo_url VARCHAR NOT NULL
);
