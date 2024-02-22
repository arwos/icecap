/*
 *  Copyright (c) 2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a GPL-3.0 license that can be found in the LICENSE file.
 */

package main

import (
	icapp "go.arwos.org/icecap/app"
	"go.osspkg.com/goppy"
	"go.osspkg.com/goppy/metrics"
	"go.osspkg.com/goppy/ormmysql"
	"go.osspkg.com/goppy/web"
)

var Version = "v0.0.0-dev"

func main() {
	app := goppy.New()
	app.AppName("icecap")
	app.AppVersion(Version)
	app.Plugins(
		metrics.WithMetrics(),
		web.WithHTTP(),
		ormmysql.WithMySQL(),
	)
	app.Plugins(icapp.Plugins...)
	app.Run()
}
