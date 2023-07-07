// Copyright 2023 Juan Pablo Tosso and the OWASP Coraza contributors
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"net/http"
	"strings"
	"os"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers/gorillamux"
	"github.com/corazawaf/coraza/v3"
	"github.com/corazawaf/coraza/v3/experimental/plugins/plugintypes"
)

type validateOpenAPI struct{}

var _ plugintypes.Operator = &validateOpenAPI

func (o *validateOpenAPI) Evaluate(_ plugintypes.TransactionState, value string) bool {
	schemaFile := "/home/tinnt2/FSOFT/Anew/coraza/APISchema/api.json"
	if s := os.Getenv("SCHEMA_FILE"); s != "" {
		schemaFile = s
	}
	reqe := strings.Split(value, " ")
	uri := reqe[1]
	req, _ := http.NewRequest(http.MethodGet, uri, nil)
	// Load the OpenAPI document
	loader := openapi3.NewLoader()
	doc, _ := loader.LoadFromFile(schemaFile)

	// Find the operation (HTTP method + path) that matches the request
	router, _ := gorillamux.NewRouter(doc)
	route, pathParams, _ := router.FindRoute(req)

	// Create a RequestValidationInput
	requestValidationInput := &openapi3filter.RequestValidationInput{
		Request:    req,
		PathParams: pathParams,
		Route:      route,
	}
	httpreq := req.Context()
	// Validate the request
	if er := openapi3filter.ValidateRequest(httpreq, requestValidationInput); er != nil {
		return true
	}
	return false
}


