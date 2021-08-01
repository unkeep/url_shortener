package database

import "github.com/jackc/pgx/v4"

// ErrNotFound - returns when requested entity is not found
var ErrNotFound = pgx.ErrNoRows
