// Copyright 2023 Juan Pablo Tosso and the OWASP Coraza contributors
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"net/http"
	"os"
	"strings"

	corazawaf "command-line-arguments/home/tinnt2/FSOFT/github/CorazaWAF/main.go"

	"github.com/corazawaf/coraza/v3"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers/gorillamux"
)

type validateOpenAPI struct{}

func (o *validateOpenAPI) Init(data string) error{
	return nil
}

func (o *validateOpenAPI) Evaluate(tx *coraza.Transaction, value string) bool {
	reqe := strings.Split(value, " ")
	uri := reqe[1]
	req, _ := http.NewRequest(http.MethodGet, uri, nil)
	doc, ok := tx.Waf.Config.Get("apifile", nil)
	if !ok || doc == nil {
		return true
	}

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

var _ coraza.RuleOperator = &validateOpenAPI

