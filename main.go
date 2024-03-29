package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/jsvm"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"

	apps "webbegg.com/test/pbadmin/apps"
	"webbegg.com/test/pbadmin/utils"
)

func main() {
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}

	exeDir := filepath.Dir(exePath)
	dataDir := filepath.Join(exeDir, "./data")
	migrationsDir := filepath.Join(exeDir, "./migrations")
	hooksDir := filepath.Join(exeDir, "./plugins")
	hooksWatch := true
	hooksPool := 25
	automigrate := true

	config := pocketbase.Config{
		DefaultDataDir: dataDir,
	}
	app := pocketbase.NewWithConfig(config)

	jsvm.MustRegister(app, jsvm.Config{
		MigrationsDir: migrationsDir,
		HooksDir:      hooksDir,
		HooksWatch:    hooksWatch,
		HooksPoolSize: hooksPool,
	})

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		TemplateLang: migratecmd.TemplateLangJS,
		Automigrate:  automigrate,
		Dir:          migrationsDir,
	})

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		utils.RegisterApp("/central/", e.Router, apps.CentralDirFS)
		utils.RegisterApp("/pos/", e.Router, apps.PosDirFS)

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
