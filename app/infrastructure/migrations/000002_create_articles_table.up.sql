CREATE TABLE IF NOT EXISTS articles(
  id serial PRIMARY KEY,
  title VARCHAR (50) NOT NULL,
  content VARCHAR NOT NULL,
  author_id serial NOT NULL,
  INDEX (author_id),
  FOREIGN KEY (author_id)
  REFERENCES users(id)
  ON UPDATE CASCADE ON DELETE RESTRICT
)