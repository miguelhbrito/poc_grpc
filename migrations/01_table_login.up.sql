CREATE TABLE IF NOT EXISTS login
(
    username CHARACTER varying(200) NOT NULL,
    password CHARACTER varying(200) NOT NULL,
    PRIMARY KEY (username)
);