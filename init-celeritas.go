package main

import (
	"fmt"
	"os"

	"github.com/king-d-dev/celeritas"
)

func initApplication() *application {
	workingDir, err := os.Getwd()
	quitAppIfError(err)

	cel := celeritas.Celeritas{AppName: "cele-demo", Debug: true, Version: "1.0.0"}
	err = cel.New(workingDir)
	quitAppIfError(err)

	fmt.Println("Debug is", cel.Debug)

	app := application{App: &cel}
	return &app
}
