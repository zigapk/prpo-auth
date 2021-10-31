CREATE TABLE users
(
    uid          TEXT PRIMARY KEY,

    email        TEXT                                               NOT NULL UNIQUE,
    name         TEXT                                               NOT NULL,

    password     TEXT                                               NOT NULL,

    date_created TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE authenticated_devices
(
    user_id      TEXT REFERENCES users (uid) ON DELETE CASCADE      NOT NULL,
    token        TEXT                                               NOT NULL,

    date_created TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    last_used    TIMESTAMP WITH TIME ZONE                           NOT NULL
);