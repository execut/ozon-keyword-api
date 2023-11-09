package api

import (
    "context"
    "github.com/execut/ozon-keyword-api/pkg/ozon-keyword-api"
    "github.com/rs/zerolog/log"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

func (o *ozonAPI) CreateKeywordV1(ctx context.Context, req *ozon_keyword_api.CreateKeywordV1Request) (*ozon_keyword_api.CreateKeywordV1Response, error) {
    log.Debug().Msg("CreateKeywordV1")
    return nil, status.Error(codes.Unimplemented, "CreateKeywordV1 not implemented")
}
