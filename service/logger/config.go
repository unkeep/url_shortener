package logger

// Config is logging config
type Config struct {
	Format string `envconfig:"LOG_FORMAT" required:"true"`
	Level  string `envconfig:"LOG_LEVEL" required:"true"`
}
