ALTER TABLE Language
ADD COLUMN allow_use BOOLEAN NOT NULL DEFAULT false,
ADD COLUMN extension VARCHAR(64);