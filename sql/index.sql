-- Index user
CREATE INDEX users_id ON users (user_id);
CREATE INDEX users_email ON users (email);

-- Index film
CREATE INDEX films_id ON films (film_id);
CREATE INDEX films_name ON films (name);
CREATE INDEX films_year ON films (year);
CREATE INDEX films_rating ON films (rating);

-- Index bookmarks
CREATE INDEX bookmarks_id ON bookmarks (b_id);
CREATE INDEX bookmarks_userId ON bookmarks (user_id);
CREATE INDEX bookmarks_filmId ON bookmarks (film_id);