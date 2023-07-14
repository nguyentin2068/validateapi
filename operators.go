// Copyright 2023 Juan Pablo Tosso and the OWASP Coraza contributors
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"net/http"
	"strings"

	coraza"github.com/corazawaf/coraza/v3"
	"github.com/corazawaf/coraza/v3/operators"
    "github.com/corazawaf/coraza/v3/types"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers/gorillamux"
)
// Operator interface is used to define rule @operators
type validateOpenAPI interface {
	// Init is used during compilation to setup and cache
	// the operator
	Init(string) error
	// Evaluate is used during the rule evaluation,
	// it returns true if the operator succeeded against
	// the input data for the transaction
	Evaluate(*coraza.Transaction, string) bool
   }
type openAPIValidator struct{}

var _ coraza.Operator = &openAPIValidator{}

func(*openAPIValidator) Init(_ string) error{
	return nil
}

func (*openAPIValidator) Evaluate(_ coraza.Transaction, value string) bool {
	reqe := strings.Split(value, " ")
	methd := reqe[0]
	uri := reqe[1]
	req, _ := http.NewRequest(methd, uri, nil)
	loader := openapi3.NewLoader()
	doc, _ := loader.LoadFromFile("./APISchema/api.json")

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
func init() {
	operators.RegisterPlugin("validateOpenAPI", func() types.Operator {
  	return &openAPIValidator{}
 })
}


