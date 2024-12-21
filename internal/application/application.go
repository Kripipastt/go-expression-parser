package application

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/Kripipastt/go-expression-parser/pkg/parser"
	"github.com/joho/godotenv"
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
}

func New() *Application {
	return &Application{config: LoadConfig()}
}

type Request struct {
	Expression string `json:"expression"`
}

type Response struct {
	Result float64 `json:"result"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	var request = Request{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Internal server error"})
		return
	}

	defer r.Body.Close()

	result, err := parser.Calc(request.Expression)

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Expression is not valid"})
		return
	}

	json.NewEncoder(w).Encode(Response{Result: result})
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
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/calculate", CalcHandler)
	http.ListenAndServe(":"+app.config.Port, mux)
}
