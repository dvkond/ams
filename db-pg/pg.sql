CREATE TABLE ams.album (
	id serial4 NOT NULL,
	title varchar(128) NULL,
	artist varchar(255) NULL,
	price numeric(5, 2) NULL,
	CONSTRAINT album_pk PRIMARY KEY (id)
);

INSERT INTO album
  (title, artist, price)
VALUES
  ('Blue Train', 'John Coltrane', 56.99),
  ('Giant Steps', 'John Coltrane', 63.99),
  ('Jeru', 'Gerry Mulligan', 17.99),
  ('Sarah Vaughan', 'Sarah Vaughan', 34.98);