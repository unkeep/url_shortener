package domain

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"url_shortener/service/database"
)

// ShortURLs - short URLs BL
type ShortURLs interface {
	// Create - creates a short URl from the original one. Subsequent calls returns the same URL
	Create(ctx context.Context, originURL string) (shortURL string, err error)
	// Delete - deletes a short URL. Does not fail for non existing URL
	Delete(ctx context.Context, shortURL string) (err error)
	// DeleteByOriginURL - deletes a short URL by an original one. Does not fail for non existing URL
	DeleteByOriginURL(ctx context.Context, originURL string) (err error)
	// GetOriginURLByShortPath - gets an original URL by a short URL path. Returns ErrNotFound of URL is not found
	GetOriginURLByShortPath(ctx context.Context, shortURLPath string) (originURL string, err error)
	// GetByOriginURL - gets a short URL by an original one. Returns ErrNotFound of URL is not found
	GetByOriginURL(ctx context.Context, originURL string) (shortURL string, err error)
}

// NewShortURLs - creates ShortURLs
func NewShortURLs(urlStorage database.URLStorage, baseURL string) shortURLs {
	return shortURLs{
		URLStorage:   urlStorage,
		BaseURL:      baseURL,
		URLIDEncoder: base62URLIDEncoder{},
	}
}

type shortURLs struct {
	URLStorage   database.URLStorage
	BaseURL      string
	URLIDEncoder urlIDEncoder
}

func (u shortURLs) Create(ctx context.Context, originURL string) (string, error) {
	_, err := u.validateURL(originURL)
	if err != nil {
		return "", fmt.Errorf("validateURL: %s: %w", err.Error(), ErrInvalidParam)
	}

	id, err := u.URLStorage.Create(ctx, originURL)
	if err != nil {
		return "", fmt.Errorf("URLStorage.Create: %w", err)
	}

	shortURL, err := u.buildShortURLFromID(id)
	if err != nil {
		return "", fmt.Errorf("buildShortURLFromID: %w", err)
	}

	return shortURL, nil
}

func (u shortURLs) Delete(ctx context.Context, shortURL string) error {
	id, err := u.parseIDFromShortURL(shortURL)
	if err != nil {
		return fmt.Errorf("parseIDFromShortURL: %s: %w", err.Error(), ErrInvalidParam)
	}

	if err := u.URLStorage.DeleteByID(ctx, id); err != nil {
		return fmt.Errorf("URLStorage.Delete: %w", err)
	}

	return nil
}

func (u shortURLs) DeleteByOriginURL(ctx context.Context, originURL string) error {
	_, err := u.validateURL(originURL)
	if err != nil {
		return fmt.Errorf("validateURL: %s: %w", err.Error(), ErrInvalidParam)
	}

	if err := u.URLStorage.Delete(ctx, originURL); err != nil {
		return fmt.Errorf("URLStorage.Delete: %w", err)
	}

	return nil
}

func (u shortURLs) GetOriginURLByShortPath(ctx context.Context, shortURLPath string) (string, error) {
	encodedID := strings.TrimLeft(shortURLPath, "/")
	if encodedID == "" {
		return "", fmt.Errorf("path is empty: %w", ErrInvalidParam)
	}

	id, err := u.URLIDEncoder.Decode(encodedID)
	if err != nil {
		return "", fmt.Errorf("URLIDEncoder.Decode: %s: %w", err.Error(), ErrInvalidParam)
	}

	originURL, err := u.URLStorage.Get(ctx, id)
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			err = fmt.Errorf("%w", ErrNotFound)
		}
		return "", fmt.Errorf("URLStorage.Get: %w", err)
	}

	return originURL, nil
}

func (u shortURLs) GetByOriginURL(ctx context.Context, originURL string) (string, error) {
	_, err := u.validateURL(originURL)
	if err != nil {
		return "", fmt.Errorf("validateURL: %s: %w", err.Error(), ErrInvalidParam)
	}

	id, err := u.URLStorage.Find(ctx, originURL)
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			err = fmt.Errorf("%w", ErrNotFound)
		}
		return "", fmt.Errorf("URLStorage.Find: %w", err)
	}

	shortURL, err := u.buildShortURLFromID(id)
	if err != nil {
		return "", fmt.Errorf("buildShortURLFromID: %w", err)
	}

	return shortURL, nil
}

func (u shortURLs) buildShortURLFromID(id uint64) (string, error) {
	encodedID, err := u.URLIDEncoder.Encode(id)
	if err != nil {
		return "", fmt.Errorf("URLIDEncoder.Encode: %w", err)
	}
	return fmt.Sprintf("%s/%s", u.BaseURL, encodedID), nil
}

func (u shortURLs) parseIDFromShortURL(shortURL string) (uint64, error) {
	parsedURL, err := u.validateURL(shortURL)
	if err != nil {
		return 0, fmt.Errorf("validateURL: %w", err)
	}

	encodedID := strings.TrimLeft(parsedURL.Path, "/")

	id, err := u.URLIDEncoder.Decode(encodedID)
	if err != nil {
		return 0, fmt.Errorf("URLIDEncoder.Decode: %w", err)
	}

	return id, nil
}

func (u shortURLs) validateURL(URL string) (*url.URL, error) {
	if URL == "" {
		return nil, fmt.Errorf("URL is empty")
	}

	parsed, err := url.Parse(URL)
	if err != nil {
		return nil, fmt.Errorf("url.Parse: %w", err)
	}

	return parsed, nil
}
