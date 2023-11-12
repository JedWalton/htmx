-- Create a new table for storing films
CREATE TABLE films (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    director VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Indexes can be added for fields that you expect to query frequently
CREATE INDEX idx_films_title ON films (title);
CREATE INDEX idx_films_director ON films (director);

