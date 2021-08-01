package database

import (
	"context"

	"github.com/jackc/pgx/v4"
)

// URLStorage - is interface of URL storage
type URLStorage interface {
	// Create - saves the given URL returning it's ID. Subsequent calls returns the same ID
	Create(ctx context.Context, url string) (uint64, error)
	// Get - gets URL by it's ID. Returns ErrNotFound if URL is not found
	Get(ctx context.Context, id uint64) (string, error)
	// Find - finds URL ID. Returns ErrNotFound if URL is not found
	Find(ctx context.Context, url string) (uint64, error)
	// DeleteByID - deletes URL by it's ID. Doesnt fail if the URL with the ID is not found
	DeleteByID(ctx context.Context, id uint64) error
	// Delete - deletes URL. Doesnt fail if the URL is not found
	Delete(ctx context.Context, url string) error
}

// NewURLStorage - creates URLStorage
func NewURLStorage(conn *pgx.Conn) urlStorage {
	return urlStorage{Conn: conn}
}

// urlStorage - implements URLStorage
type urlStorage struct {
	Conn *pgx.Conn
}

func (u urlStorage) Create(ctx context.Context, url string) (uint64, error) {
	const query = `
insert into urls (url) values ($1)
on conflict (url) do update set url=excluded.url
returning id`

	var id uint64
	return id, u.Conn.QueryRow(ctx, query, url).Scan(&id)
}

func (u urlStorage) Get(ctx context.Context, id uint64) (string, error) {
	const query = `select url from urls where id=$1`
	var url string
	return url, u.Conn.QueryRow(ctx, query, id).Scan(&url)
}

func (u urlStorage) Find(ctx context.Context, url string) (uint64, error) {
	const query = `select id from urls where url=$1`
	var id uint64
	return id, u.Conn.QueryRow(ctx, query, url).Scan(&id)
}

func (u urlStorage) DeleteByID(ctx context.Context, id uint64) error {
	const query = `delete from urls where id=$1`
	_, err := u.Conn.Exec(ctx, query, id)
	return err
}

func (u urlStorage) Delete(ctx context.Context, url string) error {
	const query = `delete from urls where url=$1`
	_, err := u.Conn.Exec(ctx, query, url)
	return err
}
