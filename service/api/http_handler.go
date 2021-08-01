package api

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"url_shortener/service/domain"
)

const (
	queryArgOriginURL = "origin_url"
	queryArgShortURL  = "short_url"
)

// NewHandler - creates an HTTP handler which serves docs UI and short URLs API
func NewHandler(cfg Config, shortURLs domain.ShortURLs) httpHandler {
	swaggerFS := http.FileServer(http.Dir(cfg.DocsUIDir))
	swaggerFS = http.StripPrefix(cfg.DocsPath, swaggerFS)

	return httpHandler{
		Cfg:       cfg,
		ShortURLs: shortURLs,
		SwaggerFS: swaggerFS,
	}
}

type httpHandler struct {
	Cfg       Config
	ShortURLs domain.ShortURLs
	SwaggerFS http.Handler
}

func (h httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case strings.HasPrefix(r.URL.Path, h.Cfg.DocsPath):
		h.SwaggerFS.ServeHTTP(w, r)
	case r.URL.Path == "/api/short_url":
		switch r.Method {
		case http.MethodGet:
			h.getShortURL(w, r)
		case http.MethodPost:
			h.createShortURL(w, r)
		case http.MethodDelete:
			h.deleteShortURL(w, r)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
	default:
		h.redirectToOriginURL(w, r)
	}
}

func (h httpHandler) getShortURL(w http.ResponseWriter, r *http.Request) {
	originURL := r.URL.Query().Get(queryArgOriginURL)
	shortURL, err := h.ShortURLs.GetByOriginURL(r.Context(), originURL)
	if err != nil {
		w.WriteHeader(errToStatus(err))
		return
	}

	_, _ = w.Write([]byte(shortURL))
}

func (h httpHandler) createShortURL(w http.ResponseWriter, r *http.Request) {
	originURL := r.URL.Query().Get(queryArgOriginURL)
	shortURL, err := h.ShortURLs.Create(r.Context(), originURL)
	if err != nil {
		w.WriteHeader(errToStatus(err))
		return
	}
	_, _ = w.Write([]byte(shortURL))
}

func (h httpHandler) deleteShortURL(w http.ResponseWriter, r *http.Request) {
	originURL := r.URL.Query().Get(queryArgOriginURL)
	shortURL := r.URL.Query().Get(queryArgShortURL)

	if (originURL == "" && shortURL == "") ||
		(originURL != "" && shortURL != "") {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var err error
	if shortURL != "" {
		err = h.ShortURLs.Delete(r.Context(), shortURL)
	} else {
		err = h.ShortURLs.DeleteByOriginURL(r.Context(), originURL)
	}
	if err != nil {
		w.WriteHeader(errToStatus(err))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h httpHandler) redirectToOriginURL(w http.ResponseWriter, r *http.Request) {
	shortURLPath := r.URL.Path
	originURL, err := h.ShortURLs.GetOriginURLByShortPath(r.Context(), shortURLPath)
	if err != nil {
		w.WriteHeader(errToStatus(err))
		return
	}

	w.Header().Set("Location", originURL)
	w.WriteHeader(http.StatusMovedPermanently)
}

func errToStatus(err error) int {
	switch {
	case errors.Is(err, domain.ErrNotFound):
		return http.StatusNotFound
	case errors.Is(err, domain.ErrInvalidParam):
		return http.StatusBadRequest
	case errors.Is(err, context.Canceled):
		return 499
	default:
		return http.StatusInternalServerError
	}
}
