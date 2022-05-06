DROP TABLE IF EXISTS archive_videos;
CREATE TABLE archive_videos (
    id SERIAL PRIMARY KEY,
    broadcast_id VARCHAR(255) NOT NULL,
    title VARCHAR(255) NOT NULL,
    url TEXT,
    stremer VARCHAR(50) NOT NULL,
    platform SMALLINT NOT NULL,
    thumbnail_image TEXT,
    started_datetime TIMESTAMP NOT NULL,
    ended_datetime TIMESTAMP,
    UNIQUE(broadcast_id)
);
