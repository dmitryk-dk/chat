CREATE TABLE IF NOT EXIST user (
    userId uuid NOT NUL PRIMARY KEY,
    nickname VARCHAR(50) NOT NULL,
    regDate DATE NOT NULL,
);