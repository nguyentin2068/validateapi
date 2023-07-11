package plugin

import (
	"github.com/corazawaf/coraza/v3"
	"github.com/corazawaf/coraza/v3/internal/operators"
	"github.com/corazawaf/coraza/v3/internal/seclang"
)

func init() {
	operators.RegisterPlugin("validateOpenAPI", func() coraza.RuleOperator { return new(validateOpenAPI) })
	seclang.RegisterDirectivePlugin("apischemafile", directiveAPISchema)
}
