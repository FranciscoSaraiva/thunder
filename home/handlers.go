//Package home of Thunder
package home

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"thunder/config"
)

//Index function that runs a template of the file "index.tmpl"
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	config.TPL.ExecuteTemplate(w, "index.tmpl", nil)
}
