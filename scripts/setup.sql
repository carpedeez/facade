CREATE TABLE displays
(
    id UUID PRIMARY KEY,
    user_id VARCHAR NOT NULL,

    title VARCHAR NOT NULL,
    descr VARCHAR NOT NULL, --description
    photo_url VARCHAR NOT NULL
);

CREATE TABLE items
(
    id UUID PRIMARY KEY,
    user_id VARCHAR NOT NULL,
    display_id UUID REFERENCES displays(id) NOT NULL,

    external_link VARCHAR NOT NULL,
    social_post_link VARCHAR NOT NULL,
    photo_url VARCHAR NOT NULL
);
