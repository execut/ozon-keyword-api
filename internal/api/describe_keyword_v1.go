package api

import (
    "context"
    pb "github.com/execut/ozon-keyword-api/pkg/ozon-keyword-api"
    "github.com/rs/zerolog/log"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

func (o *ozonAPI) DescribeKeywordV1(
    ctx context.Context,
    req *pb.DescribeKeywordV1Request,
) (*pb.DescribeKeywordV1Response, error) {

    if err := req.Validate(); err != nil {
        log.Error().Err(err).Msg("DescribeKeywordV1 - invalid argument")

        return nil, status.Error(codes.InvalidArgument, err.Error())
    }

    keyword, err := o.repo.Get(ctx, req.KeywordId)
    if err != nil {
        log.Error().Err(err).Msg("DescribeKeywordV1 -- failed")

        return nil, status.Error(codes.Internal, err.Error())
    }

    if keyword == nil {
        log.Debug().Uint64("keywordId", req.KeywordId).Msg("keyword not found")
        totalKeywordNotFound.Inc()

        return nil, status.Error(codes.NotFound, "keyword not found")
    }

    log.Debug().Msg("DescribeKeywordV1 - success")

    return &pb.DescribeKeywordV1Response{
        Value: &pb.Keyword{
            Id:   keyword.ID,
            Name: keyword.Name,
        },
    }, nil
}
