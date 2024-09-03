package config

import (
	"html/template"
)

// AppConfig holds the application configuration

type AppConfig struct {
	UseCache      bool
	TamplateCache map[string]*template.Template
}
