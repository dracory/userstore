package shared

import (
	"log/slog"
	"net/http"

	"github.com/dracory/userstore"
	"github.com/gouniverse/hb"
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
	AdminHomeURL   string
	WebsiteUrl     string

	TokenizedColumns []string
	// TokenCreate      func(columnName, columnValue string) (token string, err error)
	// TokenDelete      func(token string) (err error)
	// TokenRead        func(token string) (columnName, columnValue string, err error)
	// TokenUpdate      func(token string, columnName, columnValue string) (err error)
	// TokensCreate     func(columnValueMap map[string]string) (columnTokenMap map[string]string, err error)
	// TokensUpdate     func(tokenValueMap map[string]string) (err error)
	TokensBulk func(tokensToCreate map[string]string, tokensToUpdate map[string]string, tokensToDelete []string) (createdTokens map[string]string, err error)
	TokensRead func(columnTokenMap map[string]string) (columnValueMap map[string]string, err error)
}

type PageInterface interface {
	ToTag(config Config) hb.TagInterface
}
