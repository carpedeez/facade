package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"

	"github.com/carpedeez/store/config"
	"github.com/carpedeez/store/database"
	"github.com/carpedeez/store/facade"
	"github.com/carpedeez/store/logger"
)

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
	wg := &sync.WaitGroup{}

	c, err := config.GetConfig()
	if err != nil {
		fmt.Printf("failed to get config: %v", err)
		os.Exit(1)
	}

	log := logger.GetLogger(c.LogConfig)

	q, err := database.OpenPostgres(ctx, c.DBConfig)
	if err != nil {
		log.Fatal().Msgf("failed to open database: %v\n", err)
	}
	defer func() {
		err = q.Close()
		if err != nil {
			log.Error().Err(err).Msg("error closing database")
		}
	}()

	wg.Add(1)
	go facade.Run(wg, ctx, log, c.FacadeConfig, q)

	wg.Wait()
}
