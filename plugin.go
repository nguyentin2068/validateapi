// Copyright 2023 Juan Pablo Tosso and the OWASP Coraza contributors
// SPDX-License-Identifier: Apache-2.0

package plugin

import (

	"github.com/corazawaf/coraza/v3"
	"github.com/corazawaf/coraza/v3/experimental/plugins"
)

func init() {
	operators.RegisterPlugin("validateOpenAPI", func() coraza.RuleOperator { return new(validateOpenAPI) })
}
