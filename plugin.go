package plugin

import (
	"github.com/corazawaf/coraza/v3"
	"github.com/corazawaf/coraza/v3/experimental/plugins"
)

func init() {
	plugins.RegisterOperator("validateOpenAPI", func() coraza.RuleOperator { return new(validateOpenAPI) })
}
