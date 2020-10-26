package core

import (
	"fmt"
	"html/template"
)

type Templates struct {
	Sale      *template.Template
	Transport *template.Template
}

func parseTemplates() (*Templates, error) {
	sale, err := template.ParseFiles("assets/parseTemplates/sale.html")
	if err != nil {
		return nil, fmt.Errorf("failed to parse sale template: %w", err)
	}
	transport, err := template.ParseFiles("assets/parseTemplates/transport.html")
	if err != nil {
		return nil, fmt.Errorf("failed to parse transport template: %w", err)
	}
	return &Templates{
		Sale:      sale,
		Transport: transport,
	}, nil
}
