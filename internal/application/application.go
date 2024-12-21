package application

import (
	"bufio"
	"fmt"
	"github.com/Kripipastt/go-expression-parser/internal/application/handlers"
	logger2 "github.com/Kripipastt/go-expression-parser/pkg/logger"
	"github.com/Kripipastt/go-expression-parser/pkg/logger/middleware"
	"github.com/Kripipastt/go-expression-parser/pkg/parser"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
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
	config *Config
	logger *zap.Logger
}

func New() *Application {
	return &Application{config: LoadConfig(), logger: logger2.LoggerCreate()}
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

func (app *Application) RunServer() {
	router := mux.NewRouter()
	router.Use(middleware.LoggingMiddleware(app.logger))

	router.HandleFunc("/api/v1/calculate", handlers.CalcHandler)
	router.HandleFunc("/ping", handlers.PingHandler)

	http.Handle("/", router)

	app.logger.Info("Run server on port " + app.config.Port)
	http.ListenAndServe(":"+app.config.Port, router)
}
