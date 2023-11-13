package api

import (
    "context"
    "errors"
    "github.com/execut/ozon-keyword-api/internal/repo"
    "github.com/execut/ozon-keyword-api/pkg/ozon-keyword-api"
    pb "github.com/execut/ozon-keyword-api/pkg/ozon-keyword-api"
    "github.com/rs/zerolog/log"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

func (o *ozonAPI) RemoveKeywordV1(ctx context.Context, req *ozon_keyword_api.RemoveKeywordV1Request) (*ozon_keyword_api.RemoveKeywordV1Response, error) {
    if err := req.Validate(); err != nil {
        log.Error().Err(err).Msg("RemoveKeywordV1 - invalid argument")

        return nil, status.Error(codes.InvalidArgument, err.Error())
    }

    log.Debug().Msg("RemoveKeywordV1")

    err := o.repo.Remove(ctx, req.KeywordId)
    found := errors.Is(err, repo.ErrKeywordNotFound)
    if !found && err != nil {
        log.Error().Err(err).Msg("RemoveKeywordV1 -- failed")

        return nil, status.Error(codes.Internal, err.Error())
    }

    log.Debug().Msg("RemoveKeywordV1 - success")

    return &pb.RemoveKeywordV1Response{
        Found: found,
    }, nil
}
