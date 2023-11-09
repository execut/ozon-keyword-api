package api

import (
    "context"

    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promauto"

    "github.com/rs/zerolog/log"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"

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

func (o *ozonAPI) DescribeKeywordV1(
    ctx context.Context,
    req *pb.DescribeKeywordV1Request,
) (*pb.DescribeKeywordV1Response, error) {

    if err := req.Validate(); err != nil {
        log.Error().Err(err).Msg("DescribeKeywordV1 - invalid argument")

        return nil, status.Error(codes.InvalidArgument, err.Error())
    }

    ozon, err := o.repo.DescribeKeyword(ctx, req.KeywordId)
    if err != nil {
        log.Error().Err(err).Msg("DescribeKeywordV1 -- failed")

        return nil, status.Error(codes.Internal, err.Error())
    }

    if ozon == nil {
        log.Debug().Uint64("ozonId", req.KeywordId).Msg("ozon not found")
        totalKeywordNotFound.Inc()

        return nil, status.Error(codes.NotFound, "ozon not found")
    }

    log.Debug().Msg("DescribeKeywordV1 - success")

    return &pb.DescribeKeywordV1Response{
        Value: &pb.Keyword{
            Id:  ozon.ID,
            Foo: ozon.Foo,
        },
    }, nil
}
