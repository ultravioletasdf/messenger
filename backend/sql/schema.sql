CREATE TABLE IF NOT EXISTS users (
    id integer PRIMARY KEY,
    username text NOT NULL,
    email text NOT NULL,
    display_name text,
    bio text,

    password text NOT NULL,
    created_at integer NOT NULL,
    updated_at integer NOT NULL
);

CREATE INDEX idx_users_id ON users (id);
CREATE UNIQUE INDEX idx_users_username ON users (username);
CREATE UNIQUE INDEX idx_users_email ON users (email);

CREATE TABLE IF NOT EXISTS sessions (
    token text PRIMARY KEY,
    user_id integer NOT NULL REFERENCES users (id),
    platform text NOT NULL,
    ip text NOT NULL,

    created_at integer NOT NULL
);
CREATE INDEX idx_sessions_token ON sessions (token);