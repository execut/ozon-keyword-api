package main

import (
    "github.com/execut/ozon-keyword-api/internal/app/repo"
    "github.com/execut/ozon-keyword-api/internal/app/sender"
    "github.com/execut/ozon-keyword-api/internal/config"
    "github.com/execut/ozon-keyword-api/internal/database"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/execut/ozon-keyword-api/internal/app/retranslator"
    "github.com/rs/zerolog/log"

    _ "github.com/jackc/pgx/v4"
    _ "github.com/jackc/pgx/v4/stdlib"
    _ "github.com/lib/pq"
)

func main() {
    if err := config.ReadConfigYML("config.yml"); err != nil {
        log.Fatal().Err(err).Msg("Failed init configuration")
    }
    cfg := config.GetConfigInstance()

    db, err := database.NewPostgres(database.GenerateDsnFromConfig(cfg), cfg.Database.Driver)
    if err != nil {
        log.Fatal().Err(err).Msg("Failed init postgres")
    }
    defer db.Close()

    eventsRepo := repo.NewEventRepo(db)

    sigs := make(chan os.Signal, 1)

    retranslatorConfig := retranslator.Config{
        ChannelSize:         2,
        ConsumersCount:      2,
        ConsumerInterval:    time.Millisecond,
        ConsumerBatchSize:   2,
        ProducersCount:      2,
        WorkerCount:         2,
        Repo:                eventsRepo,
        Sender:              sender.NewStubEventSender(),
        WorkerpoolBatchSize: 100,
    }

    retranslator := retranslator.NewRetranslator(retranslatorConfig)
    retranslator.Start()

    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

    <-sigs
}
