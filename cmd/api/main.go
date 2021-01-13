package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	accountapp "github.com/gustvision/backend-interview/pkg/account/app"
	accountsql "github.com/gustvision/backend-interview/pkg/account/sql"
	userapp "github.com/gustvision/backend-interview/pkg/user/app"
	usersql "github.com/gustvision/backend-interview/pkg/user/sql"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	_ "google.golang.org/grpc/encoding/gzip"
)

const (
	// Time allocated for init phase (connections + setup).
	initTO = 30 * time.Second
)

var version string

// run services.
func run(prog string, filename string) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, initTO)
	// no need for defer cancel
	_ = cancel

	// init logger (no config need)
	zerolog.TimeFieldFormat = ""

	log.Logger = log.Logger.With().Str("version", version).Str("exe", prog).Logger()

	// read config
	cfg := config{}
	if err := cfg.Populate(ctx, filename); err != nil {
		log.Error().Err(err).Msg("failed to read config")

		return
	}

	// setup postgres
	db, err := sql.Open("postgres", cfg.SQL)
	if err != nil {
		log.Error().Err(err).Msg("failed to open postgres")

		return
	}

	userSQLStore := usersql.Store{DB: db}
	accountSQLStore := accountsql.Store{DB: db}

	// init handler
	h := handler{}
	h.user = &userapp.App{
		Store: &userSQLStore,
	}
	h.account = &accountapp.App{
		Store:            &accountSQLStore,
		StoreTransaction: &accountSQLStore,
	}

	// serve http api
	go func(host string) {
		if err := h.listen(host); err != nil {
			log.Error().Err(err).Msg("failed to listen")
		}
	}(cfg.Host)
	log.Info().Msg("api up")

	// listen for signals
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

	for sig := range c {
		switch sig {
		case syscall.SIGHUP:
			fallthrough
		case syscall.SIGINT:
			fallthrough
		case syscall.SIGTERM:

			cancel()

			fmt.Println("successfully closed service")

			return
		}
	}
}

func main() {
	args := os.Args
	if len(args) != 2 { // nolint: gomnd
		fmt.Printf("Usage: ./%s configfile\n", args[0])

		return
	}

	run(args[0], args[1])
}
