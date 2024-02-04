package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Aslamdesusa/click-tracker/app/http"
)

// AppConfig struct contains the config for the application
type AppConfig struct {
	Env string
}

// Config method will return the config object
func Config() string {
	env := getEnv()
	return env
}

// Init method will start the click tracker app by performing basic tasks
// like connecting to databases - Mongodb, maxmind, redis
func Init() {
	env := Config()

	Cleanup()
	http.Init(env)
}

// Cleanup method will clean the resources before exiting the app
func Cleanup() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		for sig := range c {
			fmt.Printf("%v Signal was received\n", sig)
			os.Exit(0)
		}
	}()
}

func getEnv() string {
	env := os.Getenv("GO_TRACKER_ENV")
	if env == "" {
		env = "production"
	}
	return env
}
