-- DROP
-- DROP TABLE users;
-- DROP TABLE films;
-- DROP TABLE bookmarks;


-- INIT

CREATE TABLE IF NOT EXISTS users (
	user_id VARCHAR(255) PRIMARY KEY,
	email VARCHAR(255) UNIQUE,
	password VARCHAR(255),
	role VARCHAR(10),
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL 
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS films (
	film_id VARCHAR(255) PRIMARY KEY,
	name VARCHAR(255) UNIQUE,
	year VARCHAR(255),
	rating DOUBLE,
	image TEXT,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL 
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS bookmarks (
	b_id VARCHAR(255) PRIMARY KEY,
	user_id VARCHAR(255),
	film_id VARCHAR(255)
) ENGINE=InnoDB;

ALTER TABLE bookmarks ADD CONSTRAINT FK_UserBM FOREIGN KEY (user_id) REFERENCES users(user_id);
ALTER TABLE bookmarks ADD CONSTRAINT FK_FilmBM FOREIGN KEY (film_id) REFERENCES films(film_id);


-- INDEX
CREATE INDEX idx_users ON users(user_id, email, role);
CREATE INDEX idx_films ON films(film_id, name, year, rating);
CREATE INDEX idx_bookmarks ON bookmarks(b_id, user_id, film_id);

