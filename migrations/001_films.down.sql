-- Drop the indexes first
DROP INDEX IF EXISTS idx_films_title;
DROP INDEX IF EXISTS idx_films_director;

-- Then drop the table
DROP TABLE IF EXISTS films;
