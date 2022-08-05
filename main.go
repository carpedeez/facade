package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/carpedeez/store/config"
	"github.com/carpedeez/store/database"
	"github.com/carpedeez/store/facade"
)

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)

	c, err := config.GetConfig()
	if err != nil {
		fmt.Printf("failed to get config: %v", err)
		os.Exit(1)
	}

	q, err := database.OpenPostgres(ctx, c.DBConfig)
	if err != nil {
		fmt.Printf("failed to open database: %v\n", err)
		os.Exit(1)
	}

	go facade.Run(c.FacadeConfig, q)

	<-ctx.Done()
}
