package config

import "html/template"

//TPL represents...
var TPL *template.Template

func init() {
	TPL = template.Must(template.ParseGlob("templates/*.tmpl"))
}
