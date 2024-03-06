package apps

import (
	"embed"

	"github.com/labstack/echo/v5"
)

//go:embed all:central/dist
var centralDir embed.FS
var CentralDirFS = echo.MustSubFS(centralDir, "central/dist")

//go:embed all:linkpos/dist
var posDir embed.FS
var PosDirFS = echo.MustSubFS(posDir, "linkpos/dist")
