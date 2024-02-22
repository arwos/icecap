package webapi

import (
	"context"

	"go.arwos.org/icecap/app/shared/webstatic"

	"go.osspkg.com/goppy/plugins"
	"go.osspkg.com/goppy/web"
)

var Plugin = plugins.Inject(
	_params{},
	New,
)

type (
	_params struct {
		RouterPool web.RouterPool
		UI         *webstatic.Storage
	}

	WebApi struct {
		router web.Router
		ui     *webstatic.Storage
	}
)

func New(p _params) *WebApi {
	return &WebApi{
		router: p.RouterPool.Main(),
		ui:     p.UI,
	}
}

func (v *WebApi) Up(ctx context.Context) error {
	v.ui.InjectWebUI(v.router)
	return nil
}

func (v *WebApi) Down() error {
	return nil
}
