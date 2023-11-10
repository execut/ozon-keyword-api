package api

import (
    "context"
    "github.com/execut/ozon-keyword-api/pkg/ozon-keyword-api"
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
    return nil, status.Error(codes.Unimplemented, "RemoveKeywordV1 not implemented")
}
