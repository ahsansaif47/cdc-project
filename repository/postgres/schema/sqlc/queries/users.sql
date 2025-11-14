-- =====================
-- Check if email exists
-- =====================
-- name: CheckExistingEmail :one
SELECT EXISTS (
    SELECT 1
    FROM users
    WHERE email = $1
);

-- =====================
-- Create a new user
-- =====================
-- name: CreateUser :exec
INSERT INTO users (
    id,
    username,
    email,
    password_hash,
    auth_provider_type,
    role_id,
    phone_number,
    verified,
    is_blocked,
    created_at,
    updated_at,
    deleted_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
);

-- =====================
-- Find all users
-- =====================
-- name: FindAll :many
SELECT *
FROM users
WHERE deleted_at IS NULL;

-- =====================
-- Find user by ID
-- =====================
-- name: FindByID :one
SELECT *
FROM users
WHERE id = $1
AND deleted_at IS NULL;

-- -- =====================
-- -- Get all vendors
-- -- =====================
-- -- Assuming RoleID for vendor = 2 (adjust as needed)
-- -- name: GetAllVendors :many
-- SELECT *
-- FROM users
-- WHERE role_id = 2
-- AND deleted_at IS NULL;

-- =====================
-- Get all non-vendor users
-- =====================
-- Assuming RoleID for normal users = 1 (adjust as needed)
-- name: GetAllUsers :many
SELECT *
FROM users
WHERE role_id = 1
AND deleted_at IS NULL;

-- =====================
-- Set new password
-- =====================
-- name: SetNewPassword :exec
UPDATE users
SET password_hash = $2,
    updated_at = $3
WHERE email = $1
AND deleted_at IS NULL;

-- =====================
-- Validate user credentials
-- =====================
-- name: ValidateUserCredentials :one
SELECT *
FROM users
WHERE email = $1
AND password_hash = $2
AND deleted_at IS NULL;

-- =====================
-- Find user by email
-- =====================
-- name: FindUserByEmail :one
SELECT *
FROM users
WHERE email = $1
AND deleted_at IS NULL;