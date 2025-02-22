package main

import (
	"github.com/Kripipastt/go-expression-parser/orchestrator/internal/application"
)

func main() {
	app := application.NewApplication()
	app.MountHandlers()
	//fmt.Println(app.Execute("2 - (5 ^ (4 + 4) * 0.25) - 2 + 2 / (10 / 3 ^ 4)"))
	app.RunServer()
}
