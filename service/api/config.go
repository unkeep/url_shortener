package api

// Config - API config
type Config struct {
	DocsPath  string `envconfig:"DOCS_PATH" required:"true"`
	DocsUIDir string `envconfig:"DOCS_UI_DIR" required:"true"`
}
