CREATE TABLE users
(
    id              bigint PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    first_name      VARCHAR          NOT NULL,
    last_name       VARCHAR          NOT NULL,
    birth_date      TIMESTAMP        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    email           VARCHAR UNIQUE   NOT NULL,
    hashed_password VARCHAR          NOT NULL
);

CREATE TABLE events
(
    id       bigint PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    event_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name      VARCHAR          NOT NULL,
    description TEXT
);

CREATE TABLE user_events
(
    user_id  bigint REFERENCES users (id) ON DELETE CASCADE,
    event_id  bigint REFERENCES events (id) ON DELETE CASCADE
);

