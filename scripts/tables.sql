--- create a user with a password ---
 CREATE user gopher WITH PASSWORD "password";

--- create a database --
CREATE DATABASE cmsdb;

--- grant priv ---
GRANT ALL PRIVILEGES ON DATABASE cmsdb TO gopher;

--- create table pages ---
CREATE TABLE IF NOT EXISTS PAGES(
    id      SERIAL PRIMARY KEY,
    title   TEXT NOT NULL,
    content TEXT NOT NULL
);

--- create table posts ---
CREATE TABLE IF NOT EXISTS POSTS(
    id              SERIAL PRIMARY KEY,
    title           TEXT NOT NULL,
    content         TEXT NOT NULL,
    created_date    DATE NOT NULL
);

--- create table comments
CREATE TABLE IF NOT EXISTS COMMENTS(
    id              SERIAL PRIMARY KEY,
    author          TEXT NOT NULL,
    content         TEXT NOT NULL,
    created_date    DATE NOT NULL,
    post_id         INT referances POSTS(id)
);