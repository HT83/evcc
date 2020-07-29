package test

import (
	"github.com/andig/evcc-config/registry"
	_ "github.com/andig/evcc-config/templates"
	"gopkg.in/yaml.v3"
)

type ParsedTempalte struct {
	registry.Template
	Config map[string]interface{}
}

// ConfigFromYAML parses configuration from yaml string
func ConfigFromYAML(sample string) (map[string]interface{}, error) {
	var conf map[string]interface{}
	err := yaml.Unmarshal([]byte(sample), &conf)
	return conf, err
}

// ConfigTemplates returns configuration templates for giving class
func ConfigTemplates(class string) (res []ParsedTempalte) {
	templates := registry.TemplatesByClass(class)

	for _, tmpl := range templates {
		conf, err := ConfigFromYAML(tmpl.Sample)
		if err != nil {
			continue
		}

		parsed := ParsedTempalte{
			Template: tmpl,
			Config:   conf,
		}

		res = append(res, parsed)
	}

	return res
}
