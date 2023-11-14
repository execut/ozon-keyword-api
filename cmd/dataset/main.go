package main

import (
    "context"
    "fmt"
    "github.com/execut/ozon-keyword-api/internal/app/repo"
    "github.com/execut/ozon-keyword-api/internal/config"
    "github.com/execut/ozon-keyword-api/internal/database"
    "github.com/execut/ozon-keyword-api/internal/model"
    repo2 "github.com/execut/ozon-keyword-api/internal/repo"
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
    ctx := context.Background()

    tx, _ := db.Begin()
    eventsRepo := repo.NewEventRepo(db)
    keywordRepo := repo2.NewRepo(db)
    events := []model.KeywordEvent{}
    for i := 0; i < 1000; i++ {
        keyword := model.Keyword{
            Name: fmt.Sprintf("Test keyword %v", i),
        }
        _, err := keywordRepo.Add(ctx, &keyword)
        if err != nil {
            panic(err)
        }

        events = append(events, model.KeywordEvent{
            Type:   model.Created,
            Entity: &keyword,
        })
    }

    err = eventsRepo.Add(events)
    if err != nil {
        panic(err)
    }

    //events, err = eventsRepo.Lock(10000)
    //if err != nil {
    //    panic(err)
    //}
    //
    //var ids []uint64
    //for _, event := range events {
    //    ids = append(ids, event.ID)
    //}
    //
    //err = eventsRepo.Unlock(ids)
    //if err != nil {
    //    panic(err)
    //}
    //
    //eventsRepo.Remove(ids)

    tx.Commit()
}
