package config

import "html/template"

// holds the applicaiton config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
}
