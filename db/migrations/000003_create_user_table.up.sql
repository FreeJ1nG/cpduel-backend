CREATE TABLE IF NOT EXISTS Account (
  username VARCHAR(64) PRIMARY KEY,
  full_name VARCHAR(128) NOT NULL,
  password_hash VARCHAR(256) NOT NULL
);