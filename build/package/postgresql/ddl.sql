DROP TABLE IF EXISTS archive_videos;
CREATE TABLE archive_videos (
    id INT NOT NULL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    url TEXT NOT NULL,
    stremer VARCHAR(50) NOT NULL,
    thumbnail_image TEXT,
    started_datetime TIMESTAMP NOT NULL,
    ended_datetime TIMESTAMP
);
