package main

import (
	"context"
	"log"
	"os"

	"github.com/cristoper/gsheet/gdrive"
	"github.com/cristoper/gsheet/gsheet"
)

// global service objects used by commands
var (
	driveSvc = func() *gdrive.Service {
		svc, err := gdrive.NewServiceWithCtx(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		return svc
	}()

	sheetSvc = func() *gsheet.Service {
		svc, err := gsheet.NewServiceWithCtx(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		return svc
	}()
)

func main() {
	app.EnableBashCompletion = true
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
