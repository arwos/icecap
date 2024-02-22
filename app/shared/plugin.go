package shared

import (
	"go.arwos.org/icecap/app/shared/webstatic"
	"go.osspkg.com/goppy/plugins"
)

var Plugin = plugins.Inject(
	webstatic.Plugin,
)
