package config

import (
	"log"
	"text/template"
)

// AppConfig holts the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
