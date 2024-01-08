DROP TABLE IF EXISTS users;

CREATE TABLE
    users (
        id VARCHAR(32) PRIMARY KEY,
        password VARCHAR(255) NOT NULL,
        email VARCHAR(50) NOT NULL,
        created_at TIMESTAMP NOT NULL DEFAULT NOW (),
        updated_at TIMESTAMP NOT NULL DEFAULT NOW ()
    );

