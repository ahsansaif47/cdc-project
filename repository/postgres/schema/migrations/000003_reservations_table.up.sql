-- up.sql for creating reservations table
CREATE TABLE reservations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id),
    book_id UUID NOT NULL REFERENCES books(id),
    reserved_from TIMESTAMP WITH TIME ZONE NOT NULL,
    reserved_to TIMESTAMP WITH TIME ZONE NOT NULL,
    status VARCHAR(20) DEFAULT 'active',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Optional: prevent overlapping reservations for the same book
CREATE EXTENSION IF NOT EXISTS btree_gist;

ALTER TABLE reservations
ADD CONSTRAINT no_overlap_book
EXCLUDE USING gist (
    book_id WITH =,
    tsrange(reserved_from, reserved_to) WITH &&
);