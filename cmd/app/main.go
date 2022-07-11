package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"

	"github.com/spf13/viper"
)

type Config struct {
	Port string `mapstructure:"PORT"`
}

func main() {
	// config, err := loadConfig(".")
	// must(err)

	port := os.Getenv("PORT")
	fmt.Println(port)

	// fmt.Println(config)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	server := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			io.WriteString(w, "hello!")
		}),
	}
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			must(err)
		}
	}()

	shutdownSignalCh := make(chan os.Signal)
	signal.Notify(shutdownSignalCh, os.Interrupt)
	sig := <-shutdownSignalCh
	fmt.Printf("got signal: %v", sig)
	if err := server.Shutdown(ctx); err != nil {
		must(err)
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func loadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}

	return
}
