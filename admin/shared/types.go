package shared

import (
	"log/slog"
	"net/http"

	"github.com/gouniverse/hb"
	"github.com/gouniverse/userstore"
)

type LayoutOptions struct {
	Title      string
	Body       string
	Styles     []string
	StyleURLs  []string
	Scripts    []string
	ScriptURLs []string
}

type Config struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	Logger         *slog.Logger
	Store          userstore.StoreInterface
	Layout         func(http.ResponseWriter, *http.Request, LayoutOptions) string
	HomeURL        string
	WebsiteUrl     string

	Tokenized  []string
	Tokenize   func(clear []string) (tokens []string)
	Untokenize func(tokens []string) (clear []string)
}

type PageInterface interface {
	// hb.TagInterface
	ToTag(config Config) hb.TagInterface
}
