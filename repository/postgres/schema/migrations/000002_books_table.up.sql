CREATE TABLE books (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    author TEXT,
    niche TEXT,
    description TEXT,
    published_date DATE,
    user_id UUID REFERENCES users(id),
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE
);

-- optional index for faster filtering by niche
CREATE INDEX idx_books_niche ON books(niche);