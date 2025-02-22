package application

import (
	"bufio"
	"fmt"
	"github.com/Kripipastt/go-expression-parser/orchestrator/internal/application/api"
	"github.com/Kripipastt/go-expression-parser/orchestrator/internal/application/internal"
	"github.com/Kripipastt/go-expression-parser/orchestrator/internal/application/service"
	"github.com/Kripipastt/go-expression-parser/orchestrator/internal/config"
	"github.com/Kripipastt/go-expression-parser/orchestrator/pkg/parser"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"os"
	"strings"
)

type Application struct {
	Router *chi.Mux
}

func NewApplication() *Application {
	return &Application{Router: chi.NewRouter()}
}

func (app *Application) Execute(expression string) (float64, error) {
	return parser.Calc(expression)
}

func (app *Application) Run() error {
	fmt.Println("If you want to exit, input 'exit'")
	for {
		fmt.Println("input expression")
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("failed to read expression from console")
		}
		text = strings.TrimSpace(text)
		if text == "exit" {
			fmt.Println("application was successfully closed")
			return nil
		}
		result, _, err := parser.Parse(text)
		if err != nil {
			fmt.Println(text, " failed to calc expression: ", err)
		} else {
			fmt.Println(text, " = ", result)
		}

	}
}

func (app *Application) MountHandlers() {
	app.Router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	app.Router.Use(middleware.Heartbeat("/ping"))

	app.Router.Use(middleware.Logger)
	app.Router.Use(middleware.Recoverer)

	app.Router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))
	app.Router.Get("/swagger/doc.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./docs/swagger.json")
	})

	app.Router.Mount("/api/v1", api.Router())
	app.Router.Mount("/", service.Router())
	app.Router.Mount("/internal", internal.Router())
}

// RunServer
// @title           Expression Parser Service
// @version         0.1
// @description     API for use expression parser service
// @host      localhost:8080
func (app *Application) RunServer() {
	log.Println("Run server on port " + config.Service.Port + "\nVisit at on http://localhost:" + config.Service.Port)
	http.ListenAndServe(":"+config.Service.Port, app.Router)
}
