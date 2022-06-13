package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/MalukiMuthusi/cryptocurrencies/internal/handlers"
	"github.com/MalukiMuthusi/cryptocurrencies/internal/logger"
	"github.com/MalukiMuthusi/cryptocurrencies/internal/store/fauna"
	"github.com/MalukiMuthusi/cryptocurrencies/internal/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	// Create a deadline to wait for.
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*5, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	// initialize the storage service
	store, err := fauna.New()
	if err != nil {
		logger.Log.Fatal("failed to initialize database ", err)
	}

	// initialize the API handler
	r := handlers.SetUpRouter(store, DebugPrintRoute)

	port := viper.GetString(utils.Port)

	srv := &http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%s", port),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {

		if err := srv.ListenAndServe(); err != nil {
			logger.Log.WithFields(logrus.Fields{"serverListenError": err}).Info("server listen error")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.

	c := make(chan os.Signal, 1)

	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	logger.Log.WithField("shuttingDown", "shutdown").Info("shutdown server")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)

	os.Exit(0)
}

func init() {
	viper.AutomaticEnv()

	viper.SetEnvPrefix(utils.AppName)

	BindEnvs()

	viper.SetDefault(utils.Port, "8080")

	CheckMustBeSetEnvs([]string{utils.FaunaSecret})
}

// BindEnvs bind environment variables to specific names
func BindEnvs() {

	viper.BindEnv(utils.Port, "PORT")

}

// CheckMustBeSetEnvs ensures that the provided envs have been set
func CheckMustBeSetEnvs(envs []string) {

	for _, v := range envs {
		EnvMustBeSet(v)
	}

}

// EnvMustBeSet checks that a value has been provided
func EnvMustBeSet(key string) {
	if !viper.IsSet(key) {
		logger.Log.WithField(key, false).Fatal("not set")
	}
}

// DebugPrintRoute prints all the endpoints available
func DebugPrintRoute(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	logger.Log.WithFields(logrus.Fields{"httpMethod": httpMethod, "path": absolutePath, "handlerName": handlerName, "nuHandlers": nuHandlers}).Info("endpointRequest")
}
