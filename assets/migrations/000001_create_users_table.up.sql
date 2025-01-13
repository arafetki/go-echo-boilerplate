CREATE TABLE
    IF NOT EXISTS users (
        id varchar(36) PRIMARY KEY,
        name varchar(255) NOT NULL,
        email varchar(255) NOT NULL UNIQUE,
        avatar_url text,
        is_email_verified boolean NOT NULL,
        created_at timestamptz NOT NULL DEFAULT now (),
        updated_at timestamptz
    );