package main

import (
	"github.com/Kripipastt/go-expression-parser/internal/application"
)

func main() {
	app := application.New()
	//fmt.Println(app.Execute("2 * 2 / 4 + 3 - 26 + 3 - (5 - 50 * 3 / 10)"))
	app.RunServer()
}
