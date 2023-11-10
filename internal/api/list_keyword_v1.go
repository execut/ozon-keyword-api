package api

import (
    "context"
    "github.com/execut/ozon-keyword-api/pkg/ozon-keyword-api"
    "github.com/rs/zerolog/log"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

func (o *ozonAPI) ListKeywordV1(ctx context.Context, req *ozon_keyword_api.ListKeywordV1Request) (*ozon_keyword_api.ListKeywordV1Response, error) {
    log.Debug().Msg("ListKeywordV1")

    return nil, status.Error(codes.Unimplemented, "ListKeywordV1 not implemented")
}
