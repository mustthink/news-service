CREATE TABLE IF NOT EXISTS news
(
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    author_id BIGINT,
    topic_id BIGINT,
    title VARCHAR(255),
    content TEXT
);

CREATE INDEX title_news ON news(title);
CREATE INDEX idx_news_author_id ON news(author_id);