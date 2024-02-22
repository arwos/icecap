package webstatic

import (
	"fmt"

	"go.osspkg.com/goppy/iofile"
)

const indexFile = "/index.html"

type (
	ConfigWebStatic struct {
		Config Config `yaml:"ui_static"`
	}

	Config struct {
		Web   string `yaml:"web"`
		Admin string `yaml:"admin"`
	}
)

func (v *ConfigWebStatic) Default() {
	if len(v.Config.Web) == 0 {
		v.Config.Web = "./ui/web/dist"
	}
	if len(v.Config.Admin) == 0 {
		v.Config.Admin = "./ui/admin/dist"
	}
}

func (v *ConfigWebStatic) Validate() error {
	if len(v.Config.Web) == 0 {
		return fmt.Errorf("web ui folder not set in config")
	}
	if !iofile.Exist(v.Config.Web + indexFile) {
		return fmt.Errorf("index file for web ui not found: %s", v.Config.Web+indexFile)
	}
	if len(v.Config.Admin) == 0 {
		return fmt.Errorf("admin ui folder not set in config")
	}
	if !iofile.Exist(v.Config.Admin + indexFile) {
		return fmt.Errorf("index file for web ui not found: %s", v.Config.Admin+indexFile)
	}
	return nil
}
