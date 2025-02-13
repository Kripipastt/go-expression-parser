package application

import (
	"bufio"
	"fmt"
	"github.com/Kripipastt/go-expression-parser/internal/application/handlers"
	"github.com/Kripipastt/go-expression-parser/pkg/parser"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"os"
	"strings"
)

type Config struct {
	Port string
}

func LoadConfig() *Config {
	godotenv.Load()
	conf := new(Config)
	conf.Port = os.Getenv("PORT")
	if conf.Port == "" {
		conf.Port = "8080"
	}
	return conf
}

type Application struct {
	Router *chi.Mux
	Config *Config
}

func NewApplication() *Application {
	return &Application{Config: LoadConfig(), Router: chi.NewRouter()}
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
		result, err := parser.Calc(text)
		if err != nil {
			fmt.Println(text, " failed to calc expression: ", err)
		} else {
			fmt.Println(text, " = ", result)
		}

	}
}

func (app *Application) MountHandlers() {
	app.Router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type"},
	}))

	app.Router.Use(middleware.Logger)
	app.Router.Use(middleware.Recoverer)

	app.Router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))
	app.Router.Get("/swagger/doc.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./docs/swagger.json")
	})

	app.Router.Post("/api/v1/calculate", handlers.CalcHandler)
	app.Router.Get("/ping", handlers.PingHandler)
}

// RunServer
// @title           Expression Parser Service
// @version         0.1
// @description     API for use expression parser service
// @host      localhost:8080
func (app *Application) RunServer() {
	log.Println("Run server on port " + app.Config.Port)
	http.ListenAndServe(":"+app.Config.Port, app.Router)
}
