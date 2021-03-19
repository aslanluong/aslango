package main

import (
	"aslango/pkg/routes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"gopkg.in/yaml.v2"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/go", routes.LinkRoute)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Oops, nothing here :("))
	})

	listenAndServe := func(port string) {
		LoadEnvVars()
		http.ListenAndServe(":"+port, r)
	}

	if port, ok := os.LookupEnv("PORT"); ok {
		listenAndServe(port)
	} else {
		listenAndServe("9999")
	}
}

type EnvVars struct {
	Environment string `yaml:"APP_ENV"`
	MongoDB     struct {
		Uri string `yaml:"uri"`
	}
}

func LoadEnvVars() {
	env, err := ioutil.ReadFile("./env.yaml")
	if err != nil {
		return
	}

	var envVars map[string]interface{}
	if err := yaml.Unmarshal(env, &envVars); err != nil {
		fmt.Println("fail", err)
	} else {
		var setEnvVars func(rootKey string, values map[string]interface{})
		setEnvVars = func(rootKey string, values map[string]interface{}) {
			for key, value := range values {
				_, ok := value.(map[interface{}]interface{})
				if ok {
					valuesMarshal, _ := yaml.Marshal(value)
					var valuesUnmarshal map[string]interface{}
					yaml.Unmarshal(valuesMarshal, &valuesUnmarshal)
					setEnvVars(rootKey+key+".", valuesUnmarshal)
				} else {
					if valueStr, ok := value.(string); ok {
						os.Setenv(rootKey+key, valueStr)
					}
					if valueInt, ok := value.(int); ok {
						os.Setenv(rootKey+key, strconv.Itoa(valueInt))
					}
				}
			}
		}
		setEnvVars("", envVars)
		os.Setenv("ENV_UNMARSHAL_STRING", string(env))
	}
}

func GetEnvVars() *EnvVars {
	env := os.Getenv("ENV_UNMARSHAL_STRING")
	envVars := &EnvVars{}
	if err := yaml.Unmarshal([]byte(env), &envVars); err != nil {
		fmt.Println("fail", err)
	}
	return envVars
}
