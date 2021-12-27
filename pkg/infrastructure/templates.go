package infrastructure

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/yaien/clothes-store-api/assets"
	"html/template"
	"strings"
)

type Templates struct {
	Sale      *template.Template
	Transport *template.Template
}

var funcs = func(config *Config) template.FuncMap {
	return template.FuncMap{
		"first": func(name string) string {
			return strings.Split(name, " ")[0]
		},
		"link": func(ref string) string {
			return strings.ReplaceAll(config.SMTP.RefLink, "{ref}", ref)
		},
		"currency": func(value int) string {
			return fmt.Sprintf("$%s", humanize.Comma(int64(value)))
		},
	}
}

func parseTemplates(config *Config) (*Templates, error) {
	fn := funcs(config)
	fs := assets.FS()
	sale, err := template.New("sale.html").Funcs(fn).ParseFS(fs, "templates/sale.html")
	if err != nil {
		return nil, fmt.Errorf("failed to parse sale template: %w", err)
	}
	transport, err := template.New("transport.html").Funcs(fn).ParseFS(fs, "templates/transport.html")
	if err != nil {
		return nil, fmt.Errorf("failed to parse transport template: %w", err)
	}
	return &Templates{
		Sale:      sale,
		Transport: transport,
	}, nil
}
