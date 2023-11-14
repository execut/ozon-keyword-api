package api

import (
    "context"
    "github.com/execut/ozon-keyword-api/pkg/ozon-keyword-api"
    "github.com/rs/zerolog/log"
)

func (o *ozonAPI) ListKeywordV1(ctx context.Context, req *ozon_keyword_api.ListKeywordV1Request) (*ozon_keyword_api.ListKeywordV1Response, error) {
    log.Debug().Msg("ListKeywordV1")
    keywords, err := o.repo.List(ctx, 100, 0)
    if err != nil {
        return nil, err
    }

    var result []*ozon_keyword_api.Keyword
    for _, keyword := range keywords {
        result = append(result, &ozon_keyword_api.Keyword{
            Id:   keyword.ID,
            Name: keyword.Name,
        })
    }

    return &ozon_keyword_api.ListKeywordV1Response{Items: result}, nil
}
