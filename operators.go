// Copyright 2023 Juan Pablo Tosso and the OWASP Coraza contributors
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"net/http"
	"strings"

	"github.com/corazawaf/coraza/v3/experimental/plugins"
	"github.com/corazawaf/coraza/v3/experimental/plugins/plugintypes"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers/gorillamux"
)

type validateOpenAPI struct{}

var _ plugintypes.Operator = (*validateOpenAPI)(nil)

func newValidateOpenAPI(plugintypes.OperatorOptions) (plugintypes.Operator, error) {
	return &validateOpenAPI{}, nil
}
// func (o *validateOpenAPI) Init(data string) error{
// 	return nil
// }

func (o *validateOpenAPI) Evaluate(tx plugintypes.TransactionState, value string) bool {
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
	plugins.RegisterOperator("validateOpenAPI", newValidateOpenAPI)
}



