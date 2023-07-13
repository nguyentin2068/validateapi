package plugin

import (
	"github.com/corazawaf/coraza/v3"
	"github.com/corazawaf/coraza/v3/experimental/plugins"
	"github.com/corazawaf/coraza/v3/experimental/plugins/plugintypes"
)

func init() {
	plugins.RegisterOperator("validateOpenAPI", func() plugintypes.Operator { return new(validateOpenAPI) })
}
