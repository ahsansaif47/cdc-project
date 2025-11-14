CREATE TABLE roles (
    id SERIAL PRIMARY KEY,

    name TEXT NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Unique index for role name
CREATE UNIQUE INDEX idx_roles_name ON roles (name);

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id UUID PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255),
    auth_provider_type VARCHAR(20) NOT NULL DEFAULT 'local',
    role_id INTEGER NOT NULL,
    phone_number VARCHAR(20),
    verified BOOLEAN NOT NULL DEFAULT false,
    is_blocked BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,

    CONSTRAINT fk_users_role
        FOREIGN KEY (role_id)
            REFERENCES roles (id)
            ON UPDATE CASCADE
            ON DELETE RESTRICT
);

-- Unique Indexes
CREATE UNIQUE INDEX idx_users_username ON users (username);
CREATE UNIQUE INDEX idx_users_email ON users (email);

-- Soft delete index
CREATE INDEX idx_users_deleted_at ON users (deleted_at);