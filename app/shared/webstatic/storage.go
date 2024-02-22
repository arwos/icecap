package webstatic

import (
	"net/http"

	"go.osspkg.com/goppy/errors"
	"go.osspkg.com/goppy/web"
	"go.osspkg.com/static"
)

const (
	hContentType  = "Content-Type"
	hCacheControl = "Cache-Control"
)

type Storage struct {
	config Config
	web    *static.Cache
	admin  *static.Cache
}

func New(c *ConfigWebStatic) *Storage {
	return &Storage{
		config: c.Config,
		web:    static.New(),
		admin:  static.New(),
	}
}

func (v *Storage) Up() error {
	return errors.Wrap(
		v.web.FromDir(v.config.Web),
		v.admin.FromDir(v.config.Admin),
	)
}

func (v *Storage) Down() error {
	return nil
}

func (v *Storage) InjectWebUI(r web.Router) {
	r.Get("/", handler(v.web))
	for _, file := range v.web.List() {
		r.Get(file, handler(v.web))
	}
}

func (v *Storage) InjectAdminUI(r web.Router) {
	r.Get("/", handler(v.admin))
	for _, file := range v.admin.List() {
		r.Get(file, handler(v.admin))
	}
}

func handler(cache *static.Cache) func(ctx web.Context) {
	return func(ctx web.Context) {
		file, cached := ctx.URL().Path, true
		if len(file) == 0 {
			file, cached = indexFile, false
		}
		b, ct := cache.Get(file)
		if b == nil {
			ctx.String(http.StatusNotFound, "file not found: %s", file)
			return
		}
		if cached {
			ctx.Header().Set(hCacheControl, "max-age=86400, public")
		}
		ctx.Header().Set(hContentType, ct)
		ctx.Bytes(http.StatusOK, b)
	}
}
