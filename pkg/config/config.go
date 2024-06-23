package config

import (
	"html/template"
	"log"
)

// AppConfig holds the application config
type AppConfig struct {
	useCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
