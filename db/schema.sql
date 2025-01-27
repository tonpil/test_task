-- PostgreSQL schema
CREATE TABLE
    news (
        id BIGSERIAL PRIMARY KEY,
        title TEXT NOT NULL,
        content TEXT NOT NULL
    );

CREATE TABLE
    news_categories (
        id SERIAL PRIMARY KEY,
        news_id BIGINT NOT NULL,
        category_id BIGINT NOT NULL,
        FOREIGN KEY (news_id) REFERENCES news (id),
        FOREIGN KEY (category_id) REFERENCES categories (id)
    );

INSERT INTO
    news (title, content)
VALUES
    (
        'Breaking News',
        'This is the content of breaking news.'
    ),
    (
        'Technology Update',
        'Latest updates in technology.'
    ),
    (
        'Sports Highlights',
        'Highlights from today''s sports events.'
    );

INSERT INTO
    news_categories (news_id, category_id)
VALUES
    (1, 101),
    (2, 102),
    (3, 103),
    (3, 104);