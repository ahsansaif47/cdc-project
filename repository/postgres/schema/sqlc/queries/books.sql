-- file: books.sql

-- Insert a new book
-- name: CreateBook :one
INSERT INTO books (
    id,
    title,
    author,
    description,
    published_date,
    user_id,
    created_at,
    updated_at,
    niche
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING *;

-- Get a book by ID
-- name: GetBookByID :one
SELECT *
FROM books
WHERE id = $1;

-- List all books (optional: filtered by user_id or niche)
-- name: ListBooks :many
SELECT *
FROM books
ORDER BY created_at DESC;

-- List books by niche
-- name: ListBooksByNiche :many
SELECT *
FROM books
WHERE niche = $1
ORDER BY created_at DESC;

-- Update a book
-- name: UpdateBook :one
UPDATE books
SET
    title = $2,
    author = $3,
    description = $4,
    published_date = $5,
    user_id = $6,
    updated_at = $7,
    niche = $8
WHERE id = $1
RETURNING *;

-- Delete a book
-- name: DeleteBook :exec
DELETE FROM books
WHERE id = $1;
