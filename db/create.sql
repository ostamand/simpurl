CREATE DATABASE IF NOT EXISTS simpurl;

USE simpurl;

CREATE TABLE IF NOT EXISTS links (
    id INTEGER NOT NULL AUTO_INCREMENT,
    user_id INTEGER NOT NULL,
    symbol VARCHAR(255),
    url VARCHAR(4096),
    description VARCHAR(4096),
    note VARCHAR(4096),
    created_at DATETIME NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS users (
    id INTEGER NOT NULL AUTO_INCREMENT,
    username VARCHAR(255),
    hashed_password VARCHAR(255),
    admin BOOLEAN,
    created_at DATETIME NOT NULL,
    PRIMARY KEY(id),
    CONSTRAINT UC_users UNIQUE(username)
);

CREATE TABLE IF NOT EXISTS sessions (
    id INTEGER NOT NULL AUTO_INCREMENT,
    user_id INTEGER NOT NULL,
    token VARCHAR(255) NOT NULL,
    created_at DATETIME NOT NULL,
    expiry_at DATETIME NOT NULL,
    PRIMARY KEY(id)
);