package plugin

import (
	"github.com/corazawaf/coraza/v3"
	"github.com/getkin/kin-openapi/openapi3"
)

func directiveAPISchema(w *coraza.Waf, path string) error {
	loader := openapi3.NewLoader()
	doc, _ := loader.LoadFromFile(path)
	if err != nil {
		return err
	}
	w.Config.Set("apifile", doc)
	return nil
}