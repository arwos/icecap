package pubapi

import (
	"context"

	"go.osspkg.com/goppy/web"

	"go.osspkg.com/goppy/plugins"
)

var Plugin = plugins.Inject(
	_params{},
	New,
)

type (
	_params struct {
		RouterPool web.RouterPool
	}

	PubApi struct {
		router web.Router
	}
)

func New(p _params) *PubApi {
	return &PubApi{
		router: p.RouterPool.Main(),
	}
}

func (v *PubApi) Up(ctx context.Context) error {
	return nil
}

func (v *PubApi) Down() error {
	return nil
}
