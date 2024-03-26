package utils

import (
	"io/fs"
	"net/http"
	"strings"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/pocketbase/pocketbase/apis"
)

func RegisterApp(trailedPath string, router *echo.Echo, appDir fs.FS) {
	router.GET(strings.TrimRight(trailedPath, "/"), func(c echo.Context) error {
		return c.Redirect(http.StatusTemporaryRedirect, strings.TrimLeft(trailedPath, "/"))
	})

	router.GET(
		trailedPath+"*",
		apis.StaticDirectoryHandler(appDir, true),
		middleware.Gzip(),
	)
}
