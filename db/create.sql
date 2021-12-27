CREATE DATABASE IF NOT EXISTS shorturl;

USE shorturl;

CREATE TABLE IF NOT EXISTS links (
    id INTEGER NOT NULL AUTO_INCREMENT,
    symbol VARCHAR(255),
    url VARCHAR(4096),
    description VARCHAR(4096),
    PRIMARY KEY(id),
    CONSTRAINT UC_links UNIQUE (symbol)
);

CREATE TABLE IF NOT EXISTS users (
    id INTEGER NOT NULL AUTO_INCREMENT,
    username VARCHAR(255),
    password VARCHAR(255),
    PRIMARY KEY(id),
    CONSTRAINT UC_users UNIQUE(username)
);