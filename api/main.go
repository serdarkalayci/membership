// Package: main
package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/serdarkalayci/membership/api/adapters/comm/rest"
	"github.com/serdarkalayci/membership/api/adapters/data/postgres"

	"github.com/nicholasjackson/env"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/serdarkalayci/membership/api/util"
)

var bindAddress = env.String("BASE_URL", false, ":5500", "Bind address for rest server")

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	util.SetConstValues()
	util.SetLogLevels()
	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()

	// var dbContext application.DataContextCarrier
	// var err error
	dbContext, _ := postgres.NewDataContext()

	s := rest.NewRestServer(dbContext)

	// start the http server
	go func() {
		log.Debug().Msgf("Starting server on %s", *bindAddress)

		err := s.RunServer(bindAddress)
		if err != nil {
			log.Error().Err(err).Msg("error starting rest server")
			os.Exit(1)
		}
	}()

	// trap sigterm or interrupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Info().Msgf("Got signal: %s", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(ctx)
}
