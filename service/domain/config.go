package domain

// Config - domain config
type Config struct {
	ShortURLsBase string `envconfig:"SHORT_URLS_BASE" required:"true"`
}
