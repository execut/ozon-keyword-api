package api

import (
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promauto"

    "github.com/execut/ozon-keyword-api/internal/repo"

    pb "github.com/execut/ozon-keyword-api/pkg/ozon-keyword-api"
)

var (
    totalKeywordNotFound = promauto.NewCounter(prometheus.CounterOpts{
        Name: "ozon_keyword_api_ozon_not_found_total",
        Help: "Total number of ozons that were not found",
    })
)

type ozonAPI struct {
    pb.UnimplementedOzonKeywordApiServiceServer
    repo repo.Repo
}

// NewKeywordAPI returns api of ozon-keyword-api service
func NewKeywordAPI(r repo.Repo) pb.OzonKeywordApiServiceServer {
    return &ozonAPI{repo: r}
}
