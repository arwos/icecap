package app

import (
	"go.arwos.org/icecap/app/pubapi"
	"go.arwos.org/icecap/app/shared"
	"go.arwos.org/icecap/app/webapi"
	"go.osspkg.com/goppy/plugins"
)

var Plugins = plugins.Inject(
	pubapi.Plugin,
	webapi.Plugin,
	shared.Plugin,
)
