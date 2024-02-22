package webstatic

import "go.osspkg.com/goppy/plugins"

var Plugin = plugins.Plugin{
	Config: &ConfigWebStatic{},
	Inject: New,
}
